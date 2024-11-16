#include "nvr_s.h"

nvr_s::nvr_s(Config config)
{
    config = config;
}

nvr_s::~nvr_s()
{
}

void nvr_s::nvrDownload(const httplib::Request& request, httplib::Response& response, DownloadForm params) {
    std::string filepath{};
    if (Error err = download(params, filepath); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err), "text/plain");
        return;
    }
    
    std::ifstream fp(filepath, std::ios::binary);
    if (!fp.is_open()) {
        response.status = 200;
        response.set_content(error(Error::FileOpenFailed), "text/plain");
        return;
    }
    std::string content((std::istreambuf_iterator<char>(fp)), std::istreambuf_iterator<char>());
    fp.close();

    response.status = 200;
    response.set_header("Content-Type", "application/octet-stream");
    response.set_header("Content-Disposition", "attachment; filename=" + filepath);
    response.set_header("Content-Transfer-Encoding", "binary");
    response.set_content(content, "application/octet-stream");

    return;
}

Error nvr_s::download(DownloadForm& params, std::string& filepathR) 
{
    NET_DVR_Init();

    // 设置连接时间与重连时间
    NET_DVR_SetConnectTime(2000, 1);
    NET_DVR_SetReconnect(10000, true);

    NET_DVR_DEVICEINFO_V30 deviceInfo{};
    NvrConfig nvrConfig = config.nvr;
    char* host = new char[nvrConfig.host.length() + 1]; strcpy(host, nvrConfig.host.c_str());
    char* user = new char[nvrConfig.user.length() + 1]; strcpy(user, nvrConfig.user.c_str());
    char* password = new char[nvrConfig.password.length() + 1]; strcpy(password, nvrConfig.password.c_str());
    LONG userId = NET_DVR_Login_V30(
        host, 
        nvrConfig.port, 
        user, 
        password, 
        &deviceInfo);
    if (userId != Error::Nil) {
        NET_DVR_Cleanup();
        return Error::UserLoginFailed;
    }
    spdlog::info("login success...");

    // 注意：目前SDK私有协议对接时64路以下的NVR的IP通道号是从33开始的，64路以及以上的NVR的IP通道从1开始
    NET_DVR_PLAYCOND downloadCond{};
    downloadCond.dwChannel = params.channel;
    downloadCond.struStartTime.dwYear = params.startTime.year;
    downloadCond.struStartTime.dwMonth = params.startTime.month;
    downloadCond.struStartTime.dwDay = params.startTime.day;
    downloadCond.struStartTime.dwHour = params.startTime.hour;
    downloadCond.struStartTime.dwMinute = params.startTime.minute;
    downloadCond.struStartTime.dwSecond = params.startTime.second;

    downloadCond.struStopTime.dwYear = params.endTime.year;
    downloadCond.struStopTime.dwMonth = params.endTime.month;
    downloadCond.struStopTime.dwDay = params.endTime.day;
    downloadCond.struStopTime.dwHour = params.endTime.hour;
    downloadCond.struStopTime.dwMinute = params.endTime.minute;
    downloadCond.struStopTime.dwSecond = params.endTime.second;

    // 按时间下载
    std::string now{};
    currentTimeStr(now);
    std::hash<std::string> hasher; 
    std::stringstream ss; ss << std::hex << hasher(now);
    std::string filepath = "/tmp/"+ss.str() + ".mp4";
    char* filename = new char[filepath.length() + 1]; strcpy(filename, filepath.c_str());
    int hPlayback = NET_DVR_GetFileByTime_V40(userId, filename, &downloadCond);
    if (hPlayback < 0) {
        spdlog::error(error(Error::DvrGetFileByTimeV40Failed));
        NET_DVR_Logout(userId);
        NET_DVR_Cleanup();
        return Error::DvrGetFileByTimeV40Failed;
    }

    // 开始下载
    if (!NET_DVR_PlayBackControl_V40(hPlayback, NET_DVR_PLAYSTART, NULL, 0, NULL, NULL)) {
        spdlog::error(error(Error::PlaybackControlFailed));
        NET_DVR_Logout(userId);
        NET_DVR_Cleanup();
        return Error::PlaybackControlFailed;
    }

    int nPos = 0;
    for (nPos = 0; nPos < 100 && nPos >= 0; nPos = NET_DVR_GetDownloadPos(hPlayback))
    {
        ss.clear();ss.str("");ss << nPos; 
        spdlog::info("Be downloading..." + ss.str() + "%");
        sleep(2);
    }
    if (!NET_DVR_StopGetFile(hPlayback))
    {
        spdlog::error(error(Error::FileStopGetFailed));
        NET_DVR_Logout(userId);
        NET_DVR_Cleanup();
        return Error::FileStopGetFailed;
    }
    if (nPos < 0 || nPos>100)
    {
        spdlog::error(error(Error::FileDownloadFailed));
        NET_DVR_Logout(userId);
        NET_DVR_Cleanup();
        return Error::FileDownloadFailed;
    }
    ss.clear();ss.str("");ss << nPos; 
    spdlog::info("Be downloading..." + ss.str() + "%");
    NET_DVR_Logout(userId);
    NET_DVR_Cleanup();
    filepathR = filename;
    return Error::Nil;
}

Error nvr_s::currentTimeStr(std::string& now) 
{
    std::time_t now_t = std::time(nullptr);
    char buf[100];
    std::strftime(buf, sizeof(buf), "%Y-%m-%d %H:%M:%S", std::localtime(&now_t));
    now = std::string(buf);
    return Error::Nil;
}

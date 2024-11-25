#include "nvr_s.h"

nvr_s::nvr_s(Config config)
{
    this->config = config;
}

nvr_s::~nvr_s()
{
    spdlog::info("nvr_s object is being deleted");
}

void nvr_s::nvr_channel(const httplib::Request& request, httplib::Response& response)
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
    if (userId < Error::Nil) {
        NET_DVR_Cleanup();
        response.status = 200;
        response.set_content(error(Error::UserLoginFailed), "text/plain");
        return;
    }
    spdlog::info("login success...");

    NET_DVR_IPPARACFG_V40 ipcfg{};
    DWORD bytesReturned = 0;
    ipcfg.dwSize = sizeof(NET_DVR_IPPARACFG_V40);
    int iGroupNo = 0;
    bool resCode = NET_DVR_GetDVRConfig(userId, NET_DVR_GET_IPPARACFG_V40, iGroupNo, &ipcfg, sizeof(NET_DVR_IPPARACFG_V40), &bytesReturned);
    if (!resCode) {
        NET_DVR_Logout(userId);
        NET_DVR_Cleanup();
        response.status = 200;
        response.set_content(error(Error::DvrGetConfigFailed), "text/plain");
        return;
    }

    nlohmann::json records = nlohmann::json::array();
    for (int i = 0; i != ipcfg.dwDChanNum; i ++) 
    {
        NET_DVR_PICCFG_V30 channelInfo;
        bytesReturned = 0;
        channelInfo.dwSize = sizeof(NET_DVR_PICCFG_V30);
        int channelNum = i + ipcfg.dwStartDChan;
        NET_DVR_GetDVRConfig(userId, NET_DVR_GET_PICCFG_V30, channelNum, &channelInfo,  sizeof(NET_DVR_PICCFG_V30), &bytesReturned);

        std::string chanName(reinterpret_cast<const char*>(channelInfo.sChanName));
        std::string username(reinterpret_cast<const char*>(ipcfg.struIPDevInfo[i].sUserName));
        std::string password(reinterpret_cast<const char*>(ipcfg.struIPDevInfo[i].sPassword));
        std::string deviceId(reinterpret_cast<const char*>(ipcfg.struIPDevInfo[i].szDeviceID));

        nlohmann::json record {};
        record["chanNum"] = channelNum;
        record["chanName"] = chanName;
        record["username"] = username;
        record["password"] = password;
        record["deviceId"] = deviceId;
        record["host"] = ipcfg.struIPDevInfo[i].struIP.sIpV4;
        record["port"] = ipcfg.struIPDevInfo[i].wDVRPort;
        records.push_back(record);
    }

    nlohmann::json res;
    res["data"] = records;

    // 释放 SDK 资源
    NET_DVR_Logout(userId);
    NET_DVR_Cleanup();

    response.status = 200;
    response.set_content(res.dump(), "application/json");

    return;
}

void nvr_s::nvr_download(const httplib::Request& request, httplib::Response& response, DownloadForm params) 
{
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
    if (userId < Error::Nil) {
        spdlog::error(NET_DVR_GetLastError());
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
    current_time_str(now);
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

Error nvr_s::current_time_str(std::string& now) 
{
    std::time_t now_t = std::time(nullptr);
    char buf[100];
    std::strftime(buf, sizeof(buf), "%Y-%m-%d %H:%M:%S", std::localtime(&now_t));
    now = std::string(buf);
    return Error::Nil;
}

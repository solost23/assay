#include <iostream>
#include <fstream>
#include <vector>
#include <ctime>
#include <sstream>

#include "httplib.h"
#include "HCNetSDK.h"

#include "../forms/nvr.h"
#include "../configs/config.h"

int download(DownloadForm&, std::string&);
int currentTimeStr(std::string&);

int nvrDownload(const httplib::Request& request, httplib::Response& response, DownloadForm params) {
    std::string filepath;
    int res = download(params, filepath);
    if (res != 0) {
        std::cerr << "err: " << res << std::endl;
        return -1;
    }
    
    std::ifstream fp(filepath, std::ios::binary);
    if (!fp.is_open()) {
        response.status = 404;
        response.set_content("file not found", "text/plain");
        return -1;
    }
    std::string content((std::istreambuf_iterator<char>(fp)), std::istreambuf_iterator<char>());
    fp.close();

    response.status = 200;
    response.set_header("Content-Type", "application/octet-stream");
    response.set_header("Content-Disposition", "attachment; filename=" + filepath);
    response.set_header("Content-Transfer-Encoding", "binary");
    response.set_content(content, "application/octet-stream");

    return 0;
}

int download(DownloadForm& params, std::string& filepathR) {
    NET_DVR_Init();

    // 设置连接时间与重连时间
    NET_DVR_SetConnectTime(2000, 1);
    NET_DVR_SetReconnect(10000, true);

    NET_DVR_DEVICEINFO_V30 deviceInfo;
    NvrConfig nvrConfig = serverConfig.nvr;
    char* host = new char[nvrConfig.host.length() + 1]; strcpy(host, nvrConfig.host.c_str());
    char* user = new char[nvrConfig.user.length() + 1]; strcpy(user, nvrConfig.user.c_str());
    char* password = new char[serverConfig.nvr.password.length() + 1]; strcpy(password, serverConfig.nvr.password.c_str());
    LONG userId = NET_DVR_Login_V30(
        host, 
        nvrConfig.port, 
        user, 
        password, 
        &deviceInfo);
    if (userId != 0) {
        std::cerr << "Login error, " << NET_DVR_GetLastError() << std::endl; NET_DVR_Cleanup(); return -1;
    }

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
        std::cerr << "NET_DVR_GetFileByTime_V40 fail, last err: " << NET_DVR_GetLastError() << std::endl; NET_DVR_Logout(userId); NET_DVR_Cleanup(); return -1;
    }

    // 开始下载
    if (!NET_DVR_PlayBackControl_V40(hPlayback, NET_DVR_PLAYSTART, NULL, 0, NULL, NULL)) {
        std::cerr << "Play back control failed " << NET_DVR_GetLastError() << std::endl; NET_DVR_Logout(userId); NET_DVR_Cleanup(); return -1;
    }

    int nPos = 0;
    for (nPos = 0; nPos < 100 && nPos >= 0; nPos = NET_DVR_GetDownloadPos(hPlayback))
    {
        std::cout << "Be downloading... " << nPos << "%" << std::endl;
        sleep(2);
    }
    if (!NET_DVR_StopGetFile(hPlayback))
    {
        std::cerr << "failed to stop get file " << NET_DVR_GetLastError() << std::endl; NET_DVR_Logout(userId); NET_DVR_Cleanup(); return -1;
    }
    if (nPos < 0 || nPos>100)
    {
        std::cout << "download err " << NET_DVR_GetLastError() << std::endl; NET_DVR_Logout(userId); NET_DVR_Cleanup(); return -1;
    }
    std::cout << "Be downloading..." << nPos << std::endl;
    NET_DVR_Logout(userId);
    NET_DVR_Cleanup();
    filepathR = filename;
    return 0;
}

// 获取当前时间
int currentTimeStr(std::string& now) {
    std::time_t now_t = std::time(nullptr);
    char buf[100];
    std::strftime(buf, sizeof(buf), "%Y-%m-%d %H:%M:%S", std::localtime(&now_t));
    now = std::string(buf);
    return 0;
}

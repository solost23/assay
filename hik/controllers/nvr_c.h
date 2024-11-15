#pragma once

#include "../services/nvr_s.h"

class nvr_c
{
private:
    ServerConfig serverConfig;

    Error parse(const httplib::Request& request, int& value, std::string field);
public:
    nvr_c(ServerConfig config);
    ~nvr_c();

    // 视频下载
    void download(const httplib::Request& request, httplib::Response& response);
};




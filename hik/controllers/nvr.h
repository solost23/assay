#pragma once

#include "services/nvr.h"

class NvrController
{
private:
    Config config;

    Error parse(const httplib::Request& request, int& value, std::string field);
public:
    NvrController(const Config& conf): config(conf) {};
    ~NvrController();

    // 通道列表
    void channel(const httplib::Request& request, httplib::Response& response);
    // 视频下载
    void download(const httplib::Request& request, httplib::Response& response);
};

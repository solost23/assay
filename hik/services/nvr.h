#pragma once

#include <iostream>
#include <fstream>
#include <vector>
#include <ctime>
#include <sstream>
#include <memory>

#include "httplib.h"
#include <nlohmann/json.hpp>
#include "spdlog/spdlog.h"

#include "HCNetSDK.h"

#include "configs/config.h"
#include "forms/nvr.h"

class NvrService
{
private:
    std::shared_ptr<Config> config;
    
    Error download(DownloadForm&, std::string&);
    Error current_time_str(std::string& now);
public:
    NvrService(const std::shared_ptr<Config> conf): config(conf) {};
    ~NvrService();

    void nvr_channel(const httplib::Request& request, httplib::Response& response);
    void nvr_download(const httplib::Request& request, httplib::Response& response, DownloadForm params);
};

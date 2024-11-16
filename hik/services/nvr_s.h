#pragma once

#include <iostream>
#include <fstream>
#include <vector>
#include <ctime>
#include <sstream>

#include "HCNetSDK.h"
#include <httplib.h>
#include "spdlog/spdlog.h"
#include "../configs/config_s.h"
#include "../forms/nvr_f.h"


class nvr_s
{
private:
    Config config;
    
    Error download(DownloadForm&, std::string&);
    Error currentTimeStr(std::string& now);
public:
    nvr_s(Config config);
    ~nvr_s();

    void nvrDownload(const httplib::Request& request, httplib::Response& response, DownloadForm params) ;
};

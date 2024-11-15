#pragma once

#include <iostream>
#include <fstream>
#include <vector>
#include <ctime>
#include <sstream>

#include "HCNetSDK.h"
#include <httplib.h>
#include "../configs/config.h"
#include "../forms/nvr.h"


class nvr_s
{
private:
    ServerConfig serverConfig;
    
    Error download(DownloadForm&, std::string&);
    Error currentTimeStr(std::string& now);
public:
    nvr_s(ServerConfig config);
    ~nvr_s();

    void nvrDownload(const httplib::Request& request, httplib::Response& response, DownloadForm params) ;
};

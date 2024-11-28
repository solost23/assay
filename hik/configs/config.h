#pragma once

#include <iostream>
#include <fstream>

#include "spdlog/spdlog.h"
#include "yaml-cpp/yaml.h"

#include "util/error.h"

class Config
{
    private:
        struct NvrConfig
        {
            std::string host;
            int port;
            std::string user;
            std::string password;
        };

        std::string name;
        std::string version;
        int port;
        int thread;
        NvrConfig nvr;
    public:
        Config(std::string filepath);

        std::string get_name();
        std::string get_version();
        int get_port();
        int get_thread();
        
        std::string get_nvr_host();
        int get_nvr_port();
        std::string get_nvr_user();
        std::string get_nvr_password();
};

#pragma once

#include <iostream>
#include <fstream>

#include "spdlog/spdlog.h"

#include "yaml-cpp/yaml.h"

#include "../util/error.h"

struct NvrConfig 
{
    std::string host;
    int port;
    std::string user;
    std::string password;
};

struct Config 
{
    std::string name;
    std::string version;
    int port;
    int thread;
    NvrConfig nvr;
};

class config_s
{
private:
    
public:
    Config config;

    config_s(std::string path);
    ~config_s();

    // Error initConfig(std::string configPath, ServerConfig& serverConfig);
};

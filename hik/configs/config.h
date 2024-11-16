#pragma once

#include <iostream>
#include <fstream>

#include "yaml-cpp/yaml.h"

#include "../util/error.h"

struct NvrConfig 
{
    std::string host;
    int port;
    std::string user;
    std::string password;
};

struct ServerConfig 
{
    std::string name;
    std::string version;
    int port;
    int thread;
    NvrConfig nvr;
};

class config
{
private:
public:
    config();
    ~config();

    Error initConfig(std::string configPath, ServerConfig& serverConfig);
};

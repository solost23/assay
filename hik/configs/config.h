#include <iostream>
#include <fstream>
#include "yaml-cpp/yaml.h"

struct NvrConfig {
    std::string host;
    int port;
    std::string user;
    std::string password;
};

struct ServerConfig {
    std::string name;
    std::string version;
    int port;
    NvrConfig nvr;
};

extern ServerConfig serverConfig;

int config(std::string);

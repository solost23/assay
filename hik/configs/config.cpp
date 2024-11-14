#include <iostream>
#include <fstream>
#include "yaml-cpp/yaml.h"

#include "config.h"

ServerConfig serverConfig{};

int config(std::string configPath) {
    std::ifstream fp(configPath);
    if (!fp) {
        std::cerr << "Open Config File err" << std::endl;
        return -1;
    }

    YAML::Node config = YAML::Load(fp);
        
    try {
        serverConfig.name = config["name"].as<std::string>();
        serverConfig.version = config["version"].as<std::string>();
        serverConfig.port = config["port"].as<int>();
        serverConfig.nvr.host = config["nvr"]["host"].as<std::string>();
        serverConfig.nvr.port = config["nvr"]["port"].as<int>();
        serverConfig.nvr.user = config["nvr"]["user"].as<std::string>();
        serverConfig.nvr.password = config["nvr"]["password"].as<std::string>();
    } catch (const YAML::Exception& e) {
        std::cerr << "Read Config File err: " << e.what() << std::endl;
        return -1;
    }
    return 0;
}

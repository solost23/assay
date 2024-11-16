#include "config.h"

config::config()
{
}

config::~config()
{
}

Error config::initConfig(std::string configPath, ServerConfig& serverConfig)
{
    std::ifstream fp(configPath);
    if (!fp) {
        return Error::FileOpenFailed;
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
        return Error::FileReadFailed;
    }
    return Error::Nil;
}

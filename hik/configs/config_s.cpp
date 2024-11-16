#include "config_s.h"

config_s::config_s(std::string path)
{
    std::ifstream fp(path);
    if (!fp) {
        spdlog::critical(error(Error::FileOpenFailed));
    }

    YAML::Node config_y = YAML::Load(fp);
        
    try {
        config.name = config_y["name"].as<std::string>();
        config.version = config_y["version"].as<std::string>();
        config.port = config_y["port"].as<int>();
        config.nvr.host = config_y["nvr"]["host"].as<std::string>();
        config.nvr.port = config_y["nvr"]["port"].as<int>();
        config.nvr.user = config_y["nvr"]["user"].as<std::string>();
        config.nvr.password = config_y["nvr"]["password"].as<std::string>();
    } catch (const YAML::Exception& e) {
        spdlog::critical(error(Error::FileReadFailed));
    }
}

config_s::~config_s()
{
}

// Error config::initConfig(std::string configPath, ServerConfig& serverConfig)
// {
//     std::ifstream fp(configPath);
//     if (!fp) {
//         return Error::FileOpenFailed;
//     }

//     YAML::Node config = YAML::Load(fp);
        
//     try {
//         serverConfig.name = config["name"].as<std::string>();
//         serverConfig.version = config["version"].as<std::string>();
//         serverConfig.port = config["port"].as<int>();
//         serverConfig.nvr.host = config["nvr"]["host"].as<std::string>();
//         serverConfig.nvr.port = config["nvr"]["port"].as<int>();
//         serverConfig.nvr.user = config["nvr"]["user"].as<std::string>();
//         serverConfig.nvr.password = config["nvr"]["password"].as<std::string>();
//     } catch (const YAML::Exception& e) {
//         return Error::FileReadFailed;
//     }
//     return Error::Nil;
// }

#include "config.h"

Config::Config(std::string path)
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

Config::~Config()
{
}

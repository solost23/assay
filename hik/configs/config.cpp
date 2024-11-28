#include "config.h"

Config::Config(std::string filepath)
{
    std::ifstream fp(filepath);
    if (!fp) {
        spdlog::critical(error(Error::FileOpenFailed));
    }

    YAML::Node config_y = YAML::Load(fp);

    try {
        name = config_y["name"].as<std::string>();
        version = config_y["version"].as<std::string>();
        port = config_y["port"].as<int>();
        nvr.host = config_y["nvr"]["host"].as<std::string>();
        nvr.port = config_y["nvr"]["port"].as<int>();
        nvr.user = config_y["nvr"]["user"].as<std::string>();
        nvr.password = config_y["nvr"]["password"].as<std::string>();
    } catch (const YAML::Exception& e) {
        spdlog::critical(error(Error::FileReadFailed));
    }
}

std::string Config::get_name() 
{
    return this->name;
}

std::string Config::get_version() 
{
    return this->version;
}

int Config::get_port()
{
    return this->port;
}

int Config::get_thread()
{
    return this->thread;
}

std::string Config::get_nvr_host()
{
    return this->nvr.host;
}

int Config::get_nvr_port()
{
    return this->nvr.port;
}

std::string Config::get_nvr_user() 
{
    return this->nvr.user;
}

std::string Config::get_nvr_password() 
{
    return this->nvr.password;
}

#include <iostream>

#include "controllers/nvr_c.h"

int main() {
    // 读取配置文件
    config_s* config = new config_s("./configs/config.yml");
    // if (Error err = config_s->initConfig("./configs/config.yml", serverConfig); err != Error::Nil) {
    //     // 日志打印
    //     spdlog::error(error(err));
    //     return err;
    // }

    // 开启http服务
    httplib::Server server;
    nvr_c* nvr = new nvr_c(config->config);
    server.Get("/api/hik/nvr/download", [nvr](const httplib::Request &request, httplib::Response& response) {
        nvr->download(request, response);
    });

    std::stringstream ss; ss << config->config.port;
    spdlog::info("server start: " + ss.str());
    server.listen("0.0.0.0", config->config.port, config->config.thread);
    return Error::Nil;
}

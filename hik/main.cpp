#include <iostream>

#include "controllers/nvr_c.h"

int main() {
    // 读取配置文件
    config_s config = config_s("./configs/config.yml");

    // 开启http服务
    httplib::Server server;
    // nvr_c* nvr_controller = new nvr_c(config.config);
    nvr_c nvr_controller = nvr_c(config.config);
    server.Get("/api/hik/nvr/download", [nvr_controller](const httplib::Request &request, httplib::Response& response) {
        nvr_controller.download(request, response);
    });

    std::stringstream ss; ss << config.config.port;
    spdlog::info("server start: " + ss.str());
    server.listen("0.0.0.0", config.config.port, config.config.thread);
    return Error::Nil;
}

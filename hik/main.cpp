#include <iostream>

#include "controllers/nvr.h"

int main() {
    // 读取配置文件
    Config config{"./configs/config.yml"};

    // // 开启http服务
    httplib::Server server;
    NvrController* nvr_controller = new NvrController(config);

    server.Get("/api/hik/nvr/channel", [nvr_controller](const httplib::Request& request, httplib::Response& response) {
        nvr_controller->channel(request, response);
    });
    server.Get("/api/hik/nvr/download", [nvr_controller](const httplib::Request& request, httplib::Response& response) {
        nvr_controller->download(request, response);
    });

    std::stringstream ss; ss << config.get_port();
    spdlog::info("server start: " + ss.str());
    server.listen("0.0.0.0", config.get_port(), config.get_thread());

    spdlog::info("server stop");
    return Error::Nil;
}

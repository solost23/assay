#include <iostream>

#include "controllers/nvr_c.h"

int main() {
    ServerConfig serverConfig{};
    // 读取配置文件
    config* config_s = new config();
    if (Error err = config_s->initConfig("./configs/config.yml", serverConfig); err != Error::Nil) {
        // 日志打印
        spdlog::error(error(err));
        return err;
    }

    // 开启http服务
    httplib::Server server;
    nvr_c* nvr = new nvr_c(serverConfig);
    server.Get("/api/hik/nvr/download", [nvr](const httplib::Request &request, httplib::Response& response) {
        nvr->download(request, response);
    });

    std::stringstream ss; ss << serverConfig.port;
    spdlog::info("server start: " + ss.str());
    server.listen("0.0.0.0", serverConfig.port, serverConfig.thread);
    return Error::Nil;
}

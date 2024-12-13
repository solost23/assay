#include <iostream>
#include <thread>
#include <chrono>
#include <atomic>
#include <csignal>
#include <sstream>

#include "controllers/nvr.h"

#define WEBCONFIGPATH "./configs/config.yml"

void signaler(int sig);
void start();

std::atomic<bool> running(true);

int main() 
{
    std::signal(SIGINT, signaler);
    std::signal(SIGTERM, signaler);

    std::thread http_thread(start);
    spdlog::info("HTTP server start. Press Ctrl+C to stop.");

    while(running) {
        if (!running) break;
        std::this_thread::sleep_for(std::chrono::milliseconds(200));
    }

    spdlog::info("HTTP server stoped.");

    return 0;
}

void signaler(int sig)
{
    running = false;

    std::ostringstream oss;
    oss << "Interrupt signal ("; oss << sig; oss << ") received";
    spdlog::info(oss.str());
}

void start()
{
    // 读取配置文件
    std::shared_ptr<Config> config = std::make_shared<Config>(WEBCONFIGPATH);

    // 开启http服务
    std::shared_ptr<httplib::Server> server = std::make_shared<httplib::Server>();

    std::shared_ptr<NvrController> nvr_controller = std::make_shared<NvrController>(config);

    server->Get("/api/hik/nvr/channel", [nvr_controller](const httplib::Request& request, httplib::Response& response) {
        nvr_controller->channel(request, response);
    });
    server->Get("/api/hik/nvr/download", [nvr_controller](const httplib::Request& request, httplib::Response& response) {
        nvr_controller->download(request, response);
    });

    std::stringstream ss; ss << config->get_port();
    spdlog::info("server start: " + ss.str());

    server->listen("0.0.0.0", config->get_port(), config->get_thread());

    spdlog::info("server stop");
}
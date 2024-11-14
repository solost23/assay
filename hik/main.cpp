#include <iostream>

#include "httplib.h"

#include "configs/config.h"
#include "services/nvr.h"

int parse(const httplib::Request&, int&, std::string);

int main() {
    // 读取配置文件
    config("./configs/config.yml");

    // 开启http服务
    httplib::Server server;
    server.Get("/api/hik/nvr/download", [](const httplib::Request& request, httplib::Response &response) {
        // 接收参数并打印
        DownloadForm params{};
        parse(request, params.channel, "channel");
        parse(request, params.startTime.year, "startYear");
        parse(request, params.startTime.month, "startMonth");
        parse(request, params.startTime.day, "startDay");
        parse(request, params.startTime.hour, "startHour");
        parse(request, params.startTime.minute, "startMinute");
        parse(request, params.startTime.second, "startSecond");

        parse(request, params.endTime.year, "endYear");
        parse(request, params.endTime.month, "endMonth");
        parse(request, params.endTime.day, "endDay");
        parse(request, params.endTime.hour, "endHour");
        parse(request, params.endTime.minute, "endMinute");
        parse(request, params.endTime.second, "endSecond");

        return nvrDownload(request, response, params);
    });

    std::cout << "server port: " << serverConfig.port << std::endl;
    server.listen("0.0.0.0", serverConfig.port);
    return 0;
}

int parse(const httplib::Request& request, int& value, std::string field) {
    if (!request.has_param(field)) {
        return -1;
    }
    value = atoi(request.get_param_value(field).c_str());
    return 0;
}

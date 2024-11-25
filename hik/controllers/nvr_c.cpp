#include "nvr_c.h"

nvr_c::nvr_c(Config config)
{
    this->config = config;
}

nvr_c::~nvr_c()
{
    spdlog::info("nvr_c object is being deleted");
}

void nvr_c::channel(const httplib::Request& request, httplib::Response& response) 
{
    nvr_s nvr_service = nvr_s(config);
    nvr_service.nvr_channel(request, response);
}

void nvr_c::download(const httplib::Request& request, httplib::Response& response)
{
    // 接收参数并打印
    DownloadForm params{};
    if (Error err = parse(request, params.channel, "channel"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": channel", "text/plain");
        return ;
    }
    if (Error err = parse(request, params.startTime.year, "startYear"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": startYear", "text/plain");
        return ;
    }
    if (Error err = parse(request, params.startTime.month, "startMonth"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": startMonth", "text/plain");
        return ;
    }
    if (Error err = parse(request, params.startTime.day, "startDay"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": startDay", "text/plain");
        return ;
    }
    if (Error err = parse(request, params.startTime.hour, "startHour"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": startHour", "text/plain");
        return ;
    }
    
    if (Error err = parse(request, params.startTime.minute, "startMinute"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": startMinute", "text/plain");
        return ;
    }
    
    if (Error err = parse(request, params.startTime.second, "startSecond"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": startSecond", "text/plain");
        return ;
    }
    
    if (Error err = parse(request, params.endTime.year, "endYear"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": endYear", "text/plain");
        return ;
    }
    
    if (Error err = parse(request, params.endTime.month, "endMonth"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": endMonth", "text/plain");
        return ;
    }
    
    if (Error err = parse(request, params.endTime.day, "endDay"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": endDay", "text/plain");
        return ;
    }
    
    if (Error err = parse(request, params.endTime.hour, "endHour"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": endHour", "text/plain");
        return ;
    }
    
    if (Error err = parse(request, params.endTime.minute, "endMinute"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": endMinute", "text/plain");
        return ;
    }
    
    if (Error err = parse(request, params.endTime.second, "endSecond"); err != Error::Nil) {
        response.status = 200;
        response.set_content(error(err) + ": endSecond", "text/plain");
        return ;
    }

    nvr_s nvr_service = nvr_s(config);
    nvr_service.nvr_download(request, response, params);
    return;
}

Error nvr_c::parse(const httplib::Request& request, int& value, std::string field) 
{
    if (!request.has_param(field)) {
        return Error::BadRequest;
    }
    value = atoi(request.get_param_value(field).c_str());
    return Error::Nil;
}
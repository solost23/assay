#include "nvr.h"

NvrController::~NvrController()
{
    spdlog::info("nvr_c object is being deleted");
}

void NvrController::channel(const httplib::Request& request, httplib::Response& response) 
{
    std::shared_ptr<NvrService> nvr_service = std::make_shared<NvrService>(this->config);
    nvr_service->nvr_channel(request, response);
}

void NvrController::download(const httplib::Request& request, httplib::Response& response)
{
    // 接收参数并打印
    DownloadForm params{};
    if (!parse(request, params.channel, "channel") 
    && !parse(request, params.channel, "startYear") && !parse(request, params.startTime.month, "startMonth") 
    && !parse(request, params.startTime.day, "startDay")&& !parse(request, params.startTime.hour, "startHour") 
    && !parse(request, params.startTime.minute, "startMinute") && !parse(request, params.startTime.second, "startSecond") 
    && !parse(request, params.endTime.year, "endYear") && !parse(request, params.endTime.month, "endMonth") 
    && !parse(request, params.endTime.day, "endDay") && !parse(request, params.endTime.hour, "endHour") 
    && !parse(request, params.endTime.minute, "endMinute") && !parse(request, params.endTime.second, "endSecond")) 
    {
        response.status = 200;
        response.set_content("params valid", "text/plain");
        return ;
    }

    std::shared_ptr<NvrService> nvr_service = std::make_shared<NvrService>(this->config);
    nvr_service->nvr_download(request, response, params);
    return;
}

inline Error NvrController::parse(const httplib::Request& request, int& value, std::string field) 
{
    if (!request.has_param(field)) {
        return Error::BadRequest;
    }
    value = atoi(request.get_param_value(field).c_str());
    return Error::Nil;
}
#include "nvr_c.h"


// NvrController::NvrController(ServerConfig config):serverConfig(config)
// {

// }
//     NvrController::~NvrController()
//     {

//     }
//     // 视频下载
//     void NvrController::download(const httplib::Request& request, httplib::Response& response) 
//     {
//         // 接收参数并打印
//         DownloadForm params{};
//         if (Error err = parse(request, params.channel, "channel"); err != Error::Nil) {
            
//         }
//         if (Error err = parse(request, params.startTime.year, "startYear"); err != Error::Nil) {
            
//         }
//         if (Error err = parse(request, params.startTime.month, "startMonth"); err != Error::Nil) {
            
//         }
//         if (Error err = parse(request, params.startTime.day, "startDay"); err != Error::Nil) {
            
//         }
//         if (Error err = parse(request, params.startTime.hour, "startHour"); err != Error::Nil) {
            
//         }
        
//         if (Error err = parse(request, params.startTime.minute, "startMinute"); err != Error::Nil) {
            
//         }
        
//         if (Error err = parse(request, params.startTime.second, "startSecond"); err != Error::Nil) {
            
//         }
        
//         if (Error err = parse(request, params.endTime.year, "endYear"); err != Error::Nil) {
            
//         }
        
//         if (Error err = parse(request, params.endTime.month, "endMonth"); err != Error::Nil) {
            
//         }
        
//         if (Error err = parse(request, params.endTime.day, "endDay"); err != Error::Nil) {
            
//         }
        
//         if (Error err = parse(request, params.endTime.hour, "endHour"); err != Error::Nil) {
            
//         }
        
//         if (Error err = parse(request, params.endTime.minute, "endMinute"); err != Error::Nil) {
            
//         }
        
//         if (Error err = parse(request, params.endTime.second, "endSecond"); err != Error::Nil) {
            
//         }

//         nvrDownload(request, response, serverConfig, params);
        
//         return ;
//     };


// Error NvrController::parse(const httplib::Request& request, int& value, std::string field) 
// {
//     if (!request.has_param(field)) {
//         return Error::ParamParseFailed;
//     }
//     value = atoi(request.get_param_value(field).c_str());
//     return Error::Nil;
// }


nvr_c::nvr_c(ServerConfig config)
{
    serverConfig = config;
}

nvr_c::~nvr_c()
{
}

void nvr_c::download(const httplib::Request& request, httplib::Response& response)
{
    // 接收参数并打印
    DownloadForm params{};
    if (Error err = parse(request, params.channel, "channel"); err != Error::Nil) {
        
    }
    if (Error err = parse(request, params.startTime.year, "startYear"); err != Error::Nil) {
        
    }
    if (Error err = parse(request, params.startTime.month, "startMonth"); err != Error::Nil) {
        
    }
    if (Error err = parse(request, params.startTime.day, "startDay"); err != Error::Nil) {
        
    }
    if (Error err = parse(request, params.startTime.hour, "startHour"); err != Error::Nil) {
        
    }
    
    if (Error err = parse(request, params.startTime.minute, "startMinute"); err != Error::Nil) {
        
    }
    
    if (Error err = parse(request, params.startTime.second, "startSecond"); err != Error::Nil) {
        
    }
    
    if (Error err = parse(request, params.endTime.year, "endYear"); err != Error::Nil) {
        
    }
    
    if (Error err = parse(request, params.endTime.month, "endMonth"); err != Error::Nil) {
        
    }
    
    if (Error err = parse(request, params.endTime.day, "endDay"); err != Error::Nil) {
        
    }
    
    if (Error err = parse(request, params.endTime.hour, "endHour"); err != Error::Nil) {
        
    }
    
    if (Error err = parse(request, params.endTime.minute, "endMinute"); err != Error::Nil) {
        
    }
    
    if (Error err = parse(request, params.endTime.second, "endSecond"); err != Error::Nil) {
        
    }

    nvr_s* nvrService = new nvr_s(serverConfig);
    nvrService->nvrDownload(request, response, params);
    return ;
}

Error nvr_c::parse(const httplib::Request& request, int& value, std::string field)
{
    if (!request.has_param(field)) {
        return Error::ParamParseFailed;
    }
    value = atoi(request.get_param_value(field).c_str());
    return Error::Nil;
}
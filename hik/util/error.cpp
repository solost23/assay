#include <iostream>
#include "error.h"

std::string error(Error err) {
    switch (err) {
        case Error::Nil:
            return "nil";
        case Error::FileOpenFailed:
            return "file open failed";
        case Error::FileReadFailed:
            return "file read failed";
        case Error::UserLoginFailed:
            return "user login failed";
        case Error::DvrGetFileByTimeV40Failed:
            return "dvr_get_file_by_time_v40 failed";
        case Error::PlaybackControlFailed:
            return "playback control failed";
        case Error::FileStopGetFailed:
            return "file stop get failed";
        case Error::FileDownloadFailed:
            return "file download failed";
        case Error::DvrGetConfigFailed:
            return "dvr_get_config failed";
        case Error::IntervalServerFailed:
            return "interval server failed";
        case Error::BadRequest:
            return "bad request";
        case Error::NotFound:
            return "resource not found";
    }
    return "";
}

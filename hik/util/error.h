#pragma once

enum Error {
    Nil = 0, 
    FileOpenFailed = -1,
    FileReadFailed = -2, 
    UserLoginFailed = -4, 
    DvrGetFileByTimeV40Failed = -5, 
    PlaybackControlFailed = -6, 
    FileStopGetFailed = -7,
    FileDownloadFailed = -8, 
    IntervalServerFailed = 1500, 
    BadRequest = 1400,
    NotFound = 404, 
};

std::string error(Error);

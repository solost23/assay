struct TimeForm {
    int year;
    int month;
    int day;
    int hour;
    int minute;
    int second;
};

struct DownloadForm {
    int channel;
    TimeForm startTime;
    TimeForm endTime;
};

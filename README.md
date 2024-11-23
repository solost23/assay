## assay

### 启动推流服务
```bash
sudo apt-get install ffmpeg

# 代码中实现这个命令
ffmpeg -i rtsp://admin:asdfghj1@192.168.1.70:554/Streaming/Channels/101 -c:v libx264 -preset veryfast -tune zerolatency -f flv rtmp://localhost:1935/live/stream-70

ffmpeg -i rtsp://admin:asdfghj1@192.168.1.111:554/Streaming/Channels/101 -c:v libx264 -preset veryfast -tune zerolatency -f flv rtmp://localhost:1935/live/stream-111
```

上述方式延迟较高，下面考虑 webrtc 技术

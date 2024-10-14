FROM golang:alpine AS builder 

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    GOPROXY=https://goproxy.cn/

WORKDIR /build 
COPY . . 

RUN go build -o assay ./main.go 
RUN apk add tzdata

FROM scratch 

COPY --from=builder /build/assay / 
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

EXPOSE 8083 
CMD [ "/assay" ]
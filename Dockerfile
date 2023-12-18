FROM golang:alpine as builder

# 设置工作目录
WORKDIR /build

# 配置环境变量
ENV GOPROXY=https://goproxy.cn,direct
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .
RUN go mod tidy
RUN go build -o main main.go

FROM alpine as prod

WORKDIR /build

COPY --from=builder /build/main .
COPY /config/config.yaml /build/config/

CMD ["./main"]
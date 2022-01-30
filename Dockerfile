# syntax = docker/dockerfile:experimental

FROM swr.cn-east-3.myhuaweicloud.com/jackrush/golang:latest as builder
ARG APP_NAME
WORKDIR /build
copy . .
RUN --mount=type=cache,target=/go,id=gomod,sharing=locked \
    GOPROXY=https://goproxy.cn CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -v -tags timetzdata -o ${APP_NAME}

FROM swr.cn-east-3.myhuaweicloud.com/jackrush/alpine:latest
ARG APP_NAME
ENV APP_PATH=/${APP_NAME}
COPY --from=builder /build/${APP_NAME} /
CMD $APP_PATH

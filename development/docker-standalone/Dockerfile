FROM toolsetlink/build-upgradelink:v2.0.1 AS builder

WORKDIR /go/src/upgradelink-standalone/

COPY . .

# 定义构建参数（由 Buildx 自动注入，对应目标架构）
ARG TARGETOS
ARG TARGETARCH

# 为每个子项目单独处理依赖（根据项目实际目录调整）
# 交叉编译：根据目标架构编译二进制文件
# CGO_ENABLED=0 禁用 CGO，确保静态编译（避免依赖系统库）

# 1. upgradelink-admin-core 项目
WORKDIR /go/src/upgradelink-standalone/upgradelink-admin-core/server
RUN go mod tidy \
    && go mod download

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/src/upgradelink-standalone/upgradelink-admin-core/server/api/core-api /go/src/upgradelink-standalone/upgradelink-admin-core/server/api/core.go
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/src/upgradelink-standalone/upgradelink-admin-core/server/rpc/core-rpc /go/src/upgradelink-standalone/upgradelink-admin-core/server/rpc/core.go

# 2. upgradelink-admin-file 项目
WORKDIR /go/src/upgradelink-standalone/upgradelink-admin-file
RUN go mod tidy \
    && go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/src/upgradelink-standalone/upgradelink-admin-file/server/file /go/src/upgradelink-standalone/upgradelink-admin-file/server/fms.go

# 3. upgradelink-admin-message 项目
WORKDIR /go/src/upgradelink-standalone/upgradelink-admin-message
RUN go mod tidy \
    && go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/src/upgradelink-standalone/upgradelink-admin-message/server/message /go/src/upgradelink-standalone/upgradelink-admin-message/server/mcms.go

# 4. upgradelink-admin-upgrade 项目
WORKDIR /go/src/upgradelink-standalone/upgradelink-admin-upgrade
RUN go mod tidy \
    && go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/src/upgradelink-standalone/upgradelink-admin-upgrade/server/upgrade /go/src/upgradelink-standalone/upgradelink-admin-upgrade/server/upgrade.go

# 5. upgradelink-api 项目
WORKDIR /go/src/upgradelink-standalone/upgradelink-api
RUN go mod tidy \
    && go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/src/upgradelink-standalone/upgradelink-api/server/api/api /go/src/upgradelink-standalone/upgradelink-api/server/api/api.go



FROM toolsetlink/alpine-upgradelink:v2.0.0 AS prod

LABEL description="upgradelink-standalone服务"

ARG ENV=prod

ENV RUN_MODE=$ENV APP_NAME="upgradelink-standalone"

WORKDIR /link/upgradelink-standalone

RUN mkdir -p /link/upgradelink-standalone/etc /link/upgradelink-standalone/build

COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin-core/server/api/core-api  /link/upgradelink-standalone/
COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin-core/server/api/etc/core-api.yaml  /link/upgradelink-standalone/etc/

COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin-core/server/rpc/core-rpc  /link/upgradelink-standalone/
COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin-core/server/rpc/etc/core-rpc.yaml  /link/upgradelink-standalone/etc/

COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin-file/server/file  /link/upgradelink-standalone/
COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin-file/server/etc/fms.yaml  /link/upgradelink-standalone/etc/

COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin-message/server/message  /link/upgradelink-standalone/
COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin-message/server/etc/mcms.yaml  /link/upgradelink-standalone/etc/

COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin-upgrade/server/upgrade  /link/upgradelink-standalone/
COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin-upgrade/server/etc/upgrade.yaml  /link/upgradelink-standalone/etc/

COPY --from=builder /go/src/upgradelink-standalone/upgradelink-api/server/api/api  /link/upgradelink-standalone/
COPY --from=builder /go/src/upgradelink-standalone/upgradelink-api/server/api/etc/api.yaml  /link/upgradelink-standalone/etc/


# 创建 Nginx html 目录
RUN mkdir -p /usr/share/nginx/html
COPY upgradelink-admin-ui/apps/simple-admin-core/dist /usr/share/nginx/html/

# 拷贝配置文件
COPY development/docker-standalone/nginx.conf /etc/nginx/
COPY development/docker-standalone/logrotate-nginx /link/upgradelink-standalone/logrotate-nginx
COPY development/docker-standalone/supervisord.conf /link/upgradelink-standalone/supervisord.conf


# 创建启动脚本和设置定时任务，修正路径引用
RUN sed -i '$a\0 0 * * * logrotate -f /link/upgradelink-standalone/logrotate-nginx' /var/spool/cron/crontabs/root \
    && echo '''mkdir -p /link/logs/nginx && \
        mkdir -p /link/logs/supervisor && \
        supervisord -c /link/upgradelink-standalone/supervisord.conf''' > start.sh \
    && chmod a+x start.sh

EXPOSE 8080

CMD [ "/bin/bash", "start.sh" ]


# 在当前目录执行
# docker buildx build -t toolsetlink/upgradelink-standalone:v2.0.6 --platform=linux/amd64,linux/arm64 -f ./Dockerfile ../../ --push


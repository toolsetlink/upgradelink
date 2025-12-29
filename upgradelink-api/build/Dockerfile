FROM toolsetlink/build-upgradelink:v2.0.2 AS builder

WORKDIR /go/src/upgradelink-api/

COPY . .

# 定义构建参数（由 Buildx 自动注入，对应目标架构）
ARG TARGETOS
ARG TARGETARCH

RUN go mod tidy \
    && go mod download

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/src/upgradelink-api/server/api/api /go/src/upgradelink-api/server/api/api.go


# 缩小包体积
RUN tar zcvf /go/src/upgradelink-api/server/api.tar.gz -C /go/src/upgradelink-api/server api/api api/etc


FROM toolsetlink/alpine-upgradelink:v2.0.0 AS prod

LABEL description="upgradelink-api"

ARG ENV=prod

ENV RUN_MODE=$ENV APP_NAME="upgradelink-api"

WORKDIR /link/upgradelink-api

COPY --from=builder /go/src/upgradelink-api/server/api.tar.gz  .

# 拷贝 Nginx 配置文件和构建好的文件到 Nginx 的 html 目录
COPY --from=builder /go/src/upgradelink-api/build/nginx.conf /etc/nginx/
COPY --from=builder /go/src/upgradelink-api/build/logrotate-nginx /link/upgradelink-api/build/logrotate-nginx
COPY --from=builder /go/src/upgradelink-api/build/supervisord.conf /link/upgradelink-api/build/supervisord.conf

RUN tar -zxvf api.tar.gz \
    && rm -f api.tar.gz \
    && sed -i '$a\0 0 * * * logrotate -f /link/upgradelink-api/build/logrotate-nginx' /var/spool/cron/crontabs/root \
    && echo '''mkdir -p /link/logs/nginx && \
        mkdir -p /link/logs/supervisor && \
        supervisord -c /link/upgradelink-api/build/supervisord.conf''' > start.sh \
    && chmod a+x start.sh

EXPOSE 80

CMD [ "/bin/bash", "start.sh" ]

# 在当面模块根目录下执行  /upgradelink/upgradelink-api

# amd架构
# docker buildx build --no-cache -t upgradelink-api:2.4.0 --platform=linux/amd64 -f build/Dockerfile . --load
# arm架构
# docker buildx build --no-cache -t upgradelink-api:2.4.0 --platform=linux/arm64 -f build/Dockerfile . --load


# 打包开源镜像
# docker buildx build -t toolsetlink/upgradelink-api:2.4.0 --platform=linux/amd64,linux/arm64 -f build/Dockerfile . --push

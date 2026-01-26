FROM toolsetlink/build-upgradelink:v2.0.1 AS builder

WORKDIR /go/src/upgradelink-admin/

COPY . .

# 定义构建参数（由 Buildx 自动注入，对应目标架构）
ARG TARGETOS
ARG TARGETARCH

RUN go mod tidy \
    && go mod download

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/src/upgradelink-admin/server/api/admin /go/src/upgradelink-admin/server/api/admin.go

# 缩小包体积
RUN tar zcvf /go/src/upgradelink-admin/server/admin.tar.gz -C /go/src/upgradelink-admin/server api/admin api/etc

FROM toolsetlink/alpine-upgradelink:v2.0.0 AS prod

LABEL description="upgradelink-admin"

ARG ENV=prod

ENV RUN_MODE=$ENV APP_NAME="upgradelink-admin"

WORKDIR /link/upgradelink-admin

COPY --from=builder /go/src/upgradelink-admin/server/admin.tar.gz  .

# 拷贝 HdiffPatch 文件
# 根据不同架构选择对应的 HDiffPatch 文件路径
ARG TARGETARCH
RUN --mount=type=bind,from=builder,source=/go/src/upgradelink-admin/build/HDiffPatch,target=/tmp/hdiffpatch \
    if [ "$TARGETARCH" = "arm64" ]; then \
        cp /tmp/hdiffpatch/linux_arm64/hdiffz /usr/local/bin/hdiffz; \
    else \
        cp /tmp/hdiffpatch/linux64/hdiffz /usr/local/bin/hdiffz; \
    fi && \
    chown root:root /usr/local/bin/hdiffz

COPY --from=builder /go/src/upgradelink-admin/build/nginx.conf /etc/nginx/
COPY --from=builder /go/src/upgradelink-admin/build/logrotate-nginx /link/upgradelink-admin/build/logrotate-nginx
COPY --from=builder /go/src/upgradelink-admin/build/supervisord.conf /link/upgradelink-admin/build/supervisord.conf

RUN tar -zxvf admin.tar.gz \
    && rm -f admin.tar.gz \
    && sed -i '$a\0 0 * * * logrotate -f /link/upgradelink-admin/build/logrotate-nginx' /var/spool/cron/crontabs/root \
    && echo '''mkdir -p /link/logs/nginx && \
        mkdir -p /link/logs/supervisor && \
        supervisord -c /link/upgradelink-admin/build/supervisord.conf''' > start.sh \
    && chmod a+x start.sh

EXPOSE 80

CMD [ "/bin/bash", "start.sh" ]

# 在当面模块根目录下执行  /upgradelink/upgradelink-admin

# amd架构
# docker buildx build  -t upgradelink-admin:2.9.0 --platform=linux/amd64 -f build/Dockerfile-api . --load
# arm架构
# docker buildx build  -t upgradelink-admin:2.9.0 --platform=linux/arm64 -f build/Dockerfile-api . --load

# 打包开源镜像
# docker buildx build -t toolsetlink/upgradelink-admin:2.9.1 --platform=linux/amd64,linux/arm64 -f build/Dockerfile . --push


FROM toolsetlink/build-upgradelink:v2.0.2 AS builder

WORKDIR /go/src/upgradelink-standalone/

COPY . .

# 定义构建参数（由 Buildx 自动注入，对应目标架构）
ARG TARGETOS
ARG TARGETARCH

# 为每个子项目单独处理依赖并编译（保持原有逻辑）
# 1. upgradelink-admin 项目
WORKDIR /go/src/upgradelink-standalone/upgradelink-admin
RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/src/upgradelink-standalone/upgradelink-admin/server/api/admin /go/src/upgradelink-standalone/upgradelink-admin/server/api/admin.go

# 2. upgradelink-api 项目
WORKDIR /go/src/upgradelink-standalone/upgradelink-api
RUN go mod tidy && go mod download
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /go/src/upgradelink-standalone/upgradelink-api/server/api/api /go/src/upgradelink-standalone/upgradelink-api/server/api/api.go


FROM toolsetlink/alpine-upgradelink:v2.0.0 AS prod

LABEL description="upgradelink-standalone服务"

ARG ENV=prod
ENV RUN_MODE=$ENV APP_NAME="upgradelink-standalone"

WORKDIR /link/upgradelink-standalone

# 创建必要目录
RUN mkdir -p /link/upgradelink-standalone/etc /link/upgradelink-standalone/build \
    /usr/share/nginx/html /link/logs/nginx /link/logs/supervisor

# 复制二进制文件和使用的配置文件（仅保留development/docker-standalone下的配置）
COPY --from=builder /go/src/upgradelink-standalone/upgradelink-admin/server/api/admin  ./
COPY --from=builder /go/src/upgradelink-standalone/development/docker-standalone/upgradelink-admin/admin.yaml  ./etc

COPY --from=builder /go/src/upgradelink-standalone/upgradelink-api/server/api/api  ./
COPY --from=builder /go/src/upgradelink-standalone/development/docker-standalone/upgradelink-api/api.yaml  ./etc

# 拷贝 task项目所需 HdiffPatch 文件
# 根据不同架构选择对应的 HDiffPatch 文件路径
ARG TARGETARCH
RUN --mount=type=bind,from=builder,source=/go/src/upgradelink-standalone/upgradelink-admin/build/HDiffPatch,target=/tmp/hdiffpatch \
    if [ "$TARGETARCH" = "arm64" ]; then \
        cp /tmp/hdiffpatch/linux_arm64/hdiffz /usr/local/bin/hdiffz; \
    else \
        cp /tmp/hdiffpatch/linux64/hdiffz /usr/local/bin/hdiffz; \
    fi && \
    chown root:root /usr/local/bin/hdiffz 

# 复制前端文件和服务配置
COPY upgradelink-admin-ui/apps/admin/dist /usr/share/nginx/html/
COPY development/docker-standalone/nginx.conf /etc/nginx/
COPY development/docker-standalone/logrotate-nginx ./
COPY development/docker-standalone/supervisord.conf ./

# 关键优化：删除未使用的配置文件（根据实际项目中冗余配置的路径调整）
# 假设项目中存在以下未使用的配置文件路径（例如子项目自带的默认配置）
RUN rm -rf \
    /go/src/upgradelink-standalone/upgradelink-admin/server/etc \
    /go/src/upgradelink-standalone/upgradelink-api/server/etc \
    # 如需删除其他冗余配置
    # /go/src/upgradelink-standalone/configs \
    # /go/src/upgradelink-standalone/local-conf \
    # 清理复制过程中可能带入的空目录或临时文件
    && find . -type d -empty -delete


# 设置定时任务和启动脚本
RUN sed -i '$a\0 0 * * * logrotate -f /link/upgradelink-standalone/logrotate-nginx' /var/spool/cron/crontabs/root \
    && echo '''supervisord -c /link/upgradelink-standalone/supervisord.conf''' > start.sh \
    && chmod a+x start.sh

EXPOSE 80 8888

CMD [ "/bin/bash", "start.sh" ]

# 在当前目录执行
# docker buildx build -t toolsetlink/upgradelink-standalone:3.0.0 --platform=linux/amd64,linux/arm64 -f ./Dockerfile ../../ --push


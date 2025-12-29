FROM toolsetlink/alpine-upgradelink:v2.0.0 AS prod

# 设置工作目录
WORKDIR /app

# 创建 Nginx html 目录
RUN mkdir -p /usr/share/nginx/html

COPY build/nginx.conf /etc/nginx/nginx.conf
COPY build/logrotate-nginx /app/build/
COPY build/supervisord.conf /app/build/
COPY apps/simple-admin-core/dist /usr/share/nginx/html/

# 添加并配置启动脚本
RUN echo "0 0 * * * logrotate -f /app/build/logrotate-nginx" >> /var/spool/cron/crontabs/root \
    && echo "mkdir -p /link/logs/nginx && \
     mkdir -p /link/logs/supervisor && \
     supervisord -c /app/build/supervisord.conf" > start.sh \
    && chmod a+x start.sh

EXPOSE 80

CMD [ "/bin/bash", "start.sh"]

# 在当面模块根目录下执行  /upgradelink/upgradelink-admin-ui

# amd架构
# docker buildx build --no-cache -t upgradelink-admin-ui:2.3.1 --platform=linux/amd64 -f build/Dockerfile . --load
# arm架构
# docker buildx build --no-cache -t upgradelink-admin-ui:2.3.1 --platform=linux/arm64 -f build/Dockerfile . --load

# 构建命令
# pnpm install
# pnpm build:simple-admin-core
# 打包开源镜像
# docker buildx build -t toolsetlink/upgradelink-admin-ui:2.4.0 --platform=linux/amd64,linux/arm64 -f build/Dockerfile . --push



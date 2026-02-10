## 使用构建镜像
FROM docker.m.daocloud.io/library/node:22.12.0 AS builder

RUN npm config set -g registry http://registry.npmmirror.com

ENV TZ=Asia/Shanghai

ARG registry=http://registry.npmmirror.com
ARG disturl=https://npm.taobao.org/dist

RUN yarn config set disturl $disturl
RUN yarn config set registry $registry

## 设置工作目录
WORKDIR /app

## 复制项目文件到工作目录
COPY . .

## 安装依赖
RUN npm i -g corepack && CI=true pnpm install

## 构建项目
RUN pnpm build:admin

# 使用生产镜像
FROM toolsetlink/alpine-upgradelink:v2.0.0 AS prod

# 设置工作目录
WORKDIR /app

# 创建 Nginx html 目录
RUN mkdir -p /usr/share/nginx/html

COPY build/nginx.conf /etc/nginx/nginx.conf
COPY build/logrotate-nginx /app/build/
COPY build/supervisord.conf /app/build/
COPY apps/admin/dist /usr/share/nginx/html/

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
# docker buildx build --no-cache -t upgradelink-admin-ui:3.0.0 --platform=linux/amd64 -f build/Dockerfile . --load
# arm架构
# docker buildx build --no-cache -t upgradelink-admin-ui:3.0.0 --platform=linux/arm64 -f build/Dockerfile . --load

# 构建命令
# pnpm install
# pnpm build:admin
# 打包开源镜像
# docker buildx build -t toolsetlink/upgradelink-admin-ui:3.0.0 --platform=linux/amd64,linux/arm64 -f build/Dockerfile . --push



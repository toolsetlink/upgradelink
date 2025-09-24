FROM nginx:1.27.0-alpine

COPY ./apps/simple-admin-core/dist/ /usr/share/nginx/html/
COPY ./scripts/deploy/nginx.conf /etc/nginx/nginx.conf

LABEL MAINTAINER="toolsetlink@163.com"

EXPOSE 80

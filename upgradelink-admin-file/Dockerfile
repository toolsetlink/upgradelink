FROM nginx:1.27.0-alpine

# Define the project name | 定义项目名称
ARG PROJECT=fms
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=fms.yaml
# Define the author | 定义作者
ARG AUTHOR="yuansu.china.work@gmail.com"

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

ENV TZ=Asia/Shanghai
RUN apk update --no-cache && apk add --no-cache tzdata

COPY ./${PROJECT}_api ./
COPY ./etc/${CONFIG_FILE} ./etc/
COPY deploy/nginx/default.conf /etc/nginx/conf.d/
COPY deploy/nginx/entrypoint.sh /docker-entrypoint.d

RUN ["chmod", "+x", "/docker-entrypoint.d/entrypoint.sh"]

EXPOSE 80
EXPOSE 9102


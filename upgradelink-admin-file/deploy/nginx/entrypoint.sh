#!/bin/sh

nginx -c  /etc/nginx/nginx.conf &

/app/fms_api -f /app/etc/fms.yaml
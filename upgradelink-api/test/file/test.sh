#!/bin/bash

export LANG=en_US.UTF-8

# 常量定义
accessKey="mui2W50H1j-OC4xD6PgQag"
secretKey="PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc"
#apiURL="http://api.upgrade.toolsetlink.com/v1/url/upgrade"
apiURL="http://0.0.0.0:8888/v1/url/upgrade"
uri="/v1/url/upgrade"

# 生成 RFC3339 格式时间的函数
generate_rfc3339_time() {
    if [[ "$(uname)" == "Darwin" ]]; then
        # macOS 系统处理方式
        # 先获取当前时间，精确到秒
        local base_time=$(date +"%Y-%m-%dT%H:%M:%S")
        # 拼接成完整的 RFC3339 格式时间
        echo "${base_time}+08:00"
    else
        # Linux 系统处理方式
        # 使用 date 命令直接生成 RFC3339 格式时间
        date +"%Y-%m-%dT%H:%M:%S+08:00"
    fi
}


# 1. 构造请求体
requestBody='{
    "fileKey": "LOYlLXNy7wV3ySuh0XgtSg",
    "versionCode": 1,
    "appointVersionCode": 0,
    "devModelKey": "12312",
    "devKey": "12312"
}'

# 2. 生成请求参数
timestamp=$(generate_rfc3339_time)
echo "生成的 RFC3339 格式时间: $timestamp"
nonce=$(head -c 8 /dev/urandom | xxd -p)

# 3. 生成签名
signStr="body=$requestBody&nonce=$nonce&secretKey=$secretKey&timestamp=$timestamp&url=$uri"
signature=$(printf "%s" "$signStr" | md5sum | awk '{print $1}')

# 4. 打印时间戳
echo "timestamp: $timestamp"

# 5. 发送 HTTP 请求
response=$(curl -s -X POST \
  -H "X-Timestamp: $timestamp" \
  -H "X-Nonce: $nonce" \
  -H "X-AccessKey: $accessKey" \
  -H "X-Signature: $signature" \
  -H "Content-Type: application/json" \
  -d "$requestBody" \
  "$apiURL")

# 6. 获取响应状态码
statusCode=$(curl -s -o /dev/null -w "%{http_code}" -X POST \
  -H "X-Timestamp: $timestamp" \
  -H "X-Nonce: $nonce" \
  -H "X-AccessKey: $accessKey" \
  -H "X-Signature: $signature" \
  -H "Content-Type: application/json" \
  -d "$requestBody" \
  "$apiURL")

# 7. 处理响应
echo "Status Code: $statusCode"
echo "Response: $response"
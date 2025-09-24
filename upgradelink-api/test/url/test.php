<?php
// 常量定义
$accessKey = "mui2W50H1j-OC4xD6PgQag";
$secretKey = "PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc";
$apiURL = "http://api.upgrade.toolsetlink.com/v1/url/upgrade";
$uri = "/v1/url/upgrade";

// 生成 RFC3339Nano 格式时间的函数
function generate_rfc3339nano_time() {
    $base_time = date("Y-m-d\TH:i:s");
    $microseconds = microtime(true) - floor(microtime(true));
    $nanos = sprintf("%06d", $microseconds * 1000000);
    return $base_time . "." . $nanos . "+08:00";
}

// 1. 构造请求体
$requestBody = '{
    "urlKey": "SkEgKQ4SyLmzazl31fJnAw",
    "versionCode": 1,
    "appointVersionCode": 0,
    "devModelKey": "stv1",
    "devKey": "LOYlLXNy7w"
}';

// 2. 生成请求参数
$timestamp = generate_rfc3339nano_time();
echo "生成的 RFC3339Nano 格式时间: $timestamp\n";
$nonce = bin2hex(random_bytes(8));

// 3. 生成签名
$signStr = "body=$requestBody&nonce=$nonce&secretKey=$secretKey&timestamp=$timestamp&url=$uri";
$signature = md5($signStr);

// 4. 打印时间戳
echo "timestamp: $timestamp\n";

// 5. 发送 HTTP 请求
$ch = curl_init();
curl_setopt($ch, CURLOPT_URL, $apiURL);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_setopt($ch, CURLOPT_POST, true);
curl_setopt($ch, CURLOPT_POSTFIELDS, $requestBody);
curl_setopt($ch, CURLOPT_HTTPHEADER, [
    "X-Timestamp: $timestamp",
    "X-Nonce: $nonce",
    "X-AccessKey: $accessKey",
    "X-Signature: $signature",
    "Content-Type: application/json"
]);
$response = curl_exec($ch);

// 6. 获取响应状态码
$statusCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

// 7. 处理响应
echo "Status Code: $statusCode\n";
echo "Response: $response\n";
?>
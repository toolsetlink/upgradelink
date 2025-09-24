package main

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type UpgradeRequest struct {
	FileKey            string `json:"fileKey"`
	VersionCode        int    `json:"versionCode"`
	AppointVersionCode int    `json:"appointVersionCode,omitempty"`
	DevModelKey        string `json:"devModelKey,omitempty"`
	DevKey             string `json:"devKey,omitempty"`
}

type UpgradeResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	accessKey = "mui2W50H1j-OC4xD6PgQag"                      // 替换为实际AccessKey
	secretKey = "PEbdHFGC0uO_Pch7XWBQTMsFRxKPQAM2565eP8LJ3gc" // 替换为实际SecretKey
	//apiURL    = "http://api.upgrade.toolsetlink.com/v1/file/upgrade"
	apiURL = "http://0.0.0.0:8888/v1/file/upgrade"
)

func main() {
	// 1. 构造请求体
	requestBody := UpgradeRequest{
		FileKey:            "LOYlLXNy7wV3ySuh0XgtSg",
		VersionCode:        1,
		AppointVersionCode: 0,
		DevModelKey:        "stv1",
		DevKey:             "asdasdsadsa",
	}

	// 2. 序列化请求体
	bodyBytes, _ := json.Marshal(requestBody)
	bodyStr := string(bodyBytes)

	// 3. 生成请求参数
	timestamp := time.Now().Format(time.RFC3339)
	nonce := generateNonce()
	fmt.Println("nonce: ", nonce)
	uri := "/v1/file/upgrade"

	// 4. 生成签名
	signature := generateSignature(bodyStr, nonce, secretKey, timestamp, uri)
	fmt.Println("signature: ", signature)

	// 5. 创建HTTP请求
	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiURL, bytes.NewBuffer(bodyBytes))

	fmt.Println("timestamp: ", timestamp)

	// 6. 设置请求头
	req.Header.Set("X-Timestamp", timestamp)
	req.Header.Set("X-Nonce", nonce)
	req.Header.Set("X-AccessKey", accessKey)
	req.Header.Set("X-Signature", signature)
	req.Header.Set("Content-Type", "application/json")

	// 7. 发送请求
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 8. 处理响应
	respBody, _ := io.ReadAll(resp.Body)
	var result UpgradeResponse
	json.Unmarshal(respBody, &result)

	fmt.Printf("Status Code: %d\n", resp.StatusCode)
	fmt.Printf("Response: %+v\n", result)
}

// 生成随机Nonce (16位十六进制)
func generateNonce() string {
	bytes := make([]byte, 8)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

// 生成请求签名
func generateSignature(body, nonce, secretKey, timestamp, uri string) string {
	var parts []string

	if body != "" {
		parts = append(parts, "body="+body)
	}

	parts = append(parts,
		"nonce="+nonce,
		"secretKey="+secretKey,
		"timestamp="+timestamp,
		"url="+uri,
	)

	signStr := strings.Join(parts, "&")
	hash := md5.Sum([]byte(signStr))
	return hex.EncodeToString(hash[:])
}

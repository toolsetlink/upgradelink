package cdn

import (
	"fmt"
	"net/url"
)

// ReturnCdnUrl 返回cdn连接
func ReturnCdnUrl(CdnUrl string, originalURL string) (string, error) {

	newDomain := CdnUrl

	// 解析原始URL
	parsedOriginal, err := url.Parse(originalURL)
	if err != nil {
		return "", fmt.Errorf("解析原始URL失败: %v", err)
	}

	// 解析新域名
	parsedNewDomain, err := url.Parse(newDomain)
	if err != nil {
		return "", fmt.Errorf("解析新域名失败: %v", err)
	}

	// 构建新的URL
	newURL := parsedNewDomain.Scheme + "://" + parsedNewDomain.Host + parsedOriginal.Path

	// 如果原始URL有查询参数，则添加到新URL
	if parsedOriginal.RawQuery != "" {
		newURL += "?" + parsedOriginal.RawQuery
	}

	// 如果原始URL有片段标识符，则添加到新URL
	if parsedOriginal.Fragment != "" {
		newURL += "#" + parsedOriginal.Fragment
	}

	return newURL, nil
}

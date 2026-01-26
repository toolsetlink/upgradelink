package common

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// ConvertToInt64Slice 函数用于将以逗号分隔的字符串转换为 []int64 切片
func ConvertToInt64Slice(s string) ([]int64, error) {
	// 使用 strings.Split 函数按逗号分割字符串
	parts := strings.Split(s, ",")
	// 创建一个 []int64 切片，用于存储转换后的整数
	result := make([]int64, 0, len(parts))
	// 遍历分割后的字符串切片
	for _, part := range parts {
		// 去除字符串两端的空白字符
		part = strings.TrimSpace(part)
		if part == "" {
			// 如果为空字符串，则跳过
			continue
		}
		// 使用 strconv.ParseInt 函数将字符串转换为 int64 类型
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			// 如果转换过程中出现错误，则返回错误信息
			return nil, err
		}
		// 将转换后的整数添加到结果切片中
		result = append(result, num)
	}
	return result, nil
}

// GetCurrentTime 获取当前时间
func GetCurrentTime() time.Time {
	return time.Now()
}

// IsTimeOutside 判断当前时间是否在指定的时间范围外
func IsTimeOutside(start, end time.Time) bool {
	now := time.Now()
	return now.Before(start) || now.After(end)
}

// CalculateTimeDifferenceSeconds  计算两个时间差的秒
func CalculateTimeDifferenceSeconds(t1, t2 time.Time) int {
	diff := t2.Sub(t1)
	return int(diff.Seconds())
}

// ParseRFC3339ToTime ParseRFC3339 函数将 RFC3339 格式的字符串转换为 *time.Time 类型
// 如果输入的字符串格式不符合 RFC3339 标准，返回 nil 和错误信息
func ParseRFC3339ToTime(timeStr string) (*time.Time, error) {
	// 使用 time.Parse 函数解析时间字符串，RFC3339 是 Go 语言内置的标准时间格式
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return nil, fmt.Errorf("解析时间失败: %w", err)
	}
	return &t, nil
}

// SemVerToInt64 将SemVer格式的版本号转换为整数
// 格式为MAJOR.MINOR.PATCH，转换为MAJOR*1000000 + MINOR*1000 + PATCH
func SemVerToInt64(version string) (int64, error) {
	// 去除可能存在的前导v字符
	version = strings.TrimPrefix(version, "v")

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return 0, fmt.Errorf("版本号格式不正确，需要MAJOR.MINOR.PATCH格式")
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil || major < 0 {
		return 0, fmt.Errorf("MAJOR部分必须是正整数")
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil || minor < 0 {
		return 0, fmt.Errorf("MINOR部分必须是正整数")
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil || patch < 0 {
		return 0, fmt.Errorf("PATCH部分必须是正整数")
	}

	return int64(major*1000000 + minor*1000 + patch), nil
}

// ConvertLink 函数用于将原链接转换为带有代理前缀的链接
func ConvertLink(originalLink string) string {
	// 去除原链接的 https:// 部分
	linkWithoutProtocol := strings.TrimPrefix(originalLink, "https://")
	// 拼接上代理前缀
	convertedLink := fmt.Sprintf("https://gh-proxy.com/%s", linkWithoutProtocol)
	return convertedLink
}

// CalculateMD5 计算 struct 的 MD5 值
func CalculateMD5(obj interface{}) (string, error) {
	// 将 struct 转换为 JSON 字节
	jsonData, err := json.Marshal(obj)
	if err != nil {
		return "", fmt.Errorf("JSON 编码错误: %v", err)
	}

	// 计算 MD5 哈希
	hash := md5.Sum(jsonData)

	// 将哈希转换为十六进制字符串
	return hex.EncodeToString(hash[:]), nil
}

// SplitStringToIntSlice
// 使用 strings.Split 函数将字符串按逗号拆分成字符串切片
func SplitStringToIntSlice(s string) ([]int, error) {
	strs := strings.Split(s, ",")
	// 创建一个整数切片，用于存储转换后的整数
	ints := make([]int, len(strs))
	// 遍历字符串切片，将每个字符串元素转换为整数
	for i, str := range strs {
		// 使用 strconv.Atoi 函数将字符串转换为整数
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		// 将转换后的整数添加到整数切片中
		ints[i] = num
	}
	return ints, nil
}

// IntSliceToString  将整数切片按逗号拼接成字符串
func IntSliceToString(ints []int) string {
	strs := make([]string, len(ints))
	for i, num := range ints {
		strs[i] = strconv.Itoa(num)
	}
	return strings.Join(strs, ",")
}

// StringToTime 将YYYY-MM-DD HH:mm:ss格式的字符串转换为*time.Time
func StringToTime(timeStr string) (*time.Time, error) {
	// 定义时间格式
	layout := "2006-01-02 15:04:05"
	// 解析字符串为时间对象
	//t, err := time.Parse(layout, timeStr)
	t, err := time.ParseInLocation(layout, timeStr, time.Local)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// IsStartTimeBeforeEndTime  判断开始时间是否小于结束时间
func IsStartTimeBeforeEndTime(startTime, endTime string) bool {
	startTime1, _ := StringToTime(startTime)
	endTime1, _ := StringToTime(endTime)

	return startTime1.Before(*endTime1)
}

// SemVerToInt 将SemVer格式的版本号转换为整数
// 格式为MAJOR.MINOR.PATCH，转换为MAJOR*1000000 + MINOR*1000 + PATCH
func SemVerToInt(version string) (int, error) {
	// 去除可能存在的前导v字符
	version = strings.TrimPrefix(version, "v")

	parts := strings.Split(version, ".")
	if len(parts) != 3 {
		return 0, fmt.Errorf("版本号格式不正确，需要MAJOR.MINOR.PATCH格式")
	}

	major, err := strconv.Atoi(parts[0])
	if err != nil || major < 0 {
		return 0, fmt.Errorf("MAJOR部分必须是正整数")
	}

	minor, err := strconv.Atoi(parts[1])
	if err != nil || minor < 0 {
		return 0, fmt.Errorf("MINOR部分必须是正整数")
	}

	patch, err := strconv.Atoi(parts[2])
	if err != nil || patch < 0 {
		return 0, fmt.Errorf("PATCH部分必须是正整数")
	}

	return major*1000000 + minor*1000 + patch, nil
}

// BytesToMBString 字节转 MB
func BytesToMBString(bytes uint64) string {
	mb := float64(bytes) / (1024 * 1024)
	return fmt.Sprintf("%.2f MB", mb)
}

// IntUnixMilliToTime 仅支持 UnixMilli 格式（毫秒级）时间戳转换为*time.Time
// 时间戳单位：毫秒，默认使用本地时区（CST+0800）
func IntUnixMilliToTime(milliTimestamp int) (*time.Time, error) {
	// 转换为int64，避免32位系统溢出，适配time.UnixMilli参数要求
	ts := int64(milliTimestamp)

	// 无效时间戳校验：Unix时间戳（毫秒级）不支持负数
	if ts < 0 {
		return nil, errors.New("无效的毫秒级时间戳，不支持负数")
	}

	// 仅处理毫秒级时间戳，使用UnixMilli转换并指定本地时区
	t := time.UnixMilli(ts).In(time.Local)
	return &t, nil
}

// IntUnixToTime 仅支持 Unix 格式（秒级）时间戳转换为*time.Time
// 时间戳单位：秒，默认使用本地时区（CST+0800）
func IntUnixToTime(secTimestamp int) (*time.Time, error) {
	// 转换为int64，避免32位系统溢出，适配time.Unix参数要求
	ts := int64(secTimestamp)

	// 无效时间戳校验：Unix时间戳（秒级）不支持负数
	if ts < 0 {
		return nil, errors.New("无效的秒级时间戳，不支持负数")
	}

	// 仅处理秒级时间戳，使用Unix转换（纳秒部分为0）并指定本地时区
	t := time.Unix(ts, 0).In(time.Local)
	return &t, nil
}

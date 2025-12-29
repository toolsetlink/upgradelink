package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

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

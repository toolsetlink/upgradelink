package utils

import (
	"fmt"
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
	t, err := time.Parse(layout, timeStr)
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

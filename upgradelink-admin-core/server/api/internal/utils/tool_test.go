package utils

import (
	"testing"
)

// TestStringToTime 测试 StringToTime 函数
func TestStringToTime(t *testing.T) {
	// 定义测试用例：结构体包含输入、预期结果、是否预期错误
	testCases := []struct {
		name        string // 用例名称
		input       string // 输入时间字符串
		expected    string // 预期格式化后的时间字符串（统一格式便于校验）
		expectError bool   // 是否预期出现错误
	}{
		{
			name:        "合法时间格式-正常转换",
			input:       "2025-12-25 10:30:45",
			expected:    "2025-12-25 10:30:45",
			expectError: false,
		},
		{
			name:        "非法时间格式-缺少秒数",
			input:       "2025-12-25 10:30",
			expected:    "",
			expectError: true,
		},
		{
			name:        "非法时间格式-分隔符错误",
			input:       "2025/12/25 10:30:45",
			expected:    "",
			expectError: true,
		},
		{
			name:        "边界时间-1970年初",
			input:       "1970-01-01 08:00:01",
			expected:    "1970-01-01 08:00:01",
			expectError: false,
		},
	}

	// 遍历执行测试用例
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 调用待测试函数
			actualTime, err := StringToTime(tc.input)

			// 校验错误是否符合预期
			if (err != nil) != tc.expectError {
				t.Fatalf("错误状态不符合预期：预期是否报错 %v，实际报错 %v，错误信息：%v",
					tc.expectError, err != nil, err)
			}

			// 若预期无错误，进一步校验时间结果
			if !tc.expectError {
				actualStr := actualTime.Format("2006-01-02 15:04:05")
				if actualStr != tc.expected {
					t.Errorf("时间转换结果不符合预期：预期 %s，实际 %s", tc.expected, actualStr)
				}
			}
		})
	}
}

// TestIntUnixMilliToTime 测试 IntUnixMilliToTime 函数（仅支持毫秒级时间戳）
func TestIntUnixMilliToTime(t *testing.T) {
	testCases := []struct {
		name        string
		input       int    // 输入毫秒级时间戳
		expected    string // 预期格式化时间
		expectError bool
	}{
		{
			name:        "有效毫秒级-1000000毫秒（对应1970-01-01 08:00:01）",
			input:       1000000,
			expected:    "1970-01-01 08:00:01",
			expectError: false,
		},
		{
			name:        "有效毫秒级-当前时间附近（示例值，可根据实际调整）",
			input:       1735689600000, // 2025-01-01 00:00:00 CST
			expected:    "2025-01-01 00:00:00",
			expectError: false,
		},
		{
			name:        "无效毫秒级-负数时间戳",
			input:       -1000,
			expected:    "",
			expectError: true,
		},
		{
			name:        "有效毫秒级-1000毫秒（对应1970-01-01 08:00:01）",
			input:       1000,
			expected:    "1970-01-01 08:00:01",
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualTime, err := IntUnixMilliToTime(tc.input)

			if (err != nil) != tc.expectError {
				t.Fatalf("错误状态不符合预期：预期是否报错 %v，实际报错 %v，错误信息：%v",
					tc.expectError, err != nil, err)
			}

			if !tc.expectError {
				actualStr := actualTime.Format("2006-01-02 15:04:05")
				if actualStr != tc.expected {
					t.Errorf("时间转换结果不符合预期：预期 %s，实际 %s", tc.expected, actualStr)
				}
			}
		})
	}
}

// TestIntUnixToTime 测试 IntUnixToTime 函数（仅支持秒级时间戳）
func TestIntUnixToTime(t *testing.T) {
	testCases := []struct {
		name        string
		input       int    // 输入秒级时间戳
		expected    string // 预期格式化时间
		expectError bool
	}{
		{
			name:        "有效秒级-1000秒（对应1970-01-01 08:00:01）",
			input:       1000,
			expected:    "1970-01-01 08:00:01",
			expectError: false,
		},
		{
			name:        "有效秒级-当前时间附近（示例值，可根据实际调整）",
			input:       1735689600, // 2025-01-01 00:00:00 CST
			expected:    "2025-01-01 00:00:00",
			expectError: false,
		},
		{
			name:        "无效秒级-负数时间戳",
			input:       -1000,
			expected:    "",
			expectError: true,
		},
		{
			name:        "有效秒级-0秒（对应1970-01-01 08:00:00）",
			input:       0,
			expected:    "1970-01-01 08:00:00",
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualTime, err := IntUnixToTime(tc.input)

			if (err != nil) != tc.expectError {
				t.Fatalf("错误状态不符合预期：预期是否报错 %v，实际报错 %v，错误信息：%v",
					tc.expectError, err != nil, err)
			}

			if !tc.expectError {
				actualStr := actualTime.Format("2006-01-02 15:04:05")
				if actualStr != tc.expected {
					t.Errorf("时间转换结果不符合预期：预期 %s，实际 %s", tc.expected, actualStr)
				}
			}
		})
	}
}

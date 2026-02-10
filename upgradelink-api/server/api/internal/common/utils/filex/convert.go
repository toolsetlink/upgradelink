package filex

import (
	"upgradelink-api/server/api/internal/common/enum"
)

// ConvertFileTypeToUint8 converts file type string to uint8.
func ConvertFileTypeToUint8(fileType string) uint8 {
	switch fileType {
	case "other":
		return enum.Other
	case "image":
		return enum.Image
	case "video":
		return enum.Video
	case "audio":
		return enum.Audio
	default:
		return enum.Other
	}
}

// 反向转换函数
func ConvertUint8ToFileType(fileTypeUint8 uint8) string {
	switch fileTypeUint8 {
	case enum.Other:
		return "other"
	case enum.Image:
		return "image"
	case enum.Video:
		return "video"
	case enum.Audio:
		return "audio"
	default:
		return "other"
	}
}

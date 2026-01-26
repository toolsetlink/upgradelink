package filex

import (
	"upgradelink-admin/server/api/internal/common/enum"
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

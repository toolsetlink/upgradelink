package filex

import (
	"context"

	"github.com/suyuan32/simple-admin-common/enum/errorcode"
	"github.com/zeromicro/go-zero/core/errorx"

	"upgradelink-admin-file/internal/svc"
)

func CheckOverSize(ctx context.Context, svCtx *svc.ServiceContext, fileType string, size int64) error {
	if fileType == "image" && size > svCtx.Config.UploadConf.MaxImageSize {
		return errorx.NewCodeError(errorcode.InvalidArgument,
			svCtx.Trans.Trans(ctx, "file.overSizeError"))
	} else if fileType == "video" && size > svCtx.Config.UploadConf.MaxVideoSize {
		return errorx.NewCodeError(errorcode.InvalidArgument,
			svCtx.Trans.Trans(ctx, "file.overSizeError"))
	} else if fileType == "audio" && size > svCtx.Config.UploadConf.MaxAudioSize {
		return errorx.NewCodeError(errorcode.InvalidArgument,
			svCtx.Trans.Trans(ctx, "file.overSizeError"))
	} else if fileType != "image" && fileType != "video" && fileType != "audio" &&
		size > svCtx.Config.UploadConf.MaxOtherSize {
		return errorx.NewCodeError(errorcode.InvalidArgument,
			svCtx.Trans.Trans(ctx, "file.overSizeError"))
	}

	return nil
}

package userctx

import (
	"context"
	"upgradelink-admin/server/api/internal/common/http_error"
)

// GetUserIDFromCtx returns user id from context.
func GetUserIDFromCtx(ctx context.Context) (string, error) {
	userId, ok := ctx.Value("userId").(string)
	if !ok {
		return "", http_error.NewApiInternalError("failed to get user id from context")
	}
	return userId, nil
}

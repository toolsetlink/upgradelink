package rolectx

import (
	"context"
	"slices"
	"strings"
	"upgradelink-admin/server/api/internal/common/http_error"
)

// GetRoleIDFromCtx returns role id from context.
func GetRoleIDFromCtx(ctx context.Context) ([]string, error) {
	roleId, ok := ctx.Value("roleId").(string)
	if !ok {
		return nil, http_error.NewApiInternalError("failed to get role id from context")
	}
	roleIds := strings.Split(roleId, ",")
	slices.Sort(roleIds)
	return roleIds, nil
}

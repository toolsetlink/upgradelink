package db_error

import (
	"net/http"
	"upgradelink-admin/server/api/internal/common/http_error"
	"upgradelink-admin/server/api/internal/common/i18n"

	"github.com/zeromicro/go-zero/core/logx"

	"upgradelink-admin/server/api/internal/common/msg/log_msg"
	"upgradelink-admin/server/api/internal/ent"
)

// DefaultEntError returns errors dealing with default functions.
func DefaultEntError(logger logx.Logger, err error, detail any) error {
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logger.Errorw(err.Error(), logx.Field("detail", detail))
			return http_error.NewCodeError(http.StatusNotFound, i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logger.Errorw(err.Error(), logx.Field("detail", detail))
			return http_error.NewCodeError(http.StatusBadRequest, i18n.ConstraintError)
		case ent.IsValidationError(err):
			logger.Errorw(err.Error(), logx.Field("detail", detail))
			return http_error.NewCodeError(http.StatusBadRequest, i18n.ValidationError)
		case ent.IsNotSingular(err):
			logger.Errorw(err.Error(), logx.Field("detail", detail))
			return http_error.NewCodeError(http.StatusBadRequest, i18n.NotSingularError)
		default:
			logger.Errorw(log_msg.DatabaseError, logx.Field("detail", err.Error()))
			return http_error.NewCodeError(http.StatusInternalServerError, i18n.DatabaseError)
		}
	}
	return err
}

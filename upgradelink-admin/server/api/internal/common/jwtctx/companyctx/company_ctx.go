package companyctx

import (
	"context"
	"strconv"
)

// GetCompanyIDPointerFromCtx returns company id from context.
func GetCompanyIDPointerFromCtx(ctx context.Context) *int {
	companyIdInt := 0
	companyId, ok := ctx.Value("companyId").(string)
	if !ok {
		return &companyIdInt
	}
	companyIdInt, err := strconv.Atoi(companyId)
	if err != nil {
		return &companyIdInt
	}
	return &companyIdInt
}

func GetCompanyIDFromCtx(ctx context.Context) int {
	companyId, ok := ctx.Value("companyId").(string)
	if !ok {
		return 0
	}
	companyIdInt, err := strconv.Atoi(companyId)
	if err != nil {
		return 0
	}
	return companyIdInt
}

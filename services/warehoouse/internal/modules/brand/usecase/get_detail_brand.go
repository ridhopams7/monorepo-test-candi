package usecase

import (
	"context"
	"time"

	"monorepo/services/warehoouse/internal/modules/brand/domain"
	shareddomain "monorepo/services/warehoouse/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *brandUsecaseImpl) GetDetailBrand(ctx context.Context, id string) (result domain.ResponseBrand, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BrandUsecase:GetDetailBrand")
	defer trace.Finish()

	var data shareddomain.Brand
	repoFilter := domain.FilterBrand{ID: &id}
	data, err = uc.repoSQL.BrandRepo().Find(ctx, &repoFilter)
	if err != nil {
		return result, err
	}

	result.ID = data.ID
	result.Title = data.Title
	result.CreatedAt = data.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = data.UpdatedAt.Format(time.RFC3339)
	return
}

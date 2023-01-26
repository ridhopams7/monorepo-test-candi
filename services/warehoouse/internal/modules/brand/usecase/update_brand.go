package usecase

import (
	"context"

	"github.com/golangid/candi/tracer"
	"monorepo/services/warehoouse/internal/modules/brand/domain"
)

func (uc *brandUsecaseImpl) UpdateBrand(ctx context.Context, data *domain.RequestBrand) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BrandUsecase:UpdateBrand")
	defer trace.Finish()

	repoFilter := domain.FilterBrand{ID: &data.ID}
	existing, err := uc.repoSQL.BrandRepo().Find(ctx, &repoFilter)
	if err != nil {
		return err
	}
	existing.Title = data.Title
	err = uc.repoSQL.BrandRepo().Save(ctx, &existing)
	return
}

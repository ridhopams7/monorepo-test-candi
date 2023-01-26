package usecase

import (
	"context"
	"time"

	"monorepo/services/warehoouse/internal/modules/brand/domain"
	shareddomain "monorepo/services/warehoouse/pkg/shared/domain"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/tracer"
)

func (uc *brandUsecaseImpl) GetAllBrand(ctx context.Context, filter *domain.FilterBrand) (results []domain.ResponseBrand, meta candishared.Meta, err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BrandUsecase:GetAllBrand")
	defer trace.Finish()

	var data []shareddomain.Brand
	data, err = uc.repoSQL.BrandRepo().FetchAll(ctx, filter)
	if err != nil {
		return results, meta, err
	}
	count := uc.repoSQL.BrandRepo().Count(ctx, filter)
	meta = candishared.NewMeta(filter.Page, filter.Limit, count)

	results = make([]domain.ResponseBrand, 0)
	for _, detail := range data {
		results = append(results, domain.ResponseBrand{
			ID:        detail.ID,
			Title:     detail.Title,
			CreatedAt: detail.CreatedAt.Format(time.RFC3339),
			UpdatedAt: detail.UpdatedAt.Format(time.RFC3339),
		})
	}

	return
}

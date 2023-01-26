package usecase

import (
	"context"

	"monorepo/services/warehoouse/internal/modules/brand/domain"
	shareddomain "monorepo/services/warehoouse/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *brandUsecaseImpl) CreateBrand(ctx context.Context, req *domain.RequestBrand) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BrandUsecase:CreateBrand")
	defer trace.Finish()

	return uc.repoSQL.BrandRepo().Save(ctx, &shareddomain.Brand{
		Title: req.Title,
	})
}

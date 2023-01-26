package usecase

import (
	"context"
	
	shareddomain "monorepo/services/warehoouse/pkg/shared/domain"

	"github.com/golangid/candi/tracer"
)

func (uc *brandUsecaseImpl) DeleteBrand(ctx context.Context, id string) (err error) {
	trace, ctx := tracer.StartTraceWithContext(ctx, "BrandUsecase:DeleteBrand")
	defer trace.Finish()

	
	return uc.repoSQL.BrandRepo().Delete(ctx, &shareddomain.Brand{
		ID: id,
	})
}

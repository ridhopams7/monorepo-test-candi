package usecase

import (
	"context"

	mockrepo "monorepo/services/warehoouse/pkg/mocks/modules/brand/repository"
	mocksharedrepo "monorepo/services/warehoouse/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/warehoouse/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_brandUsecaseImpl_GetDetailBrand(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		brandRepo := &mockrepo.BrandRepository{}
		brandRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Brand{}, nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BrandRepo").Return(brandRepo)

		uc := brandUsecaseImpl{
			repoSQL: repoSQL,
		}

		_, err := uc.GetDetailBrand(context.Background(), "id")
		assert.NoError(t, err)
	})
}

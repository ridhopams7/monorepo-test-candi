package usecase

import (
	"context"

	mockrepo "monorepo/services/warehoouse/pkg/mocks/modules/brand/repository"
	mocksharedrepo "monorepo/services/warehoouse/pkg/mocks/shared/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_brandUsecaseImpl_DeleteBrand(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		brandRepo := &mockrepo.BrandRepository{}
		brandRepo.On("Delete", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BrandRepo").Return(brandRepo)

		uc := brandUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.DeleteBrand(context.Background(), "id")
		assert.NoError(t, err)
	})
}

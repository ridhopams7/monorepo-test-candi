package usecase

import (
	"context"
	"errors"

	"monorepo/services/warehoouse/internal/modules/brand/domain"
	mockrepo "monorepo/services/warehoouse/pkg/mocks/modules/brand/repository"
	mocksharedrepo "monorepo/services/warehoouse/pkg/mocks/shared/repository"
	shareddomain "monorepo/services/warehoouse/pkg/shared/domain"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_brandUsecaseImpl_UpdateBrand(t *testing.T) {
	t.Run("Testcase #1: Positive", func(t *testing.T) {

		brandRepo := &mockrepo.BrandRepository{}
		brandRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Brand{}, nil)
		brandRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BrandRepo").Return(brandRepo)

		uc := brandUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateBrand(context.Background(), &domain.RequestBrand{})
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {

		brandRepo := &mockrepo.BrandRepository{}
		brandRepo.On("Find", mock.Anything, mock.Anything).Return(shareddomain.Brand{}, errors.New("Error"))
		brandRepo.On("Save", mock.Anything, mock.Anything).Return(nil)

		repoSQL := &mocksharedrepo.RepoSQL{}
		repoSQL.On("BrandRepo").Return(brandRepo)

		uc := brandUsecaseImpl{
			repoSQL: repoSQL,
		}

		err := uc.UpdateBrand(context.Background(), &domain.RequestBrand{})
		assert.Error(t, err)
	})
}

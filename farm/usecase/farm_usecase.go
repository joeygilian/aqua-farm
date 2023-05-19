package usecase

import (
	"github.com/aqua-farm/farm"
	"github.com/aqua-farm/farm/repository"
)

type FarmUsecase interface {
	Fetch(cursor int64, num int64) ([]*farm.Farm, string, error)
}

type farmUsecase struct {
	farmRepo repository.FarmRepository
}

func NewFarmUsecase(farmRepo repository.FarmRepository) FarmUsecase {
	return &farmUsecase{farmRepo: farmRepo}
}

func (f *farmUsecase) Fetch(cursor int64, num int64) ([]*farm.Farm, string, error) {
	if num == 0 {
		num = 10
	}

	listArticle, err := f.farmRepo.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""

	return listArticle, nextCursor, nil

}

package usecase

import (
	"github.com/aqua-farm/pond"
	"github.com/aqua-farm/pond/repository"
)

type PondUsecase interface {
	Fetch(cursor int64, num int64) ([]*pond.Pond, string, error)
}

type pondUsecase struct {
	pondRepo repository.PondRepository
}

func NewPondUsecase(pondRepo repository.PondRepository) PondUsecase {
	return &pondUsecase{pondRepo: pondRepo}
}

func (f *pondUsecase) Fetch(cursor int64, num int64) ([]*pond.Pond, string, error) {
	if num == 0 {
		num = 10
	}

	listArticle, err := f.pondRepo.Fetch(cursor, num)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""

	return listArticle, nextCursor, nil

}

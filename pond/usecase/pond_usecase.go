package usecase

import (
	farm "github.com/aqua-farm/farm"
	farmUsecase "github.com/aqua-farm/farm/usecase"
	"github.com/aqua-farm/pond"
	"github.com/aqua-farm/pond/repository"
)

type PondUsecase interface {
	Fetch() ([]*pond.Pond, error)
	Store(pond *pond.Pond) (*pond.Pond, error)
	Update(pond *pond.Pond) (*pond.Pond, error)
}

type pondUsecase struct {
	pondRepo    repository.PondRepository
	farmUsecase farmUsecase.FarmUsecase
}

func NewPondUsecase(pondRepo repository.PondRepository, farmUsecase farmUsecase.FarmUsecase) PondUsecase {
	return &pondUsecase{pondRepo: pondRepo, farmUsecase: farmUsecase}
}

// usecase for fetching all existing ponds
func (p *pondUsecase) Fetch() ([]*pond.Pond, error) {
	listPond, err := p.pondRepo.Fetch()
	if err != nil {
		return nil, err
	}

	return listPond, nil
}

// usecase for storing new pond, it will check the existed farm and if the pond already existed or not
func (p *pondUsecase) Store(pond *pond.Pond) (*pond.Pond, error) {
	existedFarm, _ := p.farmUsecase.GetByID(pond.FarmID)
	if existedFarm == nil {
		return nil, farm.ErrNotFound
	}

	existedPond, _ := p.pondRepo.GetByName(pond.Name)
	if existedPond != nil {
		return nil, farm.ErrConflict
	}

	id, err := p.pondRepo.Store(pond)
	if err != nil {
		return nil, err
	}

	pond.ID = id
	return pond, nil
}

func (p *pondUsecase) Update(pond *pond.Pond) (*pond.Pond, error) {

	existedPond, _ := p.pondRepo.GetByID(pond.ID)
	if existedPond == nil {
		return nil, farm.ErrNotFound
	}

	res, err := p.pondRepo.Update(pond)
	if err != nil {
		return nil, err
	}

	return res, nil
}

package usecase

import (
	"github.com/aqua-farm/farm"
	"github.com/aqua-farm/farm/repository"
	pondRepo "github.com/aqua-farm/pond/repository"
	"github.com/labstack/gommon/log"
)

type FarmUsecase interface {
	Fetch() ([]*farm.Farm, error)
	GetByID(id int64) (*farm.Farm, error)
	Store(f *farm.Farm) (*farm.Farm, error)
	Update(f *farm.Farm) (*farm.Farm, error)
	Delete(id int64) (int64, error)
}

type farmUsecase struct {
	farmRepo repository.FarmRepository
	pondRepo pondRepo.PondRepository
}

func NewFarmUsecase(farmRepo repository.FarmRepository, pondRepo pondRepo.PondRepository) FarmUsecase {
	return &farmUsecase{farmRepo: farmRepo, pondRepo: pondRepo}
}

func (f *farmUsecase) Fetch() ([]*farm.Farm, error) {

	listFarm, err := f.farmRepo.Fetch()
	if err != nil {
		return nil, err
	}

	for idx, farm := range listFarm {
		pond, err := f.pondRepo.GetByFarmID(farm.ID)
		if err != nil {
			return nil, err
		}
		listFarm[idx].Pond = pond
	}

	return listFarm, nil
}

func (f *farmUsecase) GetByID(id int64) (*farm.Farm, error) {

	res, err := f.farmRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (fu *farmUsecase) Store(f *farm.Farm) (*farm.Farm, error) {

	existedFarm, _ := fu.farmRepo.GetByName(f.Name)
	if existedFarm != nil {
		return nil, farm.ErrConflict
	}
	log.Info(existedFarm)

	id, err := fu.farmRepo.Store(f)
	if err != nil {
		return nil, err
	}

	f.ID = id
	return f, nil
}

func (fu *farmUsecase) Update(f *farm.Farm) (*farm.Farm, error) {

	existedFarm, _ := fu.farmRepo.GetByID(f.ID)
	if existedFarm == nil {
		return nil, farm.ErrNotFound
	}
	log.Info(existedFarm)

	res, err := fu.farmRepo.Update(f)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (fu *farmUsecase) Delete(id int64) (int64, error) {

	existedFarm, _ := fu.farmRepo.GetByID(id)
	if existedFarm == nil {
		return 0, farm.ErrNotFound
	}
	log.Info(existedFarm)

	res, err := fu.farmRepo.Delete(id)
	if err != nil {
		return 0, err
	}

	return res, nil
}

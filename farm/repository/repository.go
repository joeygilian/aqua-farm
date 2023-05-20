package repository

import models "github.com/aqua-farm/farm"

type FarmRepository interface {
	Fetch() ([]*models.Farm, error)
	GetByID(id int64) (*models.Farm, error)
	GetByName(name string) (*models.Farm, error)
	Store(f *models.Farm) (int64, error)
	Update(f *models.Farm) (*models.Farm, error)
}

package repository

import models "github.com/aqua-farm/farm"

type FarmRepository interface {
	Fetch(cursor int64, num int64) ([]*models.Farm, error)
}

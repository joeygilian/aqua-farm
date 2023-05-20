package repository

import (
	"database/sql"
	"fmt"

	models "github.com/aqua-farm/farm"
	"github.com/labstack/gommon/log"
)

type postgresqlFarmRepository struct {
	Conn *sql.DB
}

func NewPostgresqlFarmRepository(Conn *sql.DB) FarmRepository {
	return &postgresqlFarmRepository{Conn}
}

func (p *postgresqlFarmRepository) fetch(query string, args ...interface{}) ([]*models.Farm, error) {

	rows, err := p.Conn.Query(query, args...)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Farm, 0)
	for rows.Next() {
		f := new(models.Farm)
		err = rows.Scan(
			&f.ID,
			&f.Name,
		)

		if err != nil {
			return nil, err
		}
		result = append(result, f)
	}

	return result, nil
}

func (p *postgresqlFarmRepository) Fetch() ([]*models.Farm, error) {
	query := `SELECT id,name FROM farm where status = true`

	return p.fetch(query)

}

func (p *postgresqlFarmRepository) GetByID(id int64) (*models.Farm, error) {
	query := `SELECT id, name
  						FROM farm WHERE id = $1 and status = true`

	a := &models.Farm{}
	err := p.Conn.QueryRow(query, id).Scan(&a.ID, &a.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("farm with id %d not found", id)
		}
	}
	return a, nil
}

func (p *postgresqlFarmRepository) GetByName(name string) (*models.Farm, error) {
	query := `SELECT id, name FROM farm WHERE name = $1 and status = true`

	a := &models.Farm{}
	err := p.Conn.QueryRow(query, name).Scan(&a.ID, &a.Name)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (p *postgresqlFarmRepository) Store(f *models.Farm) (int64, error) {
	query := "INSERT INTO farm(name) VALUES ($1) RETURNING id"
	err := p.Conn.QueryRow(query, f.Name).Scan(&f.ID)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return f.ID, nil
}

func (p *postgresqlFarmRepository) Update(f *models.Farm) (*models.Farm, error) {

	_, err := p.Conn.Exec("UPDATE farm SET name=$2, updated_date = now() WHERE id=$1", f.ID, f.Name)
	if err != nil {
		return nil, err
	}

	return f, nil
}

package repository

import (
	"database/sql"
	"fmt"

	models "github.com/aqua-farm/pond"

	"github.com/labstack/gommon/log"
)

type postgresqlPondRepository struct {
	Conn *sql.DB
}

func NewPostgresqlPondRepository(Conn *sql.DB) PondRepository {
	return &postgresqlPondRepository{Conn}
}

type PondRepository interface {
	Fetch() ([]*models.Pond, error)
	GetByID(id int64) (*models.Pond, error)
	GetByName(name string) (*models.Pond, error)
	Store(pond *models.Pond) (int64, error)
	Update(pond *models.Pond) (*models.Pond, error)
}

func (m *postgresqlPondRepository) fetch(query string, args ...interface{}) ([]*models.Pond, error) {

	rows, err := m.Conn.Query(query, args...)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	result := make([]*models.Pond, 0)
	for rows.Next() {
		f := new(models.Pond)
		err = rows.Scan(
			&f.ID,
			&f.Name,
		)

		if err != nil {
			// logrus.Error(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, nil
}

func (p *postgresqlPondRepository) Fetch() ([]*models.Pond, error) {

	query := `SELECT id, farm_id ,name FROM pond`

	return p.fetch(query)
}

func (p *postgresqlPondRepository) GetByID(id int64) (*models.Pond, error) {
	query := `SELECT id, farm_id, name
				FROM pond WHERE id = $1`

	pond := &models.Pond{}
	err := p.Conn.QueryRow(query, id).Scan(&pond.ID, &pond.FarmID, &pond.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("farm with id %d not found", id)
		}
	}
	return pond, nil
}

func (p *postgresqlPondRepository) GetByName(name string) (*models.Pond, error) {
	query := `SELECT id, farm_id, name
  						FROM pond WHERE name = $1`

	a := &models.Pond{}
	err := p.Conn.QueryRow(query, name).Scan(&a.ID, &a.Name)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (p *postgresqlPondRepository) Store(pond *models.Pond) (int64, error) {

	query := "INSERT INTO pond (farm_id, name) VALUES ($1, $2) RETURNING id"
	err := p.Conn.QueryRow(query, pond.FarmID, pond.Name).Scan(&pond.ID)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return pond.ID, nil
}

func (p *postgresqlPondRepository) Update(pond *models.Pond) (*models.Pond, error) {

	_, err := p.Conn.Exec("UPDATE pond SET name=$2, farm_id = $3 WHERE id=$1", pond.ID, pond.Name, pond.FarmID)
	if err != nil {
		return nil, err
	}

	return pond, nil
}

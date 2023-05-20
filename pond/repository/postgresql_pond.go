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
	GetByFarmID(id int64) ([]*models.Pond, error)
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
			&f.FarmID,
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

// repository for fetching all pond
func (p *postgresqlPondRepository) Fetch() ([]*models.Pond, error) {

	query := `SELECT id, farm_id ,name FROM pond where status = true`
	log.Info("masuk")
	return p.fetch(query)
}

// repository for get pond by farmID
func (p *postgresqlPondRepository) GetByFarmID(id int64) ([]*models.Pond, error) {

	query := `SELECT id, farm_id ,name FROM pond where status = true and farm_id = $1`

	return p.fetch(query, id)
}

// repository for get pond by id
func (p *postgresqlPondRepository) GetByID(id int64) (*models.Pond, error) {
	query := `SELECT id, farm_id, name
				FROM pond WHERE id = $1 and status = true`

	pond := &models.Pond{}
	err := p.Conn.QueryRow(query, id).Scan(&pond.ID, &pond.FarmID, &pond.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("farm with id %d not found", id)
		}
	}
	return pond, nil
}

// repository for get pond by name
func (p *postgresqlPondRepository) GetByName(name string) (*models.Pond, error) {
	query := `SELECT id, farm_id, name
  						FROM pond WHERE name = $1 and status = true`

	a := &models.Pond{}
	err := p.Conn.QueryRow(query, name).Scan(&a.ID, &a.Name)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// repository for storing pond
func (p *postgresqlPondRepository) Store(pond *models.Pond) (int64, error) {

	query := "INSERT INTO pond (farm_id, name) VALUES ($1, $2) RETURNING id"
	err := p.Conn.QueryRow(query, pond.FarmID, pond.Name).Scan(&pond.ID)
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return pond.ID, nil
}

// repository for update pond
func (p *postgresqlPondRepository) Update(pond *models.Pond) (*models.Pond, error) {

	_, err := p.Conn.Exec("UPDATE pond SET name=$2, farm_id = $3, updated_date = now() WHERE id=$1", pond.ID, pond.Name, pond.FarmID)
	if err != nil {
		return nil, err
	}

	return pond, nil
}

// repository for delete pond
func (p *postgresqlPondRepository) Delete(id int64) (int64, error) {

	_, err := p.Conn.Exec("UPDATE pond SET status = false, deleted_date = now() WHERE id=$1", id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

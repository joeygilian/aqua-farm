package repository

import (
	"database/sql"

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
	Fetch(cursor int64, num int64) ([]*models.Pond, error)
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

func (m *postgresqlPondRepository) Fetch(cursor int64, num int64) ([]*models.Pond, error) {
	log.Error("cursor: ", cursor)

	query := `SELECT id, farm_id ,name FROM pond`

	return m.fetch(query)

}

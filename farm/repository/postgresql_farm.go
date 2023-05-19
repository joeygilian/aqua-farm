package repository

import (
	"database/sql"

	models "github.com/aqua-farm/farm"
	"github.com/labstack/gommon/log"
)

type postgresqlFarmRepository struct {
	Conn *sql.DB
}

func NewPostgresqlFarmRepository(Conn *sql.DB) FarmRepository {
	return &postgresqlFarmRepository{Conn}
}

func (m *postgresqlFarmRepository) fetch(query string, args ...interface{}) ([]*models.Farm, error) {

	rows, err := m.Conn.Query(query, args...)

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
			// logrus.Error(err)
			return nil, err
		}
		result = append(result, f)
	}

	return result, nil
}

func (m *postgresqlFarmRepository) Fetch(cursor int64, num int64) ([]*models.Farm, error) {
	log.Error("cursor: ", cursor)

	query := `SELECT id,name FROM farm`

	return m.fetch(query)

}

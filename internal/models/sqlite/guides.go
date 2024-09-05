package sqlite

import (
	"database/sql"

	"github.com/MuhammadISayyed/tov.git/internal/models"
)

type GuideModel struct {
	DB *sql.DB
}

func (m *GuideModel) All() ([]models.Guide, error) {
	stmt := `SELECT id, title, content, createdAt FROM guides ORDER BY id DESC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	guides := []models.Guide{}
	for rows.Next() {
		g := models.Guide{}
		err := rows.Scan(&g.ID, &g.Title, &g.Content, &g.CreatedAt)
		if err != nil {
			return nil, err
		}
		guides = append(guides, g)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return guides, nil
}
package sqlite

import (
	"database/sql"

	"github.com/MuhammadISayyed/tov.git/internal/models"
)

type GuideModel struct {
	DB *sql.DB
}

func (m *GuideModel) Insert(title, content string) error {
	stmt := `INSERT INTO guides (title, content, createdAt)
	VALUES (?, ?, datetime('now'))`
	_, err := m.DB.Exec(stmt, title, content)
	return err
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
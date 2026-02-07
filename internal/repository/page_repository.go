package repository

import (
	"context"
	"appdrop-api/internal/db"
	"appdrop-api/internal/models"
)

func GetAllPages() ([]models.Page, error) {
	rows, err := db.Pool.Query(context.Background(),
		`SELECT id, name, route, is_home, created_at, updated_at FROM pages ORDER BY created_at`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pages []models.Page

	for rows.Next() {
		var p models.Page
		err := rows.Scan(&p.ID, &p.Name, &p.Route, &p.IsHome, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		pages = append(pages, p)
	}

	return pages, nil
}

func CreatePage(page models.Page) error {
	_, err := db.Pool.Exec(context.Background(),
		`INSERT INTO pages (name, route, is_home) VALUES ($1,$2,$3)`,
		page.Name, page.Route, page.IsHome,
	)
	return err
}

// Package repository provides data access operations for the AppDrop API.
// This layer handles all direct database interactions using parameterized queries.
// Repository functions use raw SQL with pgx for efficient database operations.
package repository

import (
	"appdrop-api/internal/db"
	"appdrop-api/internal/models"
	"context"
)

// GetAllPages retrieves all pages from the database.
// Returns a slice of all pages ordered by creation date, or error on database failure.
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

func CreatePage(page models.Page) (*models.Page, error) {
	// CreatePage inserts a new page into the database and returns the created page.
	// Uses RETURNING clause to get auto-generated ID and timestamps in one query.
	var createdPage models.Page
	err := db.Pool.QueryRow(context.Background(),
		`INSERT INTO pages (name, route, is_home) VALUES ($1,$2,$3) RETURNING id, name, route, is_home, created_at, updated_at`,
		page.Name, page.Route, page.IsHome,
	).Scan(&createdPage.ID, &createdPage.Name, &createdPage.Route, &createdPage.IsHome, &createdPage.CreatedAt, &createdPage.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &createdPage, nil
}

func RouteExists(route string) (bool, error) {
	// RouteExists checks if a page with the given route already exists.
	// Used to enforce route uniqueness constraint.
	var exists bool
	err := db.Pool.QueryRow(context.Background(),
		`SELECT EXISTS(SELECT 1 FROM pages WHERE route=$1)`,
		route,
	).Scan(&exists)

	return exists, err
}

func ResetHomePage() error {
	// ResetHomePage sets is_home=false for all pages.
	// Called before making a different page the home page to maintain the single home page constraint.
	_, err := db.Pool.Exec(context.Background(),
		`UPDATE pages SET is_home = false WHERE is_home = true`)
	return err
}

func GetPageByID(id string) (*models.Page, error) {
	// GetPageByID retrieves a page from the database by its UUID.
	// Returns nil if the page is not found.
	var p models.Page

	err := db.Pool.QueryRow(context.Background(),
		`SELECT id,name,route,is_home,created_at,updated_at 
		 FROM pages WHERE id=$1`, id).
		Scan(&p.ID, &p.Name, &p.Route, &p.IsHome, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &p, nil
}

func DeletePage(id string) error {
	// DeletePage removes a page and all associated widgets (due to ON DELETE CASCADE).
	_, err := db.Pool.Exec(context.Background(),
		`DELETE FROM pages WHERE id=$1`, id)
	return err
}

func UpdatePage(page models.Page) (*models.Page, error) {
	// UpdatePage modifies page details and returns the updated page.
	// Uses RETURNING clause to get updated timestamps and values in one query.
	var updatedPage models.Page
	err := db.Pool.QueryRow(context.Background(),
		`UPDATE pages 
		 SET name=$1, route=$2, is_home=$3, updated_at=NOW()
		 WHERE id=$4 RETURNING id, name, route, is_home, created_at, updated_at`,
		page.Name, page.Route, page.IsHome, page.ID,
	).Scan(&updatedPage.ID, &updatedPage.Name, &updatedPage.Route, &updatedPage.IsHome, &updatedPage.CreatedAt, &updatedPage.UpdatedAt)

	if err != nil {
		return nil, err
	}
	return &updatedPage, nil
}

func RouteExistsForOtherPage(route, id string) (bool, error) {
	// RouteExistsForOtherPage checks if a route exists on a different page.
	// Used during update to allow the same page to keep its own route.
	var exists bool
	err := db.Pool.QueryRow(context.Background(),
		`SELECT EXISTS(SELECT 1 FROM pages WHERE route=$1 AND id != $2)`,
		route, id,
	).Scan(&exists)

	return exists, err
}

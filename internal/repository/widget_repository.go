package repository

import (
	"appdrop-api/internal/db"
	"appdrop-api/internal/models"
	"context"
	"encoding/json"
)

func GetWidgetsByPageID(pageID string) ([]models.Widget, error) {

	rows, err := db.Pool.Query(context.Background(),
		`SELECT id,page_id,type,position,config,created_at,updated_at 
		 FROM widgets WHERE page_id=$1 ORDER BY position`, pageID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var widgets []models.Widget

	for rows.Next() {
		var w models.Widget
		var configJSON []byte
		err := rows.Scan(
			&w.ID, &w.PageID, &w.Type,
			&w.Position, &configJSON,
			&w.CreatedAt, &w.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		// Parse JSONB config
		if configJSON != nil {
			json.Unmarshal(configJSON, &w.Config)
		}
		widgets = append(widgets, w)
	}

	return widgets, nil
}

func GetWidgetByID(id string) (*models.Widget, error) {
	var w models.Widget
	var configJSON []byte

	err := db.Pool.QueryRow(context.Background(),
		`SELECT id,page_id,type,position,config,created_at,updated_at 
		 FROM widgets WHERE id=$1`, id).
		Scan(&w.ID, &w.PageID, &w.Type, &w.Position, &configJSON, &w.CreatedAt, &w.UpdatedAt)

	if err != nil {
		return nil, err
	}

	// Parse JSONB config
	if configJSON != nil {
		json.Unmarshal(configJSON, &w.Config)
	}

	return &w, nil
}

func CreateWidget(widget models.Widget) (*models.Widget, error) {
	var createdWidget models.Widget
	var configJSON []byte

	// Marshal config to JSON for storage
	configData, err := json.Marshal(widget.Config)
	if err != nil {
		return nil, err
	}

	err = db.Pool.QueryRow(context.Background(),
		`INSERT INTO widgets (page_id,type,position,config)
		 VALUES ($1,$2,$3,$4) RETURNING id,page_id,type,position,config,created_at,updated_at`,
		widget.PageID, widget.Type, widget.Position, string(configData),
	).Scan(&createdWidget.ID, &createdWidget.PageID, &createdWidget.Type, &createdWidget.Position, &configJSON, &createdWidget.CreatedAt, &createdWidget.UpdatedAt)

	if err != nil {
		return nil, err
	}

	// Parse returned config
	if configJSON != nil {
		json.Unmarshal(configJSON, &createdWidget.Config)
	}

	return &createdWidget, nil
}

func UpdateWidget(widget models.Widget) (*models.Widget, error) {
	var updatedWidget models.Widget
	var configJSON []byte

	// Marshal config to JSON for storage
	configData, err := json.Marshal(widget.Config)
	if err != nil {
		return nil, err
	}

	err = db.Pool.QueryRow(context.Background(),
		`UPDATE widgets 
		 SET type=$1, position=$2, config=$3, updated_at=NOW()
		 WHERE id=$4 RETURNING id,page_id,type,position,config,created_at,updated_at`,
		widget.Type, widget.Position, string(configData), widget.ID,
	).Scan(&updatedWidget.ID, &updatedWidget.PageID, &updatedWidget.Type, &updatedWidget.Position, &configJSON, &updatedWidget.CreatedAt, &updatedWidget.UpdatedAt)

	if err != nil {
		return nil, err
	}

	// Parse returned config
	if configJSON != nil {
		json.Unmarshal(configJSON, &updatedWidget.Config)
	}

	return &updatedWidget, nil
}

func DeleteWidget(id string) error {
	_, err := db.Pool.Exec(context.Background(),
		`DELETE FROM widgets WHERE id=$1`, id)
	return err
}

func ReorderWidgets(pageID string, ids []string) error {
	tx, err := db.Pool.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	for index, id := range ids {
		_, err := tx.Exec(context.Background(),
			`UPDATE widgets SET position=$1, updated_at=NOW() WHERE id=$2 AND page_id=$3`,
			index, id, pageID)
		if err != nil {
			return err
		}
	}

	return tx.Commit(context.Background())
}

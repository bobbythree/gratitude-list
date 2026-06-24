package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

func openDB() (*sql.DB, error) {
	dbPath := os.Getenv("DB_PATH")

	if dbPath == "" {
		return nil, fmt.Errorf("DB_PATH is not set")
	}

	return sql.Open("sqlite", dbPath)
}

func CreateListItemsTable(db *sql.DB) error {
	query := `
		CREATE TABLE IF NOT EXISTS listItems (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			item TEXT NOT NULL
		);
	`

	_, err := db.Exec(query)
	return err
}

func CreateListItem(db *sql.DB, listItem ListItem) error {
	_, err := db.Exec(
		"INSERT INTO listItems (item) VALUES (?)",
		listItem.Item,
	)
	return err
}

// GetListItems : get items
func GetListItems(db *sql.DB) ([]ListItem, error) {
	rows, err := db.Query("SELECT id, item FROM listItems ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []ListItem{}

	for rows.Next() {
		var item ListItem

		err := rows.Scan(&item.ID, &item.Item)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

// DeleteListItem delete
func DeleteListItem(db *sql.DB, id int64) error {
	_, err := db.Exec(
		"DELETE FROM listItems WHERE id = ?",
		id,
	)
	return err
}

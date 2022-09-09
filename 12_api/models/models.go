package models

import (
	"database/sql"
	"time"
)

// wrapper for database
type Models struct {
	DB DBModel
}


// retur models with db pool
func NewModels(db *sql.DB) Models {
	return Models {
		DB: DBModel{DB: db},
	}
}

type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}


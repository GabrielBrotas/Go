package models

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) Init(syncDbCh chan<- error) {
	log.Println("Creating tables...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	tx, err := m.DB.BeginTx(ctx, nil)

	if err != nil {
		syncDbCh <- err
		return
	}

	query_movies_table := `CREATE TABLE IF NOT EXISTS movies (
		id serial PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		description VARCHAR(255) NOT NULL,
		year INT NOT NULL,
		created_at timestamp default CURRENT_TIMESTAMP,
		updated_at timestamp default CURRENT_TIMESTAMP  
	)`

	_, err = tx.ExecContext(ctx, query_movies_table)

	if err != nil {
		log.Println("Error creating tx to movies table", err)
		tx.Rollback()
		syncDbCh <- err
		return
	}

	err = tx.Commit()

	if err != nil {
		log.Println("Error creating db", err)
		syncDbCh <- err
		return
	}

	log.Println("Tables created successfully...")
	syncDbCh <- err
}

func (m *DBModel) Get(id int) (*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `select id, title, description, year, created_at, updated_at from movies where id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (m *DBModel) All() ([]*Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT * FROM movies`

	rows, err := m.DB.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	movies := []*Movie{}
	for rows.Next() {
		var movie Movie
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
		movies = append(movies, &movie)
	}

	return movies, nil
}

type MovieInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        int    `json:"year"`
}

func (m *DBModel) Create(movie MovieInput) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO movies (title, description, year) VALUES ($1, $2, $3) RETURNING id`

	id := 0
	err := m.DB.QueryRowContext(ctx, query, movie.Title, movie.Description, movie.Year).Scan(&id)

	if err != nil {
		return id, err
	}

	return id, nil
}

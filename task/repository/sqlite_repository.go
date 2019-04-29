package repository

import (
	"database/sql"
	"github.com/hiroaki-sekine/taskforce/task"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) task.Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(query string) (int64, error) {
	result, err := r.db.Exec(query)
	if err != nil {
		return 0, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return id, err
}

func (r *Repository) ReadRow(query string) (task.Row, error) {
	row := r.db.QueryRow(query)
	return row, nil
}

func (r *Repository) ReadRows(query string) (task.Rows, error) {
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *Repository) Update(query string) error {
	result, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Delete(query string) error {
	result, err := r.db.Exec(query)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"database/sql"
	"taskforce/task"
)

type Repository struct {
	db *sql.DB
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

func (r *Repository) Read(query string) (task.Scanner, error) {
	return r.db.Query(query)
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

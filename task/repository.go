package task

type Repository interface {
	Create(query string) (int64, error)
	ReadRow(query string) (Row, error)
	ReadRows(query string) (Rows, error)
	Update(query string) error
	Delete(query string) error
}

type Row interface {
	Scan(dest ...interface{}) error
}

type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Err() error
	Close() error
}

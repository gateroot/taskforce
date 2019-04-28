package task

type Repository interface {
	Create(query string) (int64, error)
	Read(query string) (Scanner, error)
	Update(query string) error
	Delete(query string) error
}

type Scanner interface {
	Scan(dest ...interface{}) error
}

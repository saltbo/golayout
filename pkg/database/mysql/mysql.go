package mysql

type DB struct {
}

func NewDB(cfg *Config) (*DB, error) {

	return &DB{}, nil
}

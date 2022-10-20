package dao

type Book struct {
	*Dao
}

func NewBook() *Book {
	return &Book{}
}

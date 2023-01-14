package model

import "database/sql"

type booksRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *booksRepository {
	return &booksRepository{db}
}

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Published string `json:"published"`
}

func (r *booksRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	rows, err := r.db.Query(`select * from books;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Published)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	defer rows.Close()

	return books, nil
}

func (r *booksRepository) FindBook(id uint64) (*Book, error) {
	var book Book
	row := r.db.QueryRow(`select * from books where id = $1;`, id)

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Published)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *booksRepository) CreateBook(book Book) error {
	_, err := r.db.Exec(`insert into books (title, author, published) values ($1, $2, $3);`, book.Title, book.Author, book.Published)
	if err != nil {
		return err
	}

	return nil
}

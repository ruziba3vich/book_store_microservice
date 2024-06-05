package storage

import (
	"database/sql"
	"library/models"
	"log"
)

type (
	Book struct {
		db *sql.DB
	}
)

func NewBook(db *sql.DB) *Book {
	return &Book{
		db: db,
	}
}

func (a *Book) CreateNewAuthor(req models.BookRequest) (*models.BookResponse, error) {
	query := `
		INSERT INTO books (author_id, title, price, page) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, author_id, title, price, page
	`

	row := a.db.QueryRow(query, req.AuthorID, req.Title, req.Price, req.Page)

	var res models.BookResponse

	err := row.Scan(&res.ID, &res.AuthorID, &res.Title, &res.Price, &res.Page)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}
// Abduazim Yusufov
func (a *Book) GetBookByID(id int) (*models.BookResponse, error) {
	query := `
		Select id, author_id, title,price, page from books where id=$1 
	`
	var res models.BookResponse

	row := a.db.QueryRow(query, id)
	err := row.Scan(&res.ID, &res.AuthorID, &res.Title, &res.Price, &res.Page)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &res, nil
}

func (a *Book) GetbookByName(name string) (*[]models.BookResponse, error) {
	query := `
		Select id, author_id, title, price, page from books where title = $1
	`
	var (
		res []models.BookResponse
	)

	rows, err := a.db.Query(query, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var book models.BookResponse
		if err = rows.Scan(&book.ID, &book.AuthorID, &book.Title, &book.Price, &book.Page); err != nil {
			log.Println(err)
			return nil, err
		}
		res = append(res, book)

	}

	return &res, nil
}

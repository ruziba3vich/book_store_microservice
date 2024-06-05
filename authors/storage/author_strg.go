package storage

import (
	"authors/models"
	"database/sql"
	"log"
)

type (
	Author struct {
		db *sql.DB
	}
)

func NewAuthor(db *sql.DB) *Author {
	return &Author{
		db: db,
	}
}

func (a *Author) CreateNewAuthor(req models.AuthorRequest) (*models.AuthorResponse, error) {
	query := `
		INSERT INTO authors (age, u_name) 
		VALUES ($1, $2)
		RETURNING id, age, u_name
	`

	row := a.db.QueryRow(query, req.Age, req.Name)

	var res models.AuthorResponse

	err := row.Scan(&res.ID, &res.Age, &res.Name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

// Abduazim Yusufov
func (a *Author) GetAuthorByID(id int) (*models.AuthorResponse, error) {
	query := `
		Select id,age,u_name from authors where id=$1 
	`
	var res models.AuthorResponse

	row := a.db.QueryRow(query, id)
	if err := row.Scan(&res.ID, &res.Age, &res.Name); err != nil {
		log.Println(err)
		return nil, err
	}

	return &res, nil
}

func (a *Author) GetAuthorByName(name string) (*[]models.AuthorResponse, error) {
	query := `
		Select id,age,u_name from authors where u_name = $1
	`
	var (
		res []models.AuthorResponse
	)

	rows, err := a.db.Query(query, name)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var author models.AuthorResponse
		if err = rows.Scan(&author.ID, &author.Age, &author.Name); err != nil {
			log.Println(err)
			return nil, err
		}
		res = append(res, author)

	}

	return &res, nil
}

/*
CREATE TABLE IF NOT EXISTS authors(
    id SERIAL PRIMARY KEY,
    age INTEGER,
    u_name VARCHAR(64) UNIQUE
);
*/

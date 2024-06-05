package models

type (
	BookRequest struct {
		AuthorID int     `json:"author_id"`
		Title    string  `json:"title"`
		Price    float64 `json:"price"`
		Page     int     `json:"page"`
	}

	BookResponse struct {
		ID       int     `json:"id"`
		AuthorID int     `json:"author_id"`
		Title    string  `json:"title"`
		Price    float64 `json:"price"`
		Page     int     `json:"page"`
	}
)

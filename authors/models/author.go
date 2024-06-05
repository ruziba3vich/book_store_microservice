package models

type (
	AuthorRequest struct {
		Age  int    `json:"age"`
		Name string `json:"name"`
	}

	AuthorResponse struct {
		ID   int    `json:"id"`
		Age  int    `json:"age"`
		Name string `json:"name"`
	}
)

package utility

import "database/sql"

// Pagination is a default pagination structure.
type Pagination struct {
	Total       uint64 `json:"total"`
	PerPage     uint64 `json:"per_page"`
	CurrentPage uint64 `json:"current_page"`
	LastPage    uint64 `json:"last_page"`
}

// View is a default JSON response format.
type View struct {
	Response   bool        `json:"response"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

var (
	DB *sql.DB
)

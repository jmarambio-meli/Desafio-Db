package domain

type Product struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type TopProducts struct {
	Description string
	Total       int
}

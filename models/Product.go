package models

type Product struct {
	ID          int
	Title       string
	Description string
	Price       float64
	Quantity    int
	Active      bool
}

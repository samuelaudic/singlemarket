package models

import "time"

type Order struct {
	ID          int
	ClientID    int
	ProductID   int
	Quantity    int
	Price       float64
	PurchaseDate time.Time
}
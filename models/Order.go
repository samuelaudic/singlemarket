package models

type Order struct {
	id         int
	idclient   Client
	idproduct  Product
	quantity   int
	totalPrice float64
}
package controllers

import (
	"fmt"
	"singlemarket/models"
	"singlemarket/services"
	"singlemarket/views"
)

func PlaceOrder() {
	views.GetSeparator()
	var order models.Order
	clientID := views.GetTextInput("Enter client ID: ")
	fmt.Sscanf(clientID, "%d", &order.ClientID)
	productID := views.GetTextInput("Enter product ID: ")
	fmt.Sscanf(productID, "%d", &order.ProductID)
	quantity := views.GetTextInput("Enter quantity: ")
	fmt.Sscanf(quantity, "%d", &order.Quantity)

	product, err := services.GetProductByID(order.ProductID)
	if err != nil {
			fmt.Printf("Error getting product: %v\n", err)
			return
	}

	order.Price = product.Price * float64(order.Quantity)

	paymentInput := views.GetTextInput("Enter payment amount: ")
	var payment float64
	fmt.Sscanf(paymentInput, "%f", &payment)

	change := payment - order.Price
	if change < 0 {
			fmt.Println("Not enough payment provided.")
			return
	}

	err = services.PlaceOrder(order)
	if err != nil {
			fmt.Printf("Error placing order: %v\n", err)
	} else {
			fmt.Printf("Order placed successfully! Change: %.2f\n", change)
	}
}

func ExportOrdersToCSV() {
	views.GetSeparator()
	err := services.ExportAllOrdersToCSV()
	if err != nil {
		fmt.Printf("Error exporting orders to CSV: %v\n", err)
	} else {
		fmt.Println("Orders exported to CSV successfully!")
	}
}
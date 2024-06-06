package controllers

import (
	"fmt"
	"singlemarket/models"
	"singlemarket/services"
	"singlemarket/utils"
)

func PlaceOrder() {
	var order models.Order
	clientID := utils.GetTextInput("Enter client ID: ")
	fmt.Sscanf(clientID, "%d", &order.ClientID)
	productID := utils.GetTextInput("Enter product ID: ")
	fmt.Sscanf(productID, "%d", &order.ProductID)
	quantity := utils.GetTextInput("Enter quantity: ")
	fmt.Sscanf(quantity, "%d", &order.Quantity)

	product, err := services.GetProductByID(order.ProductID)
	if err != nil {
			fmt.Printf("Error getting product: %v\n", err)
			return
	}

	order.Price = product.Price * float64(order.Quantity)

	paymentInput := utils.GetTextInput("Enter payment amount: ")
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
	err := services.ExportAllOrdersToCSV()
	if err != nil {
		fmt.Printf("Error exporting orders to CSV: %v\n", err)
	} else {
		fmt.Println("Orders exported to CSV successfully!")
	}
}
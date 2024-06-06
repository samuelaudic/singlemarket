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
	price := utils.GetTextInput("Enter price: ")
	fmt.Sscanf(price, "%f", &order.Price)

	err := services.PlaceOrder(order)
	if err != nil {
		fmt.Printf("Error placing order: %v\n", err)
	} else {
		fmt.Println("Order placed successfully!")
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

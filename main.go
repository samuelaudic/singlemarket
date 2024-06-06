package main

import (
	"fmt"
	"log"
	"singlemarket/controllers"
	"singlemarket/database"
)

func main() {
	// db
	err := database.Connect()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	fmt.Println("Database connection successful and migration completed")

	// menu
	for {
		fmt.Println("Main Menu")
		fmt.Println("1. Add a product")
		fmt.Println("2. View all products")
		fmt.Println("3. Edit a product")
		fmt.Println("4. Delete a product")
		fmt.Println("5. Export products to CSV")
		fmt.Println("6. Add a client")
		fmt.Println("7. View all clients")
		fmt.Println("8. Edit a client")
		fmt.Println("9. Export clients to CSV")
		fmt.Println("10. Place an order")
		fmt.Println("11. Export orders to CSV")
		fmt.Println("12. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		fmt.Scanln()

		switch choice {
		case 1:
			controllers.AddProduct()
		case 2:
			controllers.ViewAllProducts()
		case 3:
			controllers.EditProduct()
		case 4:
			controllers.DeleteProduct()
		case 5:
			controllers.ExportProductsToCSV()
		case 6:
			controllers.AddClient()
		case 7:
			controllers.ViewAllClients()
		case 8:
			controllers.EditClient()
		case 9:
			controllers.ExportClientsToCSV()
		case 10:
			controllers.PlaceOrder()
		case 11:
			controllers.ExportOrdersToCSV()
		case 12:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
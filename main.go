package main

import (
	"fmt"
	"log"
	"singlemarket/database"
	"singlemarket/handlers"
)

func main() {
	// Test database connection and migration
	err := database.Connect()
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}
	fmt.Println("Database connection successful and migration completed")

	// Display main menu
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

		switch choice {
		case 1:
			handlers.AddProduct()
		case 2:
			handlers.ViewAllProducts()
		case 3:
			handlers.EditProduct()
		case 4:
			handlers.DeleteProduct()
		case 5:
			handlers.ExportProductsToCSV()
		case 6:
			// Add a client
		case 7:
			// View all clients
		case 8:
			// Edit a client
		case 9:
			// Export clients to CSV
		case 10:
			// Place an order
		case 11:
			// Export orders to CSV
		case 12:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}
package handlers

import (
	"fmt"
	"singlemarket/models"
	"singlemarket/services"
)

func AddProduct() {
	var title, description string
	var price float64
	var quantity int

	fmt.Print("Enter product title: ")
	fmt.Scan(&title)
	fmt.Print("Enter product description: ")
	fmt.Scan(&description)
	fmt.Print("Enter product price: ")
	fmt.Scan(&price)
	fmt.Print("Enter product quantity: ")
	fmt.Scan(&quantity)

	product := models.Product{
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		Active:      true,
	}

	err := services.AddProduct(product)
	if err != nil {
		fmt.Printf("Error adding product: %v\n", err)
	} else {
		fmt.Println("Product added successfully!")
	}
}

func ViewAllProducts() {
	products, err := services.GetAllProducts()
	if err != nil {
		fmt.Printf("Error retrieving products: %v\n", err)
		return
	}

	for _, product := range products {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Price: %.2f, Quantity: %d, Active: %t\n",
			product.ID, product.Title, product.Description, product.Price, product.Quantity, product.Active)
	}
}

func EditProduct() {
	var id int
	var title, description string
	var price float64
	var quantity int
	var active bool

	fmt.Print("Enter product ID to edit: ")
	fmt.Scan(&id)
	fmt.Print("Enter new product title: ")
	fmt.Scan(&title)
	fmt.Print("Enter new product description: ")
	fmt.Scan(&description)
	fmt.Print("Enter new product price: ")
	fmt.Scan(&price)
	fmt.Print("Enter new product quantity: ")
	fmt.Scan(&quantity)
	fmt.Print("Is the product active? (true/false): ")
	fmt.Scan(&active)

	product := models.Product{
		ID:          id,
		Title:       title,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		Active:      active,
	}

	err := services.UpdateProduct(product)
	if err != nil {
		fmt.Printf("Error updating product: %v\n", err)
	} else {
		fmt.Println("Product updated successfully!")
	}
}

func DeleteProduct() {
	var id int

	fmt.Print("Enter product ID to delete: ")
	fmt.Scan(&id)

	err := services.DeactivateProduct(id)
	if err != nil {
		fmt.Printf("Error deleting product: %v\n", err)
	} else {
		fmt.Println("Product deleted successfully!")
	}
}

func ExportProductsToCSV() {
	err := services.ExportAllProductsToCSV()
	if err != nil {
		fmt.Printf("Error exporting products to CSV: %v\n", err)
	} else {
		fmt.Println("Products exported to CSV successfully!")
	}
}

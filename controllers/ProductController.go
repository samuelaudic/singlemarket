package controllers

import (
	"fmt"
	"singlemarket/models"
	"singlemarket/services"
	"singlemarket/views"
)

func AddProduct() {
	views.GetSeparator()
	var product models.Product
	product.Title = views.GetTextInput("Enter product title: ")
	product.Description = views.GetTextInput("Enter product description: ")
	price := views.GetTextInput("Enter product price: ")
	fmt.Sscanf(price, "%f", &product.Price)
	quantity := views.GetTextInput("Enter product quantity: ")
	fmt.Sscanf(quantity, "%d", &product.Quantity)
	product.Active = true

	err := services.AddProduct(product)
	if err != nil {
		fmt.Printf("Error adding product: %v\n", err)
	} else {
		fmt.Println("Product added successfully!")
	}
}

func ViewAllProducts() {
	views.GetSeparator()
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
	views.GetSeparator()
	var id int
	fmt.Print("Enter product ID to edit: ")
	fmt.Scan(&id)

	fmt.Scanln()
	title := views.GetTextInput("Enter new product title: ")
	description := views.GetTextInput("Enter new product description: ")
	price := views.GetTextInput("Enter new product price: ")
	var priceFloat float64
	fmt.Sscanf(price, "%f", &priceFloat)
	quantity := views.GetTextInput("Enter new product quantity: ")
	var quantityInt int
	fmt.Sscanf(quantity, "%d", &quantityInt)
	active := views.GetTextInput("Is the product active? (true/false): ")
	var activeBool bool
	fmt.Sscanf(active, "%t", &activeBool)

	product := models.Product{
		ID:          id,
		Title:       title,
		Description: description,
		Price:       priceFloat,
		Quantity:    quantityInt,
		Active:      activeBool,
	}

	err := services.UpdateProduct(product)
	if err != nil {
		fmt.Printf("Error updating product: %v\n", err)
	} else {
		fmt.Println("Product updated successfully!")
	}
}

func DeleteProduct() {
	views.GetSeparator()
	var id int
	fmt.Print("Enter product ID to delete: ")
	fmt.Scan(&id)

	fmt.Scanln()

	err := services.DeactivateProduct(id)
	if err != nil {
		fmt.Printf("Error deleting product: %v\n", err)
	} else {
		fmt.Println("Product deleted successfully!")
	}
}

func ExportProductsToCSV() {
	views.GetSeparator()
	err := services.ExportAllProductsToCSV()
	if err != nil {
		fmt.Printf("Error exporting products to CSV: %v\n", err)
	} else {
		fmt.Println("Products exported to CSV successfully!")
	}
}

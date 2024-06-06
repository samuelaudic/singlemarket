package services

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"singlemarket/database"
	"singlemarket/models"
)

// AddProduct adds a new product to the database
func AddProduct(product models.Product) error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO products (title, description, price, quantity, active) VALUES (?, ?, ?, ?, ?)", product.Title, product.Description, product.Price, product.Quantity, product.Active)
	return err
}

// GetAllProducts retrieves all products from the database
func GetAllProducts() ([]models.Product, error) {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, title, description, price, quantity, active FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Quantity, &product.Active)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

// UpdateProduct updates an existing product in the database
func UpdateProduct(product models.Product) error {
	db := database.GetDB()
	_, err := db.Exec("UPDATE products SET title = ?, description = ?, price = ?, quantity = ?, active = ? WHERE id = ?", product.Title, product.Description, product.Price, product.Quantity, product.Active, product.ID)
	return err
}

// DeactivateProduct deactivates a product by setting its active field to false
func DeactivateProduct(id int) error {
	db := database.GetDB()
	_, err := db.Exec("UPDATE products SET active = ? WHERE id = ?", false, id)
	return err
}

// ExportAllProductsToCSV exports all products to a CSV file
func ExportAllProductsToCSV() error {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, title, description, price, quantity, active FROM products")
	if err != nil {
		return err
	}
	defer rows.Close()

	// Create or open CSV file
	csvFile, err := os.Create("../products.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"id", "title", "description", "price", "quantity", "active"})

	// Write rows
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Quantity, &product.Active)
		if err != nil {
			return err
		}
		record := []string{
			fmt.Sprintf("%d", product.ID),
			product.Title,
			product.Description,
			fmt.Sprintf("%.2f", product.Price),
			fmt.Sprintf("%d", product.Quantity),
			fmt.Sprintf("%t", product.Active),
		}
		writer.Write(record)
	}

	return nil
}

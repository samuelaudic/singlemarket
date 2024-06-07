package services

import (
	"encoding/csv"
	"fmt"
	"os"
	"singlemarket/database"
	"singlemarket/models"
)

func AddProduct(product models.Product) error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO products (title, description, price, quantity, active) VALUES (?, ?, ?, ?, ?)", product.Title, product.Description, product.Price, product.Quantity, product.Active)
	return err
}

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

func UpdateProduct(product models.Product) error {
	db := database.GetDB()
	_, err := db.Exec("UPDATE products SET title = ?, description = ?, price = ?, quantity = ?, active = ? WHERE id = ?", product.Title, product.Description, product.Price, product.Quantity, product.Active, product.ID)
	return err
}

func DeactivateProduct(id int) error {
	db := database.GetDB()
	_, err := db.Exec("UPDATE products SET active = ? WHERE id = ?", false, id)
	return err
}

func ExportAllProductsToCSV() error {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, title, description, price, quantity, active FROM products")
	if err != nil {
		return err
	}
	defer rows.Close()

	csvFile, err := os.Create("./exports/csv/products.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.Write([]string{"id", "title", "description", "price", "quantity", "active"})

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

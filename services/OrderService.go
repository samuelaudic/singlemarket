package services

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"singlemarket/database"
	"singlemarket/models"
	"singlemarket/utils"
	"time"
)

func PlaceOrder(order models.Order) error {
	db := database.GetDB()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO orders (client_id, product_id, quantity, price, purchase_date) VALUES (?, ?, ?, ?, ?)",
		order.ClientID, order.ProductID, order.Quantity, order.Price, time.Now())
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	client, err := GetClientByID(order.ClientID)
	if err != nil {
		return err
	}
	product, err := GetProductByID(order.ProductID)
	if err != nil {
		return err
	}
	emailBody := fmt.Sprintf("Dear %s %s,\n\nThank you for your order!\n\nProduct: %s\nQuantity: %d\nPrice: %.2f\n\nBest regards,\nSingleMarket",
		client.FirstName, client.LastName, product.Title, order.Quantity, order.Price)

		fmt.Println("envoi du mail + génération du pdf")
		fmt.Println(emailBody)

	err = utils.SendEmail(client.Email, "Your Order Confirmation", emailBody)
	if err != nil {
		return err
	}

	// err = utils.GenerateOrderPDF(order, client, product)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func GetClientByID(id int) (models.Client, error) {
	db := database.GetDB()
	row := db.QueryRow("SELECT id, first_name, last_name, phone, address, email FROM clients WHERE id = ?", id)

	var client models.Client
	err := row.Scan(&client.ID, &client.FirstName, &client.LastName, &client.Phone, &client.Address, &client.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return client, fmt.Errorf("client not found")
		}
		return client, err
	}

	return client, nil
}

func GetProductByID(id int) (models.Product, error) {
	db := database.GetDB()
	row := db.QueryRow("SELECT id, title, description, price, quantity, active FROM products WHERE id = ?", id)

	var product models.Product
	err := row.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.Quantity, &product.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, fmt.Errorf("product not found")
		}
		return product, err
	}

	return product, nil
}

func ExportAllOrdersToCSV() error {
	db := database.GetDB()
	rows, err := db.Query("SELECT orders.id, clients.first_name, clients.last_name, products.title, orders.quantity, orders.price, orders.purchase_date FROM orders JOIN clients ON orders.client_id = clients.id JOIN products ON orders.product_id = products.id")
	if err != nil {
		return err
	}
	defer rows.Close()

	csvFile, err := os.Create("./exports/csv/orders.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.Write([]string{"Order ID", "Client First Name", "Client Last Name", "Product Title", "Quantity", "Price", "Purchase Date"})

	for rows.Next() {
    var orderID int
    var clientFirstName, clientLastName, productTitle string
    var quantity int
    var price float64
    var purchaseDateStr string

    err := rows.Scan(&orderID, &clientFirstName, &clientLastName, &productTitle, &quantity, &price, &purchaseDateStr)
    if err != nil {
        return err
    }

    purchaseDate, err := time.Parse("2006-01-02 15:04:05", purchaseDateStr)
		if err != nil {
				return err
		}

    record := []string{
        fmt.Sprintf("%d", orderID),
        clientFirstName,
        clientLastName,
        productTitle,
        fmt.Sprintf("%d", quantity),
        fmt.Sprintf("%.2f", price),
        purchaseDate.Format(time.RFC3339),
    }
    writer.Write(record)
	}
	
	return nil
}

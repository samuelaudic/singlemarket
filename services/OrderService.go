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

	result, err := tx.Exec("INSERT INTO orders (client_id, product_id, quantity, price, purchase_date) VALUES (?, ?, ?, ?, ?)",
		order.ClientID, order.ProductID, order.Quantity, order.Price, time.Now())
	if err != nil {
		return err
	}

	orderID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	order.ID = int(orderID)

	// Retrieve the order to ensure all fields are populated correctly
	var purchaseDateStr string
	err = tx.QueryRow("SELECT purchase_date FROM orders WHERE id = ?", order.ID).Scan(&purchaseDateStr)
	if err != nil {
		return err
	}

	// Convert the string to time.Time
	order.PurchaseDate, err = time.Parse("2006-01-02 15:04:05", purchaseDateStr)
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

	err = utils.GenerateOrderPDF(order, client, product)
	if err != nil {
		return err
	}

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
	rows, err := db.Query(`
		SELECT 
			clients.id, 
			clients.first_name, 
			clients.last_name, 
			GROUP_CONCAT(orders.id) AS order_ids, 
			SUM(orders.price * orders.quantity) AS total_price
		FROM 
			orders 
		JOIN 
			clients ON orders.client_id = clients.id 
		GROUP BY 
			clients.id, clients.first_name, clients.last_name
	`)
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

	// Write header
	writer.Write([]string{"Client ID", "Client First Name", "Client Last Name", "Order IDs", "Total Price"})

	for rows.Next() {
		var clientID int
		var clientFirstName, clientLastName, orderIDs string
		var totalPrice float64

		err := rows.Scan(&clientID, &clientFirstName, &clientLastName, &orderIDs, &totalPrice)
		if err != nil {
			return err
		}

		record := []string{
			fmt.Sprintf("%d", clientID),
			clientFirstName,
			clientLastName,
			orderIDs,
			fmt.Sprintf("%.2f", totalPrice),
		}
		writer.Write(record)
	}

	return nil
}
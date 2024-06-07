package services

import (
	"encoding/csv"
	"fmt"
	"os"
	"singlemarket/database"
	"singlemarket/models"
)

func AddClient(client models.Client) error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO clients (first_name, last_name, phone, address, email) VALUES (?, ?, ?, ?, ?)",
		client.FirstName, client.LastName, client.Phone, client.Address, client.Email)
	return err
}

func GetAllClients() ([]models.Client, error) {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, first_name, last_name, phone, address, email FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clients := []models.Client{}
	for rows.Next() {
		var client models.Client
		err := rows.Scan(&client.ID, &client.FirstName, &client.LastName, &client.Phone, &client.Address, &client.Email)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func UpdateClient(client models.Client) error {
	db := database.GetDB()
	_, err := db.Exec("UPDATE clients SET first_name = ?, last_name = ?, phone = ?, address = ?, email = ? WHERE id = ?", client.FirstName, client.LastName, client.Phone, client.Address, client.Email, client.ID)
	return err
}

func DeactivateClient(id int) error {
	db := database.GetDB()
	_, err := db.Exec("UPDATE clients SET active = ? WHERE id = ?", false, id)
	return err
}

func ExportAllClientsToCSV() error {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, first_name, last_name, phone, address, email FROM clients")
	if err != nil {
		return err
	}
	defer rows.Close()

	csvFile, err := os.Create("./exports/csv/clients.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	writer.Write([]string{"id", "first_name", "last_name", "phone", "address", "email"})

	for rows.Next() {
		var client models.Client
		err := rows.Scan(&client.ID, &client.FirstName, &client.LastName, &client.Phone, &client.Address, &client.Email)
		if err != nil {
			return err
		}
		record := []string{
			fmt.Sprintf("%d", client.ID),
			client.FirstName,
			client.LastName,
			client.Phone,
			client.Address,
			client.Email,
		}
		writer.Write(record)
	}

	return nil
}
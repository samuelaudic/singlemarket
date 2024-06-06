package services

import (
	"encoding/csv"
	"os"
	"singlemarket/database"
	"strconv"
)

func AddClient(client Client) error {
	db := database.GetDB()
	_, err := db.Exec("INSERT INTO clients (firstname, lastname, age, email, phone) VALUES (?, ?, ?, ?, ?)", client.Firstname, client.Lastname, client.Age, client.Email, client.Phone)
	return err
}

func GetAllClients() ([]Client, error) {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, firstname, lastname, age, email, phone FROM clients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clients := []Client{}
	for rows.Next() {
		var client Client
		err := rows.Scan(&client.ID, &client.Firstname, &client.Lastname, &client.Age, &client.Email, &client.Phone)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}

func UpdateClient(client Client) error {
	db := database.GetDB()
	_, err := db.Exec("UPDATE clients SET firstname = ?, lastname = ?, age = ?, email = ?, phone = ? WHERE id = ?", client.Firstname, client.Lastname, client.Age, client.Email, client.Phone, client.ID)
	return err
}

func ExportAllClientsToCSV() error {
	db := database.GetDB()
	rows, err := db.Query("SELECT id, firstname, lastname, age, email, phone FROM clients")
	if err != nil {
		return err
	}
	defer rows.Close()

	file, err := os.Create("../clients.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for rows.Next() {
		var client Client
		err := rows.Scan(&client.ID, &client.Firstname, &client.Lastname, &client.Age, &client.Email, &client.Phone)
		if err != nil {
			return err
		}

		err = writer.Write([]string{strconv.Itoa(client.ID), client.Firstname, client.Lastname, strconv.Itoa(client.Age), client.Email, client.Phone})
		if err != nil {
			return err
		}
	}

	return nil
}
package handlers

import (
	"fmt"
	"singlemarket/models"
	"singlemarket/services"
)

func AddClient() {
	fmt.Println("Add a client")
	var client models.Client
	fmt.Print("First name: ")
	fmt.Scan(&client.FirstName)
	fmt.Print("Last name: ")
	fmt.Scan(&client.LastName)
	fmt.Print("Phone: ")
	fmt.Scan(&client.Phone)
	fmt.Print("Address: ")
	fmt.Scan(&client.Address)
	fmt.Print("Email: ")
	fmt.Scan(&client.Email)

	err := services.AddClient(client)
	if err != nil {
		fmt.Printf("Error adding client: %v\n", err)
	} else {
		fmt.Println("Client added successfully")
	}
}

func ViewAllClients() {
	clients, err := services.GetAllClients()
	if err != nil {
		fmt.Printf("Error fetching clients: %v\n", err)
	} else {
		fmt.Println("All clients:")
		for _, client := range clients {
			fmt.Printf("ID: %d, Name: %s %s, Phone: %s, Address: %s, Email: %s\n", client.ID, client.FirstName, client.LastName, client.Phone, client.Address, client.Email)
		}
	}
}

func EditClient() {
	fmt.Println("Edit a client")
	var client models.Client
	fmt.Print("Enter client ID: ")
	fmt.Scan(&client.ID)
	fmt.Print("First name: ")
	fmt.Scan(&client.FirstName)
	fmt.Print("Last name: ")
	fmt.Scan(&client.LastName)
	fmt.Print("Phone: ")
	fmt.Scan(&client.Phone)
	fmt.Print("Address: ")
	fmt.Scan(&client.Address)
	fmt.Print("Email: ")
	fmt.Scan(&client.Email)

	err := services.UpdateClient(client)
	if err != nil {
		fmt.Printf("Error updating client: %v\n", err)
	} else {
		fmt.Println("Client updated successfully")
	}
}

func ExportClientsToCSV() {
	err := services.ExportAllClientsToCSV()
	if err != nil {
		fmt.Printf("Error exporting clients to CSV: %v\n", err)
	} else {
		fmt.Println("Clients exported to clients.csv")
	}
}
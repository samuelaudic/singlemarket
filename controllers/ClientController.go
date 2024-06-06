package controllers

import (
	"fmt"
	"singlemarket/models"
	"singlemarket/services"
	"singlemarket/utils"
)

func AddClient() {
	fmt.Println("Add a client")
	var client models.Client
	client.FirstName = utils.GetTextInput("First name: ")
	client.LastName = utils.GetTextInput("Last name: ")
	client.Phone = utils.GetTextInput("Phone: ")
	client.Address = utils.GetTextInput("Address: ")
	client.Email = utils.GetTextInput("Email: ")

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
		fmt.Printf("Error retrieving clients: %v\n", err)
		return
	}

	for _, client := range clients {
		fmt.Printf("ID: %d, First Name: %s, Last Name: %s, Phone: %s, Address: %s, Email: %s\n",
			client.ID, client.FirstName, client.LastName, client.Phone, client.Address, client.Email)
	}
}

func EditClient() {
	var id int
	fmt.Print("Enter client ID to edit: ")
	fmt.Scan(&id)
	firstName := utils.GetTextInput("Enter new client first name: ")
	lastName := utils.GetTextInput("Enter new client last name: ")
	phone := utils.GetTextInput("Enter new client phone: ")
	address := utils.GetTextInput("Enter new client address: ")
	email := utils.GetTextInput("Enter new client email: ")

	client := models.Client{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		Address:   address,
		Email:     email,
	}

	err := services.UpdateClient(client)
	if err != nil {
		fmt.Printf("Error updating client: %v\n", err)
	} else {
		fmt.Println("Client updated successfully!")
	}
}

func ExportClientsToCSV() {
	err := services.ExportAllClientsToCSV()
	if err != nil {
		fmt.Printf("Error exporting clients to CSV: %v\n", err)
	} else {
		fmt.Println("Clients exported to CSV successfully!")
	}
}

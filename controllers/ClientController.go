package controllers

import (
	"fmt"
	"singlemarket/models"
	"singlemarket/services"
	"singlemarket/views"
)

func AddClient() {
	views.GetSeparator()
	fmt.Println("Add a client")
	var client models.Client
	client.FirstName = views.GetTextInput("First name: ")
	client.LastName = views.GetTextInput("Last name: ")
	client.Phone = views.GetTextInput("Phone: ")
	client.Address = views.GetTextInput("Address: ")
	client.Email = views.GetTextInput("Email: ")

	err := services.AddClient(client)
	if err != nil {
		fmt.Printf("Error adding client: %v\n", err)
	} else {
		fmt.Println("Client added successfully")
	}
}

func ViewAllClients() {
	views.GetSeparator()
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
	views.GetSeparator()
	fmt.Print("Enter client ID to edit: ")
	fmt.Scan(&id)
	// Clear the buffer after scanning an integer
	fmt.Scanln()
	firstName := views.GetTextInput("Enter new client first name: ")
	lastName := views.GetTextInput("Enter new client last name: ")
	phone := views.GetTextInput("Enter new client phone: ")
	address := views.GetTextInput("Enter new client address: ")
	email := views.GetTextInput("Enter new client email: ")

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
	views.GetSeparator()
	err := services.ExportAllClientsToCSV()
	if err != nil {
		fmt.Printf("Error exporting clients to CSV: %v\n", err)
	} else {
		fmt.Println("Clients exported to CSV successfully!")
	}
}

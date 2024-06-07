package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"singlemarket/models"

	"github.com/jung-kurt/gofpdf"
)

func GenerateOrderPDF(order models.Order, client models.Client, product models.Product) error {
	// Create the directory if it doesn't exist
	outputDir := "./exports/pdf"
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		err := os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	// Set the output file path
	outputFile := filepath.Join(outputDir, fmt.Sprintf("order_%d.pdf", order.ID))

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, fmt.Sprintf("Order ID: %d", order.ID))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Client: %s %s", client.FirstName, client.LastName))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Product: %s", product.Title))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Quantity: %d", order.Quantity))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Price: %.2f", order.Price))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Date: %s", FormatDateInEnglish(order.PurchaseDate)))

	err := pdf.OutputFileAndClose(outputFile)
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %v", err)
	}

	return nil
}
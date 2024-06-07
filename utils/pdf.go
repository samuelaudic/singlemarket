package utils

import (
	"fmt"
	"singlemarket/models"

	"github.com/jung-kurt/gofpdf"
)

func GenerateOrderPDF(order models.Order, client models.Client, product models.Product) error {
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
	pdf.Cell(40, 10, fmt.Sprintf("Date: %s", order.PurchaseDate.Format("2006-01-02 15:04:05")))

	err := pdf.OutputFileAndClose(fmt.Sprintf("order_%d.pdf", order.ID))
	if err != nil {
		return fmt.Errorf("failed to generate PDF: %v", err)
	}

	return nil
}
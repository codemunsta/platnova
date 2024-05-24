package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/codemunsta/platnova/utils"
	"github.com/jung-kurt/gofpdf"
)

func main() {
	// Get Mockup Data
	data, err := utils.MockUpData()
	if err != nil {
		fmt.Println(err)
	}
	
	// Initialize PDF file
	pdf := gofpdf.New(gofpdf.OrientationPortrait, "mm", "A4", "")
	margin := 10.0
	pdf.SetMargins(margin, margin, margin+10)
	pdf.AddPage()

	// Add Logo
	logoPath := "images/Revolut.png"
	logoWidth := 30.0
	logoHeight := 20.0
	pdf.Image(logoPath, margin+2, margin+2, logoWidth, logoHeight, false, "", 0, "")
	
	// EUR Header
	textX := 210.0 - margin
	textY := margin
	pdf.SetFont("Arial", "B", 16)
	pdf.SetXY(textX, textY+4)
	pdf.CellFormat(0, 10, "EUR Statement", "", 1, "R", false, 0, "")
	pdf.SetFont("Arial", "", 8)
	currentDate := time.Now().Format("02 January 2006")
	pdf.SetXY(textX, textY+8)
	pdf.CellFormat(0, 10, "Generated on the "+currentDate, "", 1, "R", false, 0, "")
	pdf.SetXY(textX, textY+11)
	pdf.CellFormat(0, 10, "Revolut Bank UAB", "", 1, "R", false, 0, "")

	// User address
	pdf.SetFont("Arial", "B", 11)
	pdf.SetXY(margin+4, margin+logoHeight+5)
	pdf.CellFormat(0, 10, strings.ToUpper(data.User.Name), "", 1, "L", false, 0, "")

	// Split the address by commas or periods and add each line
	pdf.SetFont("Arial", "", 8)
	addressLines := strings.FieldsFunc(data.User.Address, func(r rune) bool {
		return r == ',' || r == '.'
	})
	pdf.SetXY(margin+4, margin+logoHeight+11)
	offset := margin+logoHeight+15
	for _, line := range addressLines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			pdf.CellFormat(0, 10, trimmedLine, "", 1, "L", false, 0, "")
		}
		pdf.SetXY(margin+4, offset)
		offset += 4
	}

	// Bank details
	for _, bank := range data.Iban {
		pdf.SetFont("Arial", "B", 8)
		pdf.SetXY(margin+105, offset+4)
		pdf.CellFormat(0, 10, strings.ToUpper("IBAN"), "", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "", 8)
		pdf.SetXY(margin+120, offset+4)
		pdf.CellFormat(0, 10, strings.ToUpper(bank.Iban), "", 1, "L", false, 0, "")
		offset += 4

		// BIC
		pdf.SetFont("Arial", "B", 8)
		pdf.SetXY(margin+105, offset+4)
		pdf.CellFormat(0, 10, strings.ToUpper("BIC"), "", 1, "L", false, 0, "")
		pdf.SetFont("Arial", "", 8)
		pdf.SetXY(margin+120, offset+4)
		pdf.CellFormat(0, 10, strings.ToUpper(bank.Bic), "", 1, "L", false, 0, "")
		offset += 8
	}

	pdf.Ln(-1)
	pdf.SetFont("Arial", "B", 11)
	pdf.SetX(margin+4)
	pdf.CellFormat(0, 10, "Balance summary", "", 1, "L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(margin+5)
	pdf.SetFont("Arial", "B", 8)
	headers := []string{"Product", "Opening Balance", "Money Out", "Money In", "Closing Balance"}
	for _, header := range headers {
		if header == "Product" {
			pdf.CellFormat(68, 9, header, "B", 0, "L", false, 0, "")
		} else if header == "Closing Balance" {
			pdf.CellFormat(34, 9, header, "B", 0, "R", false, 0, "")
		} else {
			pdf.CellFormat(26, 9, header, "B", 0, "L", false, 0, "")
		}
	}
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 9)
	for _, product := range data.Products {
		pdf.SetX(margin+5)
		pdf.CellFormat(68, 8, product.Product, "B", 0, "L", false, 0, "")
		pdf.CellFormat(26, 8, product.Balance, "B", 0, "L", false, 0, "")
		pdf.CellFormat(26, 8, product.MoneyOut, "B", 0, "L", false, 0, "")
		pdf.CellFormat(26, 8, product.MoneyIn, "B", 0, "L", false, 0, "")
		pdf.CellFormat(34, 8, product.ClosingBalance, "B", 0, "R", false, 0, "")
		pdf.Ln(-1)
	}

	pdf.SetX(margin+5)
	pdf.CellFormat(68, 8, "Total", "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 8, fmt.Sprintf("$%.2f", 2.52), "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 8, fmt.Sprintf("$%.2f", 1944.09), "", 0, "L", false, 0, "")
	pdf.CellFormat(26, 8, fmt.Sprintf("$%.2f", 1944.09), "", 0, "L", false, 0, "")
	pdf.CellFormat(34, 8, fmt.Sprintf("$%.2f", 36.43), "", 0, "R", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 5)
	pdf.SetX(margin+4)
	pdf.CellFormat(0, 10, "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley ", "", 1, "L", false, 0, "")
	pdf.Ln(-1)

	// transactions
	pdf.SetFont("Arial", "B", 11)
	pdf.SetX(margin+4)
	pdf.CellFormat(0, 10, "Account transactions from 1 February 2023 to 29 March 2023", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 8)
	pdf.SetX(margin+5)

	// Table headers
	headers = []string{"Date", "Description", "Money Out", "Money In", "Balance"}
	for _, header := range headers {
		if header == "Description" {
			pdf.CellFormat(68, 9, header, "B", 0, "L", false, 0, "")
		} else if header == "Balance" {
			pdf.CellFormat(38, 9, header, "B", 0, "R", false, 0, "")
		} else {
			if header == "Date"{
				pdf.CellFormat(24, 9, header, "B", 0, "L", false, 0, "")
			} else {
				pdf.CellFormat(24, 9, header, "B", 0, "L", false, 0, "")
			}
		}
	}
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 9)
	for _, transaction := range data.Transactions {
		pdf.SetX(margin+5)
		pdf.CellFormat(24, 8, transaction.Date, "B", 0, "L", false, 0, "")
		pdf.CellFormat(68, 8, transaction.Description, "B", 0, "L", false, 0, "")
		pdf.CellFormat(24, 8, transaction.MoneyOut, "B", 0, "L", false, 0, "")
		pdf.CellFormat(24, 8, transaction.MoneyIn, "B", 0, "L", false, 0, "")
		pdf.CellFormat(38, 8, transaction.Balance, "B", 0, "R", false, 0, "")
		pdf.Ln(-1)
	}

	pdf.Ln(-1)
	pdf.Ln(-1)

	// footer

	// qr code
	logoPath = "images/qr.png"
	logoWidth = 12.0
	logoHeight = 12.0
	pdf.Image(logoPath, margin+5, margin+240, logoWidth, logoHeight, false, "", 0, "")

	pdf.SetFont("Arial", "", 5)
	pdf.SetX(margin+50)
	pdf.MultiCell(140, 3, "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque in convallis ex. Praesent condimentum ultricies finibus. Duis neque velit, dignissim vitae augue nec, bibendum gravida lectus. Proin vitae mauris pellentesque, tristique purus eget, tempor metus. Nunc tincidunt dui metus, sit amet vehicula dolor varius a. Suspendisse diam purus, ultricies eu lacinia sit amet, iaculis non tellus. Sed tempus vehicula tortor nec posuere. Suspendisse porttitor posuere metus. Nulla enim nibh, pretium interdum erat a, ullamcorper convallis lorem. Proin efficitur, nulla quis ullamcorper pretium, turpis ligula molestie sapien, nec scelerisque mauris nisl quis enim. Praesent mi dolor, gravida eget velit a, convallis venenatis sem. Mauris ut massa sapien. Etiam molestie ipsum eu rhoncus semper.", "", "L", false)

	// Save the PDF to a file
	err = pdf.OutputFileAndClose("pdf/platnova.pdf")
	if err != nil {
		panic(err)
	}
}
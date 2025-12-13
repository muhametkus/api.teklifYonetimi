package main

import (
	"bytes"
	"log"
	"path/filepath"
	"text/template"

	"api.teklifYonetimi/internal/models"
	"api.teklifYonetimi/internal/utils"
)

func main() {

	// 1️⃣ Fake quotation data (test için)
	quotation := models.Quotation{
		Title:    "Web Sitesi Teklifi",
		Customer: "ABC Ltd",
		Status:   "PENDING",
		Total:    13000,
		Items: []models.QuotationItem{
			{
				ItemName:  "Tasarım",
				Quantity:  1,
				UnitPrice: 5000,
				Total:     5000,
			},
			{
				ItemName:  "Geliştirme",
				Quantity:  1,
				UnitPrice: 8000,
				Total:     8000,
			},
		},
	}

	// 2️⃣ Template dosyasını bul
	templatePath, _ := filepath.Abs("internal/templates/quotation.html")

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	// 3️⃣ Template render et
	var htmlBuffer bytes.Buffer
	err = tmpl.Execute(&htmlBuffer, quotation)
	if err != nil {
		log.Fatal(err)
	}

	// 4️⃣ HTML → PDF
	err = utils.GeneratePDFFromHTML(
		htmlBuffer.String(),
		"quotation_test.pdf",
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("PDF oluşturuldu (template render edildi)")
}

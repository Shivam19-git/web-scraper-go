package main

import (
	"fmt"
	"log"

	// importing colly
	"github.com/gocolly/colly"
	"github.com/xuri/excelize/v2"
)

/*
Colly's main entity is the Collector. A Collector allows you to perform HTTP requests. Also, it gives you access to the web scraping callbacks offered by the Colly interface.
*/

type ProductData struct {
	url   string
	image string
	name  string
	price string
}

func main() {
	// scraping logic...
	scrapeURL := "https://www.scrapingcourse.com/ecommerce/"
	c := colly.NewCollector(colly.AllowedDomains("www.scrapingcourse.com", "scrapingcourse.com"))

	// Create a slice to hold the product data
	var products []ProductData

	// Select all li.product HTML product elements in the page using Colly
	c.OnHTML("li.product", func(h *colly.HTMLElement) {
		product := ProductData{
			url:   h.ChildAttr("a", "href"),
			image: h.ChildAttr("img", "src"),
			name:  h.ChildText("h2"),
			price: h.ChildText(".price"),
		}
		products = append(products, product)
		
		fmt.Println("Name : ", product.name)
		fmt.Println("URL : ", product.url)
		fmt.Println("Image : ", product.image)
		fmt.Println("Price : ", product.price)
		fmt.Print("\n")
	})

	// Visiting a web page using Visit()
	c.Visit(scrapeURL)
	
	if len(products) == 0 {
		return
	}
	// To import the scraped data into an Excel sheet, you can use the excelize package in Go. This package allows you to create and manipulate Excel files.

	// Create a new Excel file
	file := excelize.NewFile()
	sheet := "Sheet1"

	// Set the header
	file.SetCellValue(sheet, "A1", "Name")
	file.SetCellValue(sheet, "B1", "URL")
	file.SetCellValue(sheet, "C1", "Image")
	file.SetCellValue(sheet, "D1", "Price")

	// Populate the sheet with the product data
	for i, product := range products {
		file.SetCellValue(sheet, fmt.Sprintf("A%d", i+2), product.name)
		file.SetCellValue(sheet, fmt.Sprintf("B%d", i+2), product.url)
		file.SetCellValue(sheet, fmt.Sprintf("C%d", i+2), product.image)
		file.SetCellValue(sheet, fmt.Sprintf("D%d", i+2), product.price)
	}

	// Save the file
	if err := file.SaveAs("products.xlsx"); err != nil {
		log.Fatalf("Failed to save file: %v", err)
	}

	fmt.Println("Data successfully written to products.xlsx")
}

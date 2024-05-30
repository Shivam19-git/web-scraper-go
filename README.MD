# E-commerce Web Scraper

This project is a web scraper for extracting product data from an e-commerce website. It uses the `Colly` library for web scraping and the `excelize` library to export the scraped data into an Excel file.

## Website
The website I am using for this is https://www.scrapingcourse.com/ecommerce/
## Introduction

The web scraper visits an e-commerce website, extracts details of products listed on the page, and saves the data into an Excel file. This tool can be useful for data analysis, price comparison, or monitoring product listings.

## Features

- Scrapes product data including name, URL, image URL, and price.
- Exports the scraped data to an Excel file.
- Handles pagination (if implemented).

## Installation

To get started with the project, follow these steps:

1. Clone the repository:
    
    git clone https://github.com/Shivam19-git/web-scraper-go.git
    cd ecommerce-web-scraper
    
2. Install the necessary dependencies:
    
    go get -u github.com/gocolly/colly
    go get -u github.com/xuri/excelize/v2
    

## Usage

1. Update the `scrapeURL` variable in the `main.go` file with the URL of the  e-commerce page you want to scrape.
2. Run the scraper:

    go run main.go

3. The scraped data will be saved into an Excel file named `products.xlsx` in the current directory.


## Dependencies

- [Colly](https://github.com/gocolly/colly): A Go library for web scraping.
- [Excelize](https://github.com/xuri/excelize): A Go library for reading and writing Excel files.

You can install these dependencies using the following commands:

go get -u github.com/gocolly/colly
go get -u github.com/xuri/excelize/v2

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to contribute to this project by opening issues or submitting pull requests. If you have any questions, please reach out to the project maintainer.

Enjoy scraping! 

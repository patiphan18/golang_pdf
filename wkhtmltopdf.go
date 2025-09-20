package main

import (
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	// create a new page
	page := wkhtmltopdf.NewPageReaderFromString(`
	<!DOCTYPE html>
	<html>
		<head>
			<title>Test PDF</title>
		</head>
		<body>
			<h1 style="color:blue;">Hello PDF from Golang!</h1>
			<p>This PDF is generated using wkhtmltopdf.</p>
		</body>
	</html>`)

	pdfg.AddPage(page)

	// create pdf
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// write file
	err = pdfg.WriteFile("output.pdf")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("PDF generated: output.pdf")
}

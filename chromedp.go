package main

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// HTML content
	content := `
	<!DOCTYPE html>
	<html>
	<head>
		<title>Test PDF</title>
	</head>
	<body>
		<h1 style="color:blue;">Hello PDF from Golang!</h1>
		<p>This PDF is generated using chromedp.</p>
	</body>
	</html>
	`

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// allocate
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var pdfBuf []byte

	// run chromedp tasks
	err := chromedp.Run(ctx,
		chromedp.Navigate("data:text/html,"+content),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuf, _, err = page.PrintToPDF().Do(ctx)
			return err
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	// save PDF file
	if err := ioutil.WriteFile("output.pdf", pdfBuf, 0644); err != nil {
		log.Fatal(err)
	}

	log.Println("PDF generated: output.pdf")
}

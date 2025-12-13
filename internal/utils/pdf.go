package utils

import (
	"context"
	"net/url"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/page"
)

// GeneratePDFFromHTML
// htmlContent: render edilmiş HTML
func GeneratePDFFromHTML(htmlContent string, outputPath string) error {

	// 1️⃣ Chrome context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// 2️⃣ Timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var pdfBuffer []byte

	// 3️⃣ HTML'i data URL'e çevir
	dataURL := "data:text/html;charset=utf-8," + url.PathEscape(htmlContent)

	// 4️⃣ Chrome ile PDF al
	err := chromedp.Run(ctx,
		chromedp.Navigate(dataURL),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			pdfBuffer, _, err = page.PrintToPDF().Do(ctx)
			return err
		}),
	)
	if err != nil {
		return err
	}

	// 5️⃣ PDF yaz
	return os.WriteFile(outputPath, pdfBuffer, 0644)
}

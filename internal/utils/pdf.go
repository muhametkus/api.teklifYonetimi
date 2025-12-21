package utils

import (
	"context"
	"net/url"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/page"
)


// GeneratePDFFromHTMLBytes
// HTML'i PDF'e çevirir ve byte slice döner
func GeneratePDFFromHTMLBytes(htmlContent string) ([]byte, error) {

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var pdfBuffer []byte

	dataURL := "data:text/html;charset=utf-8," + url.PathEscape(htmlContent)

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
		return nil, err
	}

	return pdfBuffer, nil
}


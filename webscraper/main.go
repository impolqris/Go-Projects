package main

import (
	"context"
	"log"
	"strings"

	"github.com/chromedp/chromedp"
)

const (
	ITPUrl         = "https://itpcareerexpo.com/pathways/aerospace-pathway/"
	ClassToLookFor = `a[href^="https://www.youtube.com/"]`
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var result string

	err := chromedp.Run(ctx,
		chromedp.Navigate(ITPUrl),
		chromedp.Text(ClassToLookFor, &result, chromedp.NodeVisible),
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(result))
}

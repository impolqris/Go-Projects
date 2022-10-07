package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

const (
	ITPUrl         = "https://itpcareerexpo.com/pathways/aerospace-pathway/"
	ClassToLookFor = `a[href^="https://www.youtube.com/"]`
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var result []map[string]string

	err := chromedp.Run(ctx,
		chromedp.Navigate(ITPUrl),
		chromedp.AttributesAll(ClassToLookFor, &result, chromedp.ByQueryAll),
	)

	if err != nil {
		log.Fatal(err)
	}

	for _, v := range result {
		fmt.Printf("%v\n", v)
	}
}

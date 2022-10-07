package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/chromedp/chromedp"
)

const (
	ClassToLookFor = `a[href^="https://www.youtube.com/"]`
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var result []map[string]string

	urls := []string{
		"https://itpcareerexpo.com/pathways/aerospace-pathway/",
		"https://itpcareerexpo.com/pathways/agriculture-food-natural-resources-pathway/",
		"https://itpcareerexpo.com/pathways/architecture-construction-pathway/",
		"https://itpcareerexpo.com/pathways/arts-audio-visual-technology-communications-pathway/",
		"https://itpcareerexpo.com/pathways/business-management-administration-pathway/",
		"https://itpcareerexpo.com/pathways/education-training-pathway/",
		"https://itpcareerexpo.com/pathways/energy-pathway/",
		"https://itpcareerexpo.com/pathways/engineering-pathway/",
		"https://itpcareerexpo.com/pathways/finance-pathway/",
		"https://itpcareerexpo.com/pathways/government-public-administration-pathway/",
		"https://itpcareerexpo.com/pathways/health-sciences-pathway/",
		"https://itpcareerexpo.com/pathways/hospitality-tourism-pathway/",
		"https://itpcareerexpo.com/pathways/human-services-pathway/",
		"https://itpcareerexpo.com/pathways/information-technology-pathway/",
		"https://itpcareerexpo.com/pathways/law-public-safety-corrections-security-pathway/",
		"https://itpcareerexpo.com/pathways/manufacturing-pathway/",
		"https://itpcareerexpo.com/pathways/marketing-pathway/",
		"https://itpcareerexpo.com/pathways/transportation-distribution-logistics-pathway/",
	}

	for _, v := range urls {
		err := chromedp.Run(ctx,
			chromedp.Navigate(v),
			chromedp.AttributesAll(ClassToLookFor, &result, chromedp.ByQueryAll),
		)

		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}

	fullUrls := []string{}

	for _, v := range result {
		for _, s := range v {
			if strings.Contains(s, "https://www.youtube.com") {
				fullUrls = append(fullUrls, s)
			}
		}
	}

	ytIds := []string{}

	for _, ux := range fullUrls {
		ytIds = append(ytIds, ux[32:])
	}

	fmt.Printf("%v\n", ytIds)
	fmt.Printf("%d\n", len(ytIds))
}

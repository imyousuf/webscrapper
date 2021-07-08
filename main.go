package main

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var body string
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://welcomesoftware.com/campaigns/"),
		chromedp.WaitReady("body"),
		chromedp.OuterHTML("html", &body, chromedp.ByQuery),
        chromedp.WaitVisible("h1"),
	)
	if err != nil {
		fmt.Println("Failed to open site: %v", err)
	}

	fmt.Println(body)
}

package main

import (
	"fmt"
	"net/url"
)

func main() {
	baseURL := "http://localhost:9090/api/v1/records"
	params := url.Values{}
	params.Add("description", "Watch the thrilling live cricket match between Durban Super Giants vs Paarl Royals with expert commentary and real-time score ...") // Add the query parameter with `#`
	params.Add("title", "🔥SA20 Live: Durban vs Paarl | 18th Match | Live Cricket Score &amp; Commentary")                                                          // Add the query parameter with `#`

	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println(finalURL)
}

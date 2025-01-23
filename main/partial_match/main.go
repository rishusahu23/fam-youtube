package main

import (
	"fmt"
	"net/url"
)

func main() {
	baseURL := "http://localhost:9090/api/v1/partial-match/records"
	params := url.Values{}
	params.Add("query", "Durban paarl") // Add the query parameter with `#`

	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println(finalURL)
}

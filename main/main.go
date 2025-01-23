package main

import (
	"fmt"
	"net/url"
)

func main() {
	baseURL := "http://localhost:9090/api/v1/records"
	params := url.Values{}
	params.Add("description", "Elegantly Put! 🔥❤️ #usa #israel #palestine #uk #politics #congress #news #canada #europe #australia") // Add the query parameter with `#`

	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println(finalURL)
}

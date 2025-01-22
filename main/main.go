package main

import (
	"fmt"
	"net/url"
)

func main() {
	baseURL := "http://localhost:9090/api/v1/records"
	params := url.Values{}
	params.Add("description", "The Politics 22 ม.ค. 68 I แดง-ส้ม โต้เดือด ฝุ่น 2.5 I สนทนา บก.ลายจุด - \"แดง-ส้ม\" โต้เดือด! ฝุ่น 2.5 ...") // Add the query parameter with `#`

	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println(finalURL)
}

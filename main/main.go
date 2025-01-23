package main

import (
	"fmt"
	"net/url"
)

func main() {
	baseURL := "http://localhost:9090/api/v1/records"
	params := url.Values{}
	params.Add("description", "ranjitrophy #teamindia #rohitsharma #yashasvijaiswal #shubmangill From Rohit, Jaiswal to Gill & Pant, Star Indian Players fails to ...") // Add the query parameter with `#`
	params.Add("title", "From Rohit, Jaiswal to Gill &amp; Pant, Star Players fails to impress on their Ranji Trophy comeback!")                                        // Add the query parameter with `#`

	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println(finalURL)
}

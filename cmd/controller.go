package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for {

		fmt.Println("Syncing...")
		resp, err := http.Get("http://localhost:4040")
		if err != nil {
			fmt.Println("Error:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		// Print the content of the response
		fmt.Println(string(body))
		resp.Body.Close()

		time.Sleep(5 * time.Second)
	}
}

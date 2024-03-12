package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	podName := os.Getenv("POD_NAME")
	fmt.Printf("Value of POD_NAME: %s\n", podName)
	for {

		resp, err := http.Get("http://localhost:4040")
		if err != nil {
			fmt.Println("Error:", err)
			time.Sleep(10 * time.Second)
			continue
		}

		// Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response body:", err)
			return
		}

		var electorResponse map[string]string
		err = json.Unmarshal(body, &electorResponse)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Value of electorResponse[\"name\"]: %s\n", electorResponse["name"])
		if podName == electorResponse["name"] {
			fmt.Println("Syncing...")
			fmt.Printf("Name of leader: %s\n", electorResponse["name"])
		}
		// Print the content of the response
		resp.Body.Close()
		time.Sleep(10 * time.Second)
	}
}

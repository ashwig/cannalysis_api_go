package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func makeAPIKey(vendorKey string, userKey string) string {
	return vendorKey + ":" + userKey
}

func makeMETRCAuthURI(METRCURI string, vendorKey string, userKey string) string {
	return METRCURI + makeAPIKey(vendorKey, userKey)
}

func main() {
	// Import environment variables for configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment configuration: %v", err)
	}

	vendorKey := os.Getenv("METRC_VENDOR_KEY")
	userKey := os.Getenv("METRC_USER_KEY")
	metrcAuthURI := makeMETRCAuthURI(os.Getenv("METRC_URI"), vendorKey, userKey)

	client := &http.Client{}
	if err != nil {
		log.Fatalf("Error creating HTTP client: %v", err)
	}

	resp, err := client.Get(os.Getenv("METRC_URI"))
	if err != nil {
		log.Fatalf("Error reading HTTP response: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading HTTP response: %v", err)
	}
	log.Printf("Auth: %s", metrcAuthURI)
	log.Printf("Response:\n-----\n%v", body)

	req, err := http.NewRequest("GET", "/plants/v1/active/0", nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	} else {
		req.Header.Add("Authorization", makeAPIKey(vendorKey, userKey))
		req.Header.Add("Content-Type", "application/json")
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("Error reading HTTP response: %v", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading HTTP response: %v", err)
		}
		log.Printf("Response:\n-----\n%v", body)

		// Output our response to the terminal (stdout)
		log.Printf(`The response is below:
--------------------------------------------------------------------------------
%v`, string(body))
	}
}

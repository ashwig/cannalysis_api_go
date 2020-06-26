package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	auth "github.com/abbot/go-http-auth"
	"github.com/joho/godotenv"
)

func getVendorKey() string {
	// Import environment variables for configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment configuration: %v", err)
	}

	return os.Getenv("METRC_VENDOR_KEY")
}

func getUserKey() string {
	// Import environment variables for configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment configuration: %v", err)
	}

	return os.Getenv("METRC_USER_KEY")
}

func getAPIKey() string {
	// Import environment variables for configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment configuration: %v", err)
	}

	return os.Getenv("METRC_API_KEY")
}

func generateURI() string {
	// Import environment variables for configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment configuration: %v", err)
	}

	metrcURIBase := os.Getenv("METRC_URI")

	return string(getVendorKey() + ":" + getUserKey() + "@" + metrcURIBase)
}

func secret() string {
	return string(base64.RawStdEncoding.EncodeToString([]byte(getAPIKey())))
}

func handle(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	return
}

func main() {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, string(generateURI()+"/packages/v1/0"), nil)
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	} else {
		req.Header.Add("Authorization", string("Basic "+os.Getenv("METRC_ENCODED_API")))
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
		log.Printf("Response(%v):\n-----\n%v", string(resp.StatusCode), body)

		// Output our response to the terminal (stdout)
		log.Printf(`The response is below:
--------------------------------------------------------------------------------
%v`, string(body))
	}
}

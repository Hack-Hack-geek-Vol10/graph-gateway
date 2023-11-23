package google

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// TODO:SDKからJWKsを取得する
var GoogleJWks map[string]interface{}

func GetGoogleJWKs() {
	resp, err := http.Get("https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com")
	if err != nil {
		log.Fatalf("Failed to make a request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)

	if err != nil {
		log.Fatalf("Failed to json unmarshal: %v", err)
	}

	GoogleJWks = result
}

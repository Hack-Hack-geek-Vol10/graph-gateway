package google

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
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

	parseToInterface(body)
}

func ParseGoogleJWKs(path string) {
	data, err := os.Open(path)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	body, err := io.ReadAll(data)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}

	parseToInterface(body)
}

func parseToInterface(data []byte) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result)

	if err != nil {
		log.Fatalf("Failed to json unmarshal: %v", err)
	}

	GoogleJWks = result
}

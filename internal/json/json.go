package json

import (
	"encoding/json"
	"net/http"
	"time"
)

// client: Define timeout of http.Client
var	client = &http.Client{Timeout: time.Second * 15}

// GetJSON(): Parse JSON to targeted struct
func GetJSON(url string, target any) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(&target)
}

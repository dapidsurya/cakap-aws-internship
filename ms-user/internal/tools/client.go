package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func get(url string, result interface{}) error {
	fmt.Println("Calling url:", url)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Error on http request:", err)
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)

	if err != nil {
		fmt.Println("Error on converting response:", err)
		return err
	}

	return nil
}

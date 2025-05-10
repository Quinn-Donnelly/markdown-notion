package notion

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	queryEndpointTemplate = "https://api.notion.com/v1/databases/%s/query"
)

type Client struct {
	ApiToken string
}


type DatabaseResults struct {
	Object  string   `json:"object"`
	Results []NotionObject `json:"results"`
}

type NotionObject struct {
	Object string `json:"object"`
	Id string `json:"string"`
}

func (c *Client) GetTasks(databaseID string) {
	filter, _ := json.Marshal(map[string]string{})
	filterBody := bytes.NewBuffer(filter)

	log.Println(fmt.Sprintf(queryEndpointTemplate, taskDatabaseId))
	req, err := http.NewRequest("POST", fmt.Sprintf(queryEndpointTemplate, taskDatabaseId), filterBody)
	if err != nil {
		log.Fatalf("unable to create request: %f", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiToken))
	req.Header.Set("Notion-Version", "2022-02-22")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error retrieving database information: %s", err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal("API response non 200: failed to call API")
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read response body: %s", err)
	}

	searchResults := &DatabaseResults{}
	err = json.Unmarshal(bodyBytes, &searchResults)
	if err != nil {
		log.Fatalf("Can't parse response body: %s", err)
	}

	log.Printf("Information retrieved: %v", searchResults)
}

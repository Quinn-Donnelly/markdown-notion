package notion

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/ohler55/ojg/jp"
	"github.com/ohler55/ojg/oj"
)

const (
	queryEndpointTemplate = "https://api.notion.com/v1/databases/%s/query"
)

type Client struct {
	ApiToken string
}

type DatabaseResults struct {
	Object  string         `json:"object"`
	Results []NotionObject `json:"results"`
}

type NotionObject struct {
	Object string `json:"object"`
	Id     string `json:"string"`
}

func (o *NotionObject) String() string {
	return fmt.Sprintf("%s: %v", o.Id, o.Object)
}

func (d *DatabaseResults) String() string {
	return fmt.Sprintf("object: %s, results: %#v", d.Object, d.Results)
}

type Task struct {
	Name string
}

func convertToTask(json string) []Task {
	data, _ := oj.ParseString(json)
	expression, _ := jp.ParseString(".results[*].properties.Name.title")

	thing := expression.Get(data)

	taskList := []Task{}
	for _, page := range thing {
		task := Task{}
		someShit := page.([]any)
		for _, textObjs := range someShit {
			task.Name = fmt.Sprintf("%s%s", task.Name, (textObjs.(map[string]any)["plain_text"]))
		}
		taskList = append(taskList, task)
	}

	return taskList
}

func (c *Client) GetTasks(databaseID string) {
	reader := strings.NewReader(`
		{
		  "filter": {
			  "property": "Flags for Actions",
				  "multi_select": {
					  "contains": "Focus Today"
				  }
		  }
		}
	`)

	log.Println(fmt.Sprintf(queryEndpointTemplate, databaseID))
	req, err := http.NewRequest("POST", fmt.Sprintf(queryEndpointTemplate, databaseID), reader)
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

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Can't read response body: %s", err)
	}

	searchResults := &DatabaseResults{}
	fmt.Println(convertToTask(string(bodyBytes)))
	err = json.Unmarshal(bodyBytes, &searchResults)
	if err != nil {
		log.Fatalf("Can't parse response body: %s", err)
	}
}

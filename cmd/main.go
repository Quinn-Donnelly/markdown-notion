package main

import (
	"context"
	"log"
	"os"

	"github.com/Quinn-Donnelly/markdown-notion/src/notion"
)

const (
	NOTION_API = "NOTION_API_SECRET"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("hey dog wanna tell me what database to query")
	}

	taskDatabaseId := os.Args[1]
	apiToken, found := os.LookupEnv(NOTION_API)
	if !found {
		log.Fatalf("No %s env var defined, you need a notion api token defined", NOTION_API)
	}

	client := notion.NewClient(apiToken, taskDatabaseId)

	tasks, err := client.GetFocusToday(context.Background())
	if err != nil {
		log.Fatalf("Error getting Focus today: %v", err)
	}

	for _, task := range tasks {
		log.Printf("Task: %+v\n", task)
	}
}

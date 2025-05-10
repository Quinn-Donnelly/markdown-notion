package main

import (
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

	client := notion.Client{
		ApiToken: apiToken,
	}

	client.GetTasks(taskDatabaseId)
}

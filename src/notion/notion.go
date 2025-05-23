package notion

import (
	"context"
	"log/slog"

	"github.com/jomei/notionapi"
)

type Client struct {
	Client         *notionapi.Client
	taskDatabaseId notionapi.DatabaseID
}

func NewClient(apiToken string, taskDatabaseid string) *Client {
	client := notionapi.NewClient(notionapi.Token(apiToken))

	return &Client{
		Client:         client,
		taskDatabaseId: notionapi.DatabaseID(taskDatabaseid),
	}
}

func (c *Client) GetFocusToday(ctx context.Context) ([]*Task, error) {
	resp, err := c.Client.Database.Query(ctx, c.taskDatabaseId, &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: "Flags for Actions",
			MultiSelect: &notionapi.MultiSelectFilterCondition{
				Contains: "Focus Today",
			},
		},
	})

	slog.LogAttrs(ctx, slog.LevelDebug,"results from query of database",
		slog.Group("API response",
			slog.Any("resp", resp),
			slog.Any("error", err),
		),
	)

	tasks := []*Task{}
	for _, thing := range resp.Results {
		title := thing.Properties["Name"]
		if title != nil && title.GetType() == "title" {
			name := title.(*notionapi.TitleProperty).Title[0].PlainText
			tasks = append(tasks, &Task{
				Name: name,
			})
		}
	}

	return tasks, nil
}

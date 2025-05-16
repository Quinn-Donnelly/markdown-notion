POST https://api.notion.com/v1/databases/{{task_database_id}}/query
Authorization: Bearer {{notion_api_secret}}
Notion-Version: 2022-02-22
Content-Type: application/json
{
  "filter": {
      "property": "Flags for Actions",
          "multi_select": {
              "contains": "Weekly Action"
          }
  }
}


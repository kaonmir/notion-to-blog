package blog

import (
	"context"
	"log"
	"regexp"

	"github.com/jomei/notionapi"
)

type Notion struct {
	client *notionapi.Client
}

func NewNotion(secretKey string) *Notion {
	return &Notion{
		client: notionapi.NewClient(notionapi.Token(secretKey)),
	}
}

func (n *Notion) GetPostablePages(databaseURL string) []notionapi.Page {
	databaseID := regexp.MustCompile(`https://www.notion.so/.*?/([a-z0-9]+)\??.*`).FindStringSubmatch(databaseURL)[1]

	query := notionapi.DatabaseQueryRequest{
		Filter: notionapi.OrCompoundFilter{
			notionapi.PropertyFilter{
				Property: "상태",
				Select: &notionapi.SelectFilterCondition{
					Equals: "발행 요청",
				},
			},
			notionapi.PropertyFilter{
				Property: "상태",
				Select: &notionapi.SelectFilterCondition{
					Equals: "수정 요청",
				},
			},
		},
	}
	resp, err := n.client.Database.Query(context.Background(), notionapi.DatabaseID(databaseID), &query)
	if err != nil {
		log.Fatal(err)
	}

	return resp.Results
}

func (n *Notion) ExportPage(page notionapi.Page) {
}

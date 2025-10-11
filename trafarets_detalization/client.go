package trafarets_detalization

import (
    "time"

    "github.com/execut/ozon-reports-downloader/common"
)

type Client struct {
    client *common.Client
}

func NewClient(client *common.Client) *Client {
    return &Client{
        client: client,
    }
}

func (c *Client) Download() ([]byte, error) {
    now := time.Now().In(time.UTC)
    atTo := now.Truncate(time.Hour * 24).Add(-time.Second)
    atFrom := now.Truncate(time.Hour*24).AddDate(-3, 0, 0)

    baseURL := "https://seller.ozon.ru/performance-api/seller-api/adv-api/v2/campaign_expense"
    payload := &RequestPayload{
        DateFrom:      atFrom.Format(time.DateOnly),
        DateTo:        atTo.Format(time.DateOnly),
        CampaignTypes: []string{},
    }
    bodyBytes, err := c.client.DoPostPerformanceRequest(payload, baseURL)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

type RequestPayload struct {
    DateFrom       string   `json:"dateFrom"`
    DateTo         string   `json:"dateTo"`
    CampaignSearch string   `json:"campaignSearch"`
    CampaignTypes  []string `json:"campaignTypes"`
}

package trafarets_detalization

import (
    "github.com/execut/ozon-reports-downloader/common"
    "time"
)

type Client struct {
}

func NewClient() *Client {
    return &Client{}
}

func (c *Client) Download(companyID int64, organizationID int64, cookie string) ([]byte, error) {
    now := time.Now().In(time.UTC)
    atTo := now.Truncate(time.Hour * 24).Add(-time.Second)
    atFrom := now.Truncate(time.Hour*24).AddDate(-3, 0, 0)

    baseURL := "https://seller.ozon.ru/performance-api/seller-api/adv-api/v2/campaign_expense"
    payload := &RequestPayload{
        DateFrom:      atFrom.Format(time.DateOnly),
        DateTo:        atTo.Format(time.DateOnly),
        CampaignTypes: []string{},
    }
    bodyBytes, err := common.DoPostPerformanceRequest(payload, baseURL, companyID, organizationID, cookie)
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

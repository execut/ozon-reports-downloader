package warehousing_cost_month

import (
    "strconv"
    "time"

    "github.com/execut/ozon-reports-downloader/common"
)

type Client struct {
    client    common.Client
    companyID int64
}

func NewClient(client common.Client, companyID int64) *Client {
    return &Client{
        client:    client,
        companyID: companyID,
    }
}

func (c *Client) Download() ([]byte, error) {
    now := time.Now().In(time.UTC)
    baseURL := "https://seller.ozon.ru/api/site/self-placement-gateway/api/reports/" + strconv.FormatInt(c.companyID, 10) + "/placement/periods/date/items-report/download"
    payload := &RequestPayload{
        Date: now.Format(time.DateOnly),
    }
    bodyBytes, err := c.client.DoRequest(payload, baseURL)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

type RequestPayload struct {
    Date string `json:"date"`
}

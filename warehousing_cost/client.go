package warehousing_cost

import (
    "strconv"
    "time"

    "github.com/execut/ozon-reports-downloader/common"
)

type Client struct {
    companyID int64
    client    *common.Client
}

func NewClient(companyID int64, client *common.Client) *Client {
    return &Client{companyID: companyID, client: client}
}

func (c *Client) Download() ([]byte, error) {
    now := time.Now().In(time.UTC)
    yesterday := now.AddDate(0, 0, -1).Format("2006-01-02")
    payload := RequestPayload{
        From: yesterday,
        To:   yesterday,
    }
    baseURL := "https://seller.ozon.ru/api/site/self-placement-gateway/api/reports/" + strconv.FormatInt(c.companyID, 10) + "/placement/periods/items-report/download?from=" + payload.From + "&to=" + payload.To
    bodyBytes, err := c.client.DoGetRequest(nil, baseURL)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

type RequestPayload struct {
    From string `json:"from"`
    To   string `json:"to"`
}

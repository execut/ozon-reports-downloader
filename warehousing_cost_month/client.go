package warehousing_cost_month

import (
    "github.com/execut/ozon-reports-downloader/common"
    "strconv"
    "time"
)

type Client struct {
}

func NewClient() *Client {
    return &Client{}
}

func (c *Client) Download(companyID int64, cookie string) ([]byte, error) {
    now := time.Now().In(time.UTC)
    baseURL := "https://seller.ozon.ru/api/site/self-placement-gateway/api/reports/" + strconv.FormatInt(companyID, 10) + "/placement/periods/date/items-report/download"
    payload := &RequestPayload{
        Date: now.Format(time.DateOnly),
    }
    bodyBytes, err := common.DoRequest(payload, baseURL, cookie, companyID)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

type RequestPayload struct {
    Date string `json:"date"`
}

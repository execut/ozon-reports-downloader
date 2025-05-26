package warehousing_cost

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
    yesterday := now.AddDate(0, 0, -1).Format("2006-01-02")
    payload := RequestPayload{
        From: yesterday,
        To:   yesterday,
    }
    baseURL := "https://seller.ozon.ru/api/site/self-placement-gateway/api/reports/" + strconv.FormatInt(companyID, 10) + "/placement/periods/items-report/download?from=" + payload.From + "&to=" + payload.To
    bodyBytes, err := common.DoGetRequest(nil, baseURL, cookie, companyID)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

type RequestPayload struct {
    From string `json:"from"`
    To   string `json:"to"`
}

package accruals

import (
    "fmt"
    "net/url"
    "time"

    "github.com/execut/ozon-reports-downloader/common"
)

type Client struct {
    commonClient *common.Client
    companyID    int64
}

func NewClient(commonClient *common.Client, companyID int64) *Client {
    return &Client{
        commonClient: commonClient,
        companyID:    companyID,
    }
}

func (c *Client) Download() ([]byte, error) {
    now := time.Now().In(time.UTC)
    atFrom := time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC)
    atTo := now.Truncate(time.Hour * 24).Add(-time.Second)

    baseURL := "https://seller.ozon.ru"
    resource := fmt.Sprintf("/api/site/self-gateway/api/reports/%v/accruals/product/download", c.companyID)
    params := url.Values{}
    params.Add("from", atFrom.Format(time.DateOnly))
    params.Add("to", atTo.Format(time.DateOnly))

    u, _ := url.ParseRequestURI(baseURL)
    u.Path = resource
    u.RawQuery = params.Encode()
    urlValue := fmt.Sprintf("%v", u)
    bodyBytes, err := c.commonClient.DoGetRequest(struct{}{}, urlValue)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

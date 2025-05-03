package accruals

import (
    "fmt"
    "net/url"
    "ozon_reports_downloader/common"
    "time"
)

type Client struct {
}

func NewClient() *Client {
    return &Client{}
}

func (c *Client) Download(cookie string, companyID int64) ([]byte, error) {
    now := time.Now().In(time.UTC)
    atFrom := time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC)
    atTo := now.Truncate(time.Hour * 24).Add(-time.Second)

    baseURL := "https://seller.ozon.ru"
    resource := fmt.Sprintf("/api/site/self-gateway/api/reports/%v/accruals/product/download", companyID)
    params := url.Values{}
    params.Add("from", atFrom.Format(time.DateOnly))
    params.Add("to", atTo.Format(time.DateOnly))

    u, _ := url.ParseRequestURI(baseURL)
    u.Path = resource
    u.RawQuery = params.Encode()
    urlValue := fmt.Sprintf("%v", u)
    bodyBytes, err := common.DoGetRequest(struct{}{}, urlValue, cookie, companyID)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

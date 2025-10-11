package search_promotion_orders

import (
    "encoding/json"
    "time"

    "github.com/execut/ozon-reports-downloader/common"
    "github.com/google/uuid"
)

const urlsPrefix = "https://performance.ozon.ru/seller-api/search-performance-cpo/mainpage/v1/statistic"

type Client struct {
    client *common.Client
}

func NewClient(client *common.Client) *Client {
    return &Client{client: client}
}

func (c *Client) BeginDownload() (*uuid.UUID, error) {
    now := time.Now().Round(time.Hour * 24)
    atTo := now.Truncate(time.Hour*24).AddDate(0, 0, -1)
    atFrom := now.Truncate(time.Hour*24).AddDate(0, -2, 0)
    data := StartDownloadPayload{
        TimeBounds: StartDownloadTimeBounds{
            From: atFrom,
            To:   atTo,
        },
    }
    url := urlsPrefix + "/orders/generate"

    bodyBytes, err := c.client.DoPostPerformanceRequest(data, url)
    if err != nil {
        return nil, err
    }

    response := &StartResponse{}
    err = json.Unmarshal(bodyBytes, response)
    if err != nil {
        return nil, err
    }

    uuidValue, err := uuid.Parse(response.Code)
    if err != nil {
        return nil, err
    }

    return &uuidValue, nil
}

func (c *Client) ReportsList() (*ReportResponse, error) {
    data, err := c.client.DoGetPerformanceRequest(urlsPrefix + "/reports")
    if err != nil {
        return nil, err
    }

    response := &ReportResponse{}
    err = json.Unmarshal(data, response)
    if err != nil {
        return nil, err
    }

    return response, err
}

func (c *Client) Download(uuid string) ([]byte, error) {
    data, err := c.client.DoGetPerformanceRequest(urlsPrefix + "/report?UUID=" + uuid)
    if err != nil {
        return nil, err
    }

    return data, err
}

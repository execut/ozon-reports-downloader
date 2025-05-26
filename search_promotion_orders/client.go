package search_promotion_orders

import (
    "encoding/json"
    "github.com/google/uuid"
    "github.com/execut/ozon-reports-downloader/common"
    "time"
)

const urlsPrefix = "https://performance.ozon.ru/seller-api/search-performance-cpo/mainpage/v1/statistic"

type Client struct {
}

func NewClient() *Client {
    return &Client{}
}

func (c *Client) BeginDownload(companyID int64, organizationID int64, cookie string) (*uuid.UUID, error) {
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

    bodyBytes, err := common.DoPostPerformanceRequest(data, url, companyID, organizationID, cookie)
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

func (c *Client) ReportsList(companyID int64, organizationID int64, cookie string) (*ReportResponse, error) {
    data, err := common.DoGetPerformanceRequest(urlsPrefix+"/reports", companyID, organizationID, cookie)
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

func (c *Client) Download(uuid string, companyID int64, organizationID int64, cookie string) ([]byte, error) {
    data, err := common.DoGetPerformanceRequest(urlsPrefix+"/report?UUID="+uuid, companyID, organizationID, cookie)
    if err != nil {
        return nil, err
    }

    return data, err
}

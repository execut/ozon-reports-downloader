package orders

import (
    "encoding/json"
    "strconv"
    "time"

    "github.com/execut/ozon-reports-downloader/common"
    "github.com/google/uuid"
)

type Client struct {
}

func NewClient() *Client {
    return &Client{}
}

func (c *Client) BeginDownload(orderType common.DeliveryType, companyID int64, cookie string, currentTime time.Time) (*uuid.UUID, error) {
    atTo := currentTime.Truncate(time.Hour * 24).Add(-time.Second)
    atFrom := currentTime.Truncate(time.Hour*24).AddDate(0, -3, 0)
    data := StartPayload{
        Filter: Filter{
            ProcessedAtTo:   atTo,
            ProcessedAtFrom: atFrom,
            DeliverySchema:  string(orderType),
        },
        With: With{
            AnalyticsData: true,
            JewelryCodes:  true,
        },
        Lang:      "RU",
        CompanyID: strconv.FormatInt(companyID, 10),
        SortDir:   "desc",
    }
    url := "https://seller.ozon.ru/api/report/company/postings"

    bodyBytes, err := common.DoRequest(data, url, cookie, companyID)
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

func (c *Client) Status(code *uuid.UUID, cookie string, companyID int64) (*StatusResponse, error) {
    data, err := common.DoRequest(StatusPayload{Code: code.String()}, "https://seller.ozon.ru/api/report/status", cookie, companyID)
    if err != nil {
        return nil, err
    }

    response := &StatusResponse{}
    err = json.Unmarshal(data, response)
    if err != nil {
        return nil, err
    }

    return response, err
}

func (c *Client) Download(code *uuid.UUID, companyID int64, cookie string) (*DownloadResult, error) {
    data, err := common.DoRequest(DownloadPayload{Code: code.String(), CompanyID: companyID}, "https://seller.ozon.ru/api/report/download", cookie, companyID)
    if err != nil {
        return nil, err
    }

    response := &DownloadResult{}
    err = json.Unmarshal(data, response)
    if err != nil {
        return nil, err
    }

    return response, err
}

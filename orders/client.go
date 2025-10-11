package orders

import (
    "encoding/json"
    "strconv"
    "time"

    "github.com/execut/ozon-reports-downloader/common"
    "github.com/google/uuid"
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

func (c *Client) BeginDownload(orderType common.DeliveryType, currentTime time.Time) (*uuid.UUID, error) {
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
        CompanyID: strconv.FormatInt(c.companyID, 10),
        SortDir:   "desc",
    }
    url := "https://seller.ozon.ru/api/report/company/postings"

    bodyBytes, err := c.commonClient.DoRequest(data, url)
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

func (c *Client) Status(code *uuid.UUID) (*StatusResponse, error) {
    data, err := c.commonClient.DoRequest(StatusPayload{Code: code.String()}, "https://seller.ozon.ru/api/report/status")
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

func (c *Client) Download(code *uuid.UUID) (*DownloadResult, error) {
    data, err := c.commonClient.DoRequest(DownloadPayload{Code: code.String(), CompanyID: c.companyID}, "https://seller.ozon.ru/api/report/download")
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

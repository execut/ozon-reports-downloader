package prices

import (
    "encoding/json"
    "strconv"

    "github.com/execut/ozon-reports-downloader/common"
    "github.com/google/uuid"
)

const urlsPrefix = "https://seller.ozon.ru/api/pricing-report-service/v1/report"

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

func (c *Client) BeginDownload() (*uuid.UUID, error) {
    data := StartPayload{
        IsSuperEconomEnabled: true,
        PriceColorIndex:      []string{"1", "2", "3", "0"},
        Visibility:           "ALL",
    }
    url := urlsPrefix + "/new"

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
    payload := StatusPayload{
        Code: code.String(),
    }
    data, err := c.commonClient.DoRequest(payload, urlsPrefix+"/status")
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

func (c *Client) Download(code *uuid.UUID) ([]byte, error) {
    data, err := c.commonClient.DoGetRequest(struct{}{}, "https://seller.ozon.ru/api/pricing-report-service/report/download/"+strconv.FormatInt(c.companyID, 10)+"/"+code.String())
    if err != nil {
        return nil, err
    }

    return data, nil
}

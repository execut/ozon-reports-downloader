package analytics

import (
    "encoding/json"
    "time"

    "github.com/google/uuid"

    "github.com/execut/ozon-reports-downloader/common"
)

type Client struct {
    commonClient *common.Client
}

func NewClient(commonClient *common.Client) *Client {
    return &Client{commonClient: commonClient}
}

func (c *Client) BeginDownload(prevDate time.Time) (*uuid.UUID, error) {
    now := time.Now().In(time.UTC)
    atFrom := prevDate.Truncate(time.Hour*24).AddDate(0, 0, 1)
    atTo := now.Truncate(time.Hour * 24).Add(-time.Second)
    data := StartPayload{
        Metrics:    []string{"hits_view", "session_view_pdp", "ordered_units", "hits_tocart", "conv_tocart_pdp", "hits_tocart_pdp", "revenue"},
        Dimensions: []string{"sku", "day", "modelID"},
        DateFrom:   atFrom.Format("2006-01-02"),
        DateTo:     atTo.Format("2006-01-02"),
    }
    url := "https://seller.ozon.ru/api/v1/report/data-v1-xlsx"

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
    data, err := c.commonClient.DoGetRequest(struct{}{}, "https://seller.ozon.ru/api/v1/report/status/"+code.String())
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
    data, err := c.commonClient.DoGetRequest(struct{}{}, "https://seller.ozon.ru/api/v1/report/download/"+code.String())
    if err != nil {
        return nil, err
    }

    return data, nil
}

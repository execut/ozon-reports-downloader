package analytics

import (
    "encoding/json"
    "github.com/google/uuid"
    "ozon_reports_downloader/common"
    "time"
)

type Client struct {
}

func NewClient() *Client {
    return &Client{}
}

func (c *Client) BeginDownload(prevDate time.Time, cookie string, companyID int64) (*uuid.UUID, error) {
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
    data, err := common.DoGetRequest(struct{}{}, "https://seller.ozon.ru/api/v1/report/status/"+code.String(), cookie, companyID)
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

func (c *Client) Download(code *uuid.UUID, cookie string, companyID int64) ([]byte, error) {
    data, err := common.DoGetRequest(struct{}{}, "https://seller.ozon.ru/api/v1/report/download/"+code.String(), cookie, companyID)
    if err != nil {
        return nil, err
    }

    return data, nil
}

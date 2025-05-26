package prices

import (
    "encoding/json"
    "github.com/execut/ozon-reports-downloader/common"
    "github.com/google/uuid"
    "strconv"
)

const urlsPrefix = "https://seller.ozon.ru/api/pricing-report-service/v1/report"

type Client struct {
}

func NewClient() *Client {
    return &Client{}
}

func (c *Client) BeginDownload(companyID int64, cookie string) (*uuid.UUID, error) {
    data := StartPayload{
        IsSuperEconomEnabled: true,
        PriceColorIndex:      []string{"1", "2", "3", "0"},
        Visibility:           "ALL",
    }
    url := urlsPrefix + "/new"

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
    payload := StatusPayload{
        Code: code.String(),
    }
    data, err := common.DoRequest(payload, urlsPrefix+"/status", cookie, companyID)
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
    data, err := common.DoGetRequest(struct{}{}, "https://seller.ozon.ru/api/pricing-report-service/report/download/"+strconv.FormatInt(companyID, 10)+"/"+code.String(), cookie, companyID)
    if err != nil {
        return nil, err
    }

    return data, nil
}

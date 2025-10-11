package returns

import (
    "encoding/json"
    "errors"
    "io"
    "net/http"
    "time"

    "github.com/execut/ozon-reports-downloader/common"

    "github.com/google/uuid"
)

type Client struct {
    client    *common.Client
    companyID int64
}

func NewClient(client *common.Client, companyID int64) *Client {
    return &Client{
        client:    client,
        companyID: companyID,
    }
}

func (c *Client) BeginDownload(returnsType ReturnsType, now time.Time) (*uuid.UUID, error) {
    atTo := now.Truncate(time.Hour * 24).Add(-time.Second)
    var data interface{}
    var atFrom time.Time
    switch returnsType {
    case ReturnsTypeFBOS:
        atFrom = now.Truncate(time.Hour*24).AddDate(0, -3, 0)
        data = StartPayloadFBOS{
            SellerId:       c.companyID,
            TimeZoneOffset: 3,
            Parameters: PayloadParameters{
                DateFrom:      atFrom,
                DateTo:        atTo,
                PreFilterType: 90,
            },
        }
    case ReturnsTypeRealFBS:
        atFrom = now.Truncate(time.Hour*24).AddDate(2022, 4, 1)
        data = StartPayloadRealFBS{
            SellerId:       c.companyID,
            TimeZoneOffset: 3,
            Parameters: PayloadParameters{
                DateFrom:      atFrom,
                DateTo:        atTo,
                PreFilterType: 90,
            },
        }
    default:
        return nil, errors.New("unsupported delivery type")
    }

    url := "https://seller.ozon.ru/api/site/seller-returns-report-service/generate"

    bodyBytes, err := c.client.DoRequest(data, url)
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
    data, err := c.client.DoRequest(StatusPayload{
        Code:           code.String(),
        SellerId:       c.companyID,
        TimeZoneOffset: 3,
    }, "https://seller.ozon.ru/api/site/seller-returns-report-service/status")
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

func (c *Client) Download(fileUrl string) (io.Reader, error) {
    data, err := http.Get(fileUrl)
    if err != nil {
        return nil, err
    }

    return data.Body, err
}

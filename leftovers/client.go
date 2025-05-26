package leftovers

import (
    "github.com/execut/ozon-reports-downloader/common"
)

type Client struct {
}

func NewClient() *Client {
    return &Client{}
}

func (c *Client) Download(cookie string, companyID int64) ([]byte, error) {
    bodyBytes, err := common.DoRequest(Payload{WarehouseType: "All"}, "https://seller.ozon.ru/api/som-stocks-bff/Report/GetStockApiReport", cookie, companyID)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

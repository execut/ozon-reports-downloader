package leftovers

import (
    "github.com/execut/ozon-reports-downloader/common"
)

type Client struct {
    commonClient *common.Client
}

func NewClient(commonClient *common.Client) *Client {
    return &Client{
        commonClient: commonClient,
    }
}

func (c *Client) Download() ([]byte, error) {
    bodyBytes, err := c.commonClient.DoRequest(Payload{WarehouseType: "All"}, "https://seller.ozon.ru/api/som-stocks-bff/Report/GetStockApiReport")
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

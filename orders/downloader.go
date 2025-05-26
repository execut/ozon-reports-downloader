package orders

import (
    "encoding/base64"
    "fmt"
    "time"

    "github.com/execut/ozon-reports-downloader/common"
    "github.com/execut/ozon-reports-downloader/file"
)

type Downloader struct {
    client       *Client
    deliveryType common.DeliveryType
    companyID    int64
    cookie       string
}

func NewDownloader(deliveryType common.DeliveryType, companyID int64, cookie string) *Downloader {
    return &Downloader{
        client:       &Client{},
        deliveryType: deliveryType,
        companyID:    companyID,
        cookie:       cookie,
    }
}

func (d Downloader) Download() (*file.File, error) {
    uuid, err := d.client.BeginDownload(d.deliveryType, d.companyID, d.cookie)
    if err != nil {
        return nil, err
    }

    fmt.Println("Downloading order", uuid)

    for {
        time.Sleep(time.Second)
        status, errStatus := d.client.Status(uuid, d.cookie, d.companyID)
        if errStatus != nil {
            return nil, errStatus
        }

        fmt.Println("Status:", status)
        if status.Status == "success" {
            break
        }
    }

    fileData, err := d.client.Download(uuid, d.companyID, d.cookie)
    if err != nil {
        return nil, err
    }

    data, err := base64.StdEncoding.DecodeString(fileData.Result.FileContent)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "csv"), nil
}

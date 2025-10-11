package orders

import (
    "bytes"
    "encoding/base64"
    "fmt"
    "io"
    "time"

    "github.com/execut/ozon-reports-downloader/common"
    "github.com/execut/ozon-reports-downloader/file"
)

type Downloader struct {
    time         time.Time
    client       *Client
    deliveryType common.DeliveryType
}

func NewDownloader(deliveryType common.DeliveryType, time time.Time, client *Client) *Downloader {
    return &Downloader{
        time:         time,
        client:       client,
        deliveryType: deliveryType,
    }
}

func (d Downloader) Download() (*file.File, error) {
    currentTime := d.time
    isFirstFile := true
    var data []byte
    for {
        uuid, err := d.client.BeginDownload(d.deliveryType, currentTime)
        if err != nil {
            return nil, err
        }

        fmt.Println("Downloading order", uuid)
        for {
            time.Sleep(time.Second)
            status, errStatus := d.client.Status(uuid)
            if errStatus != nil {
                return nil, errStatus
            }

            fmt.Println("Status:", status)
            if status.Status == "success" {
                break
            }
        }

        fileData, err := d.client.Download(uuid)
        if err != nil {
            return nil, err
        }

        currentData, err := base64.StdEncoding.DecodeString(fileData.Result.FileContent)
        if err != nil {
            return nil, err
        }

        buffer := bytes.NewBuffer(currentData)
        firstLine, err := buffer.ReadBytes('\n')
        if err != nil && err != io.EOF {
            return nil, err
        }

        if isFirstFile {
            data = append(data, firstLine...)
            isFirstFile = false
        }

        otherLines := buffer.Bytes()
        if len(otherLines) == 0 {
            break
        }

        data = append(data, otherLines...)

        currentTime = currentTime.AddDate(0, -3, 0)
    }

    return file.NewFile(data, "csv"), nil
}

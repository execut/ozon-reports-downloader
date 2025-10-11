package analytics

import (
    "fmt"
    "time"

    "github.com/execut/ozon-reports-downloader/file"
)

type Downloader struct {
    client   *Client
    prevDate time.Time
}

func NewDownloader(prevDate time.Time, client *Client) *Downloader {
    return &Downloader{
        client:   client,
        prevDate: prevDate,
    }
}

func (d Downloader) Download() (*file.File, error) {
    uuid, err := d.client.BeginDownload(d.prevDate)
    if err != nil {
        return nil, err
    }

    fmt.Println("Downloading analytics", uuid)

    for {
        time.Sleep(time.Second * 5)
        status, errStatus := d.client.Status(uuid)
        if errStatus != nil {
            return nil, errStatus
        }

        fmt.Println("Status:", status)
        if status.Status == "success" {
            break
        }
    }

    data, err := d.client.Download(uuid)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

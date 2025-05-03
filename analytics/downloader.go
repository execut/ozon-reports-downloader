package analytics

import (
    "fmt"
    "ozon_reports_downloader/file"
    "time"
)

type Downloader struct {
    client    *Client
    prevDate  time.Time
    cookie    string
    companyID int64
}

func NewDownloader(prevDate time.Time, cookie string, companyID int64) *Downloader {
    return &Downloader{
        client:    &Client{},
        prevDate:  prevDate,
        cookie:    cookie,
        companyID: companyID,
    }
}

func (d Downloader) Download() (*file.File, error) {
    uuid, err := d.client.BeginDownload(d.prevDate, d.cookie, d.companyID)
    if err != nil {
        return nil, err
    }

    fmt.Println("Downloading analytics", uuid)

    for {
        time.Sleep(time.Second * 5)
        status, errStatus := d.client.Status(uuid, d.cookie, d.companyID)
        if errStatus != nil {
            return nil, errStatus
        }

        fmt.Println("Status:", status)
        if status.Status == "success" {
            break
        }
    }

    data, err := d.client.Download(uuid, d.cookie, d.companyID)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

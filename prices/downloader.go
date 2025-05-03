package prices

import (
    "fmt"
    "ozon_reports_downloader/file"
    "time"
)

type Downloader struct {
    client    *Client
    companyID int64
    cookie    string
}

func NewDownloader(companyID int64, cookie string) *Downloader {
    return &Downloader{
        client:    &Client{},
        companyID: companyID,
        cookie:    cookie,
    }
}

func (d Downloader) Download() (*file.File, error) {
    uuid, err := d.client.BeginDownload(d.companyID, d.cookie)
    if err != nil {
        return nil, err
    }

    fmt.Println("Downloading prices", uuid)

    for {
        time.Sleep(time.Second * 5)
        status, errStatus := d.client.Status(uuid, d.cookie, d.companyID)
        if errStatus != nil {
            return nil, errStatus
        }

        fmt.Println("Status:", status)
        if status.Status == "done" {
            break
        }
    }

    data, err := d.client.Download(uuid, d.cookie, d.companyID)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

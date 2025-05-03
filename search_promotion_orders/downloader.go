package search_promotion_orders

import (
    "errors"
    "fmt"
    "ozon_reports_downloader/file"
    "time"
)

type Downloader struct {
    client         *Client
    companyID      int64
    organizationID int64
    cookie         string
}

func NewDownloader(companyID int64, organizationID int64, cookie string) *Downloader {
    return &Downloader{
        client:         &Client{},
        companyID:      companyID,
        organizationID: organizationID,
        cookie:         cookie,
    }
}

func (d Downloader) Download() (*file.File, error) {
    uuid, err := d.client.BeginDownload(d.companyID, d.organizationID, d.cookie)
    if err != nil {
        return nil, err
    }

    fmt.Println("Downloading search promotion orders", uuid)

    var (
        reportList *ReportResponse
        errStatus  error
    )

    for {
        time.Sleep(time.Second * 10)
        reportList, errStatus = d.client.ReportsList(d.companyID, d.organizationID, d.cookie)
        if errStatus != nil {
            return nil, errStatus
        }

        items := reportList.Items
        if len(items) == 0 {
            return nil, errors.New("items list is empty")
        }

        fmt.Println("Status:", items[0].Meta.State)
        if items[0].Meta.State == "OK" {
            break
        }
    }

    data, err := d.client.Download(reportList.Items[0].Meta.UUID, d.companyID, d.organizationID, d.cookie)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "csv"), nil
}

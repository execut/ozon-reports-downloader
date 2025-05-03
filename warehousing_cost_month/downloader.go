package warehousing_cost_month

import (
    "ozon_reports_downloader/file"
)

type Downloader struct {
    client *Client
}

func NewDownloader() *Downloader {
    return &Downloader{
        client: &Client{},
    }
}

func (d Downloader) Download(companyID int64, cookie string) (*file.File, error) {
    data, err := d.client.Download(companyID, cookie)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

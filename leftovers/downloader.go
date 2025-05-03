package leftovers

import (
    "ozon_reports_downloader/file"
)

type Downloader struct {
    client    *Client
    cookie    string
    companyID int64
}

func NewDownloader(cookie string, companyID int64) *Downloader {
    return &Downloader{
        client:    &Client{},
        cookie:    cookie,
        companyID: companyID,
    }
}

func (d Downloader) Download() (*file.File, error) {
    data, err := d.client.Download(d.cookie, d.companyID)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

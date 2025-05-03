package trafarets_detalization

import (
    "ozon_reports_downloader/file"
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
    data, err := d.client.Download(d.companyID, d.organizationID, d.cookie)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

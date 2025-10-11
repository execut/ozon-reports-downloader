package trafarets_detalization

import (
    "github.com/execut/ozon-reports-downloader/file"
)

type Downloader struct {
    client         *Client
    companyID      int64
    organizationID int64
    cookie         string
}

func NewDownloader(companyID int64, organizationID int64, cookie string, client *Client) *Downloader {
    return &Downloader{
        client:         client,
        companyID:      companyID,
        organizationID: organizationID,
        cookie:         cookie,
    }
}

func (d Downloader) Download() (*file.File, error) {
    data, err := d.client.Download()
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

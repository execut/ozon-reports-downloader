package warehousing_cost

import (
    "github.com/execut/ozon-reports-downloader/file"
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
    data, err := d.client.Download(d.companyID, d.cookie)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

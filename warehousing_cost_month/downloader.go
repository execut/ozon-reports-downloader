package warehousing_cost_month

import (
    "github.com/execut/ozon-reports-downloader/file"
)

type Downloader struct {
    client *Client
}

func NewDownloader(client *Client) *Downloader {
    return &Downloader{
        client: client,
    }
}

func (d Downloader) Download() (*file.File, error) {
    data, err := d.client.Download()
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

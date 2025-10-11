package prices

import (
    "fmt"
    "time"

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
    uuid, err := d.client.BeginDownload()
    if err != nil {
        return nil, err
    }

    fmt.Println("Downloading prices", uuid)

    for {
        time.Sleep(time.Second * 5)
        status, errStatus := d.client.Status(uuid)
        if errStatus != nil {
            return nil, errStatus
        }

        fmt.Println("Status:", status)
        if status.Status == "done" {
            break
        }
    }

    data, err := d.client.Download(uuid)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

package returns

import (
    "fmt"
    "io"
    "time"

    "github.com/execut/ozon-reports-downloader/file"
)

type Downloader struct {
    client      *Client
    returnsType ReturnsType
}

func NewDownloader(returnsType ReturnsType, client *Client) *Downloader {
    return &Downloader{
        client:      client,
        returnsType: returnsType,
    }
}

func (d Downloader) Download() (*file.File, error) {
    date := time.Now()
    uuid, err := d.client.BeginDownload(d.returnsType, date)
    if err != nil {
        return nil, err
    }

    fmt.Println("Downloading returns "+d.returnsType, uuid)

    var (
        status    *StatusResponse
        errStatus error
    )

    for {
        time.Sleep(time.Second)
        status, errStatus = d.client.Status(uuid)
        if errStatus != nil {
            return nil, errStatus
        }

        fmt.Println("Status:", status)
        if status.Status == "complete" {
            break
        }
    }

    reader, err := d.client.Download(status.Link)
    if err != nil {
        return nil, err
    }

    data, err := io.ReadAll(reader)
    if err != nil {
        return nil, err
    }

    return file.NewFile(data, "xlsx"), nil
}

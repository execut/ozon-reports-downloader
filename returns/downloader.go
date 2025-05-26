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
    companyID   int64
    cookie      string
}

func NewDownloader(returnsType ReturnsType, companyID int64, cookie string) *Downloader {
    return &Downloader{
        client:      &Client{},
        returnsType: returnsType,
        companyID:   companyID,
        cookie:      cookie,
    }
}

func (d Downloader) Download() (*file.File, error) {
    uuid, err := d.client.BeginDownload(d.companyID, d.returnsType, d.cookie)
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
        status, errStatus = d.client.Status(uuid, d.companyID, d.cookie)
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

package common

import (
    "fmt"
    "os"
    "ozon_reports_downloader/file"
)

type IDownloader interface {
    Download() (*file.File, error)
}

type Report struct {
    key        string
    downloader IDownloader
}

func NewReport(key string, downloader IDownloader) *Report {
    return &Report{key, downloader}
}

var counter = 0

func (r *Report) Run() error {
    counter++
    key := r.key
    key = fmt.Sprintf("%02d-", counter) + key
    fmt.Println("Begin report: " + key)
    ordersFile, err := r.downloader.Download()
    if err != nil {
        return err
    }

    err = r.saveFile(key, ordersFile)
    if err != nil {
        return err
    }

    return nil
}

func (r *Report) saveFile(fileName string, fileForSave *file.File) error {
    path := "reports/" + fileName + "." + fileForSave.FileType()
    fmt.Println("Saving report: " + path)
    fo, err := os.Create(path)
    if err != nil {
        return err
    }

    defer func() {
        if err := fo.Close(); err != nil {
            panic(err)
        }
    }()

    if _, err := fo.Write(fileForSave.Content()); err != nil {
        return err
    }

    return nil
}

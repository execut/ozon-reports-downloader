package common

import (
    "fmt"
    "os"

    "github.com/execut/ozon-reports-downloader/file"
)

type IDownloader interface {
    Download() (*file.File, error)
}

type Report struct {
    key        string
    downloader IDownloader
}

func (r *Report) Key() string {
    return r.key
}

func NewReport(key string, downloader IDownloader) *Report {
    return &Report{key, downloader}
}

func (r *Report) Run() error {
    fmt.Println("Begin report: " + r.key)
    ordersFile, err := r.downloader.Download()
    if err != nil {
        return err
    }

    err = r.saveFile(ordersFile)
    if err != nil {
        return err
    }

    return nil
}

func (r *Report) saveFile(fileForSave *file.File) error {
    path := "reports/" + r.key + "." + fileForSave.FileType()
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

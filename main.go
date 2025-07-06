package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"

    "github.com/execut/ozon-reports-downloader/accruals"
    "github.com/execut/ozon-reports-downloader/analytics"
    "github.com/execut/ozon-reports-downloader/common"
    "github.com/execut/ozon-reports-downloader/leftovers"
    "github.com/execut/ozon-reports-downloader/orders"
    "github.com/execut/ozon-reports-downloader/prices"
    "github.com/execut/ozon-reports-downloader/returns"
    "github.com/execut/ozon-reports-downloader/trafarets_detalization"
    "github.com/execut/ozon-reports-downloader/warehousing_cost"
    "gopkg.in/yaml.v3"
)

func main() {
    if len(os.Args) < 2 {
        panic("prev analytics date as first argument needed")
    }

    argsWithProg := os.Args[1]
    prevDate, err := time.Parse(time.DateOnly, argsWithProg)
    if err != nil {
        panic(err)
    }

    fmt.Printf("Begin report download for date %v\n", prevDate)

    config := ReadConfig()
    reports := []*common.Report{
        common.NewReport("trafarets", trafarets_detalization.NewDownloader(config.CompanyID, config.OrganizationID, config.Cookie)),
        //common.NewReport("search-promotion-orders", search_promotion_orders.NewDownloader(config.CompanyID, config.OrganizationID, config.Cookie)),
        common.NewReport("orders-fbo", orders.NewDownloader(common.DeliveryTypeFBO, config.CompanyID, config.Cookie, time.Now())),
        common.NewReport("orders-fbs", orders.NewDownloader(common.DeliveryTypeFBS, config.CompanyID, config.Cookie, time.Now())),
        common.NewReport("returns-fbos", returns.NewDownloader(returns.ReturnsTypeFBOS, config.CompanyID, config.Cookie)),
        common.NewReport("returns-realfbs", returns.NewDownloader(returns.ReturnsTypeRealFBS, config.CompanyID, config.Cookie)),
        common.NewReport("analytics", analytics.NewDownloader(prevDate, config.Cookie, config.CompanyID)),
        common.NewReport("accruals", accruals.NewDownloader(config.Cookie, config.CompanyID)),
        common.NewReport("leftovers", leftovers.NewDownloader(config.Cookie, config.CompanyID)),
        common.NewReport("warehousing-cost", warehousing_cost.NewDownloader(config.CompanyID, config.Cookie)),
        common.NewReport("prices", prices.NewDownloader(config.CompanyID, config.Cookie)),
    }

    for _, report := range reports {
        err := report.Run()
        if err != nil {
            panic(err)
        }

        time.Sleep(time.Second)
    }
}

func ReadConfig() Config {
    filename, _ := filepath.Abs("config.yml")
    yamlFile, err := os.ReadFile(filename)

    if err != nil {
        panic(err)
    }

    var config Config

    err = yaml.Unmarshal(yamlFile, &config)
    if err != nil {
        panic(err)
    }

    return config
}

type Config struct {
    Cookie         string `yaml:"cookie"`
    CompanyID      int64  `yaml:"companyID"`
    OrganizationID int64  `yaml:"organizationID"`
}

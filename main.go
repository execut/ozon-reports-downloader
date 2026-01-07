package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"

    "github.com/execut/ozon-reports-downloader/accruals"
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
    client := common.NewClient(config.Cookie, config.CompanyID, config.OrganizationID, config.SecChUa, config.UserAgent)
    reports := []*common.Report{
        common.NewReport("1-trafarets", trafarets_detalization.NewDownloader(config.CompanyID, config.OrganizationID, config.Cookie, trafarets_detalization.NewClient(client))),
        //common.NewReport("search-promotion-orders", search_promotion_orders.NewDownloader(config.CompanyID, config.OrganizationID, config.Cookie)),
        common.NewReport("2-orders-fbo", orders.NewDownloader(common.DeliveryTypeFBO, time.Now(), orders.NewClient(client, config.CompanyID))),
        common.NewReport("3-orders-fbs", orders.NewDownloader(common.DeliveryTypeFBS, time.Now(), orders.NewClient(client, config.CompanyID))),
        common.NewReport("4-returns-fbos", returns.NewDownloader(returns.ReturnsTypeFBOS, returns.NewClient(client, config.CompanyID))),
        common.NewReport("5-returns-realfbs", returns.NewDownloader(returns.ReturnsTypeRealFBS, returns.NewClient(client, config.CompanyID))),
        //common.NewReport("analytics", analytics.NewDownloader(prevDate, analytics.NewClient(client))),
        common.NewReport("6-accruals", accruals.NewDownloader(accruals.NewClient(client, config.CompanyID))),
        common.NewReport("7-leftovers", leftovers.NewDownloader(leftovers.NewClient(client))),
        common.NewReport("8-warehousing-cost", warehousing_cost.NewDownloader(warehousing_cost.NewClient(config.CompanyID, client))),
        common.NewReport("9-prices", prices.NewDownloader(prices.NewClient(client, config.CompanyID))),
    }

    for _, report := range reports {
        err := report.Run()
        if err != nil {
            log.Fatalf("report %s failed. Error: %e", report.Key(), err)
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
    SecChUa        string `yaml:"secChUa"`
    UserAgent      string `yaml:"userAgent"`
}

package common

import (
    "errors"
    "net/http"
    "strconv"

    "github.com/enetx/g"
    "github.com/enetx/surf"
)

var EmptyData = struct{}{}

type Client struct {
    httpClient     *http.Client
    surfClient     *surf.Client
    cookie         string
    companyID      int64
    organizationID int64
    secChUa        string
    userAgent      string
}

func NewClient(cookie string, companyID int64, organizationID int64, secChUa string, userAgent string) *Client {
    return &Client{cookie: cookie, companyID: companyID, organizationID: organizationID, secChUa: secChUa, userAgent: userAgent}
}

func (c *Client) DoRequest(data interface{}, url string) ([]byte, error) {
    return c.doRequest(data, url, http.MethodPost)
}

func (c *Client) DoGetRequest(data interface{}, url string) ([]byte, error) {
    return c.doRequest(data, url, http.MethodGet)
}

func (c *Client) DoPostPerformanceRequest(data interface{}, url string) ([]byte, error) {
    headerValueList := map[string]string{
        "Accept":                        "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
        "Accept-Language":               "ru",
        "Cache-control":                 "no-cache",
        "Content-Type":                  "application/json",
        "Cookie":                        c.cookie,
        "Origin":                        "https://seller.ozon.ru",
        "Priority":                      "u=1, i",
        "Sec-Ch-Ua":                     c.secChUa,
        "Sec-Ch-Ua-Mobile":              "?0",
        "Sec-Ch-Ua-Platform":            "\"Linux\"",
        "Sec-Fetch-Dest":                "empty",
        "Sec-Fetch-Mode":                "cors",
        "Sec-Fetch-Site":                "same-origin",
        "User-Agent":                    c.userAgent,
        "x-o3-adv-current-organisation": strconv.FormatInt(c.organizationID, 10),
        "X-O3-App-Name":                 "performance-sc",
        "X-O3-Company-Id":               strconv.FormatInt(c.companyID, 10),
        "X-O3-Language":                 "ru",
    }

    surfClient := surf.NewClient().
        Builder().
        Impersonate().Chrome().
        Session().
        SetHeaders(headerValueList).
        Build()

    resp := surfClient.Post(g.String(url), data).Do()
    err := resp.Err()
    if err != nil {
        return nil, err
    }

    if !resp.IsOk() {
        return nil, errors.New("wrong response status")
    }

    bodyBytes := resp.Ok().Body.Bytes()
    return bodyBytes, nil
}

//func (c *Client) DoGetPerformanceRequest(url string) ([]byte, error) {
//    // TODO need movement to serf
//    req, err := http.NewRequest("GET", url, nil)
//    if err != nil {
//        return nil, err
//    }
//
//    req.Header.Set("Accept", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
//    req.Header.Set("Accept-Language", "ru")
//    req.Header.Set("Cookie", c.cookie)
//    req.Header.Set("Origin", "https://seller.ozon.ru")
//    req.Header.Set("Priority", "u=1, i")
//    req.Header.Set("referer", "https://seller.ozon.ru/")
//    req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Google Chrome\";v=\"128\"")
//    req.Header.Set("sec-ch-ua-mobile", "?0")
//    req.Header.Set("sec-ch-ua-platform", "\"Linux\"")
//    req.Header.Set("sec-fetch-dest", "empty")
//    req.Header.Set("sec-fetch-mode", "cors")
//    req.Header.Set("sec-fetch-site", "same-site")
//    req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")
//    req.Header.Set("x-o3-adv-current-organisation", strconv.FormatInt(c.organizationID, 10))
//    req.Header.Set("x-o3-app-name", "performance-sc")
//    req.Header.Set("x-o3-company-id", strconv.FormatInt(c.companyID, 10))
//    req.Header.Set("x-o3-language", "ru")
//
//    resp, err := http.DefaultClient.Do(req)
//    if err != nil {
//        return nil, err
//    }
//    defer resp.Body.Close()
//
//    if resp.StatusCode != http.StatusOK {
//        return nil, errors.New("wrong response status " + fmt.Sprint(resp.StatusCode))
//    }
//
//    bodyBytes, err := io.ReadAll(resp.Body)
//    if err != nil {
//        return nil, err
//    }
//
//    return bodyBytes, nil
//}

func (c *Client) doRequest(data interface{}, url string, requestType string) ([]byte, error) {
    headerValueList := map[string]string{
        "Accept":             "application/json, text/plain, */*",
        "Accept-Language":    "ru",
        "Cache-control":      "no-cache",
        "Content-Type":       "application/json",
        "Cookie":             c.cookie,
        "Origin":             "https://seller.ozon.ru",
        "pragma":             "no-cache",
        "Priority":           "u=1, i",
        "Sec-Ch-Ua":          c.secChUa,
        "Sec-Ch-Ua-Mobile":   "?0",
        "Sec-Ch-Ua-Platform": "\"Linux\"",
        "Sec-Fetch-Dest":     "empty",
        "Sec-Fetch-Mode":     "cors",
        "Sec-Fetch-Site":     "same-origin",
        "User-Agent":         c.userAgent,
        "X-O3-App-Name":      "seller-ui",
        "X-O3-Company-Id":    strconv.FormatInt(c.companyID, 10),
        "X-O3-Language":      "ru",
        "X-O3-Page-Type":     "fulfillmentReports",
    }

    surfClient := surf.NewClient().
        Builder().
        Impersonate().Chrome().
        Session().
        SetHeaders(headerValueList).
        Build()

    var req *surf.Request
    if requestType == http.MethodPost {
        req = surfClient.Post(g.String(url), data)
    } else {
        if data == EmptyData {
            req = surfClient.Get(g.String(url))
        } else {
            req = surfClient.Get(g.String(url), data)
        }
    }

    resp := req.Do()
    err := resp.Err()
    if err != nil {
        return nil, err
    }

    if !resp.IsOk() {
        return nil, errors.New("wrong response status")
    }

    bodyBytes := resp.Ok().Body.Bytes()
    return bodyBytes, nil
}

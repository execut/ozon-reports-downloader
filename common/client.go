package common

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "net/http"
    "net/http/httputil"
    "strconv"
)

const (
    secChUa   = "\"Not)A;Brand\";v=\"8\", \"Chromium\";v=\"138\", \"Google Chrome\";v=\"138\""
    userAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36"
)

func DoRequest(data interface{}, url string, cookie string, companyID int64) ([]byte, error) {
    return doRequest(data, url, http.MethodPost, cookie, companyID)
}

func DoGetRequest(data interface{}, url string, cookie string, companyID int64) ([]byte, error) {
    return doRequest(data, url, http.MethodGet, cookie, companyID)
}

type loggingTransport struct{}

func (s *loggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
    bytes, _ := httputil.DumpRequestOut(r, true)
    transport := http.DefaultTransport

    resp, err := transport.RoundTrip(r)
    if err != nil {
        return nil, err
    }

    respBytes, _ := httputil.DumpResponse(resp, true)
    bytes = append(bytes, respBytes...)

    fmt.Printf("%s\n", bytes)

    return resp, err
}

func doRequest(data interface{}, url string, requestType string, cookie string, companyID int64) ([]byte, error) {
    client := http.DefaultClient
    payloadBytes, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    body := bytes.NewReader(payloadBytes)
    req, err := http.NewRequest(requestType, url, body)
    if err != nil {
        return nil, err
    }

    header := req.Header
    headerValueList := map[string]string{
        "Accept":             "application/json, text/plain, */*",
        "Accept-Language":    "ru",
        "Cache-control":      "no-cache",
        "Content-Type":       "application/json",
        "Cookie":             cookie,
        "Origin":             "https://seller.ozon.ru",
        "pragma":             "no-cache",
        "Priority":           "u=1, i",
        "Sec-Ch-Ua":          secChUa,
        "Sec-Ch-Ua-Mobile":   "?0",
        "Sec-Ch-Ua-Platform": "\"Linux\"",
        "Sec-Fetch-Dest":     "empty",
        "Sec-Fetch-Mode":     "cors",
        "Sec-Fetch-Site":     "same-origin",
        "User-Agent":         userAgent,
        "X-O3-App-Name":      "seller-ui",
        "X-O3-Company-Id":    strconv.FormatInt(companyID, 10),
        "X-O3-Language":      "ru",
        "X-O3-Page-Type":     "fulfillmentReports",
    }
    for k, v := range headerValueList {
        header.Set(k, v)
    }

    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("wrong response status " + fmt.Sprint(resp.StatusCode))
    }

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

func DoPostPerformanceRequest(data interface{}, url string, companyID int64, organizationID int64, cookie string) ([]byte, error) {
    payloadBytes, err := json.Marshal(data)
    if err != nil {
        return nil, err
    }

    body := bytes.NewReader(payloadBytes)
    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        return nil, err
    }

    header := req.Header
    headerValueList := map[string]string{
        "Accept":                        "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
        "Accept-Language":               "ru",
        "Cache-control":                 "no-cache",
        "Content-Type":                  "application/json",
        "Cookie":                        cookie,
        "Origin":                        "https://seller.ozon.ru",
        "Priority":                      "u=1, i",
        "Sec-Ch-Ua":                     secChUa,
        "Sec-Ch-Ua-Mobile":              "?0",
        "Sec-Ch-Ua-Platform":            "\"Linux\"",
        "Sec-Fetch-Dest":                "empty",
        "Sec-Fetch-Mode":                "cors",
        "Sec-Fetch-Site":                "same-origin",
        "User-Agent":                    userAgent,
        "x-o3-adv-current-organisation": strconv.FormatInt(organizationID, 10),
        "X-O3-App-Name":                 "performance-sc",
        "X-O3-Company-Id":               strconv.FormatInt(companyID, 10),
        "X-O3-Language":                 "ru",
    }
    for k, v := range headerValueList {
        header.Set(k, v)
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("wrong response status " + fmt.Sprint(resp.StatusCode))
    }

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

func DoGetPerformanceRequest(url string, companyID int64, organizationID int64, cookie string) ([]byte, error) {
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Accept", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    req.Header.Set("Accept-Language", "ru")
    req.Header.Set("Cookie", cookie)
    req.Header.Set("Origin", "https://seller.ozon.ru")
    req.Header.Set("Priority", "u=1, i")
    req.Header.Set("referer", "https://seller.ozon.ru/")
    req.Header.Set("sec-ch-ua", "\"Chromium\";v=\"128\", \"Not;A=Brand\";v=\"24\", \"Google Chrome\";v=\"128\"")
    req.Header.Set("sec-ch-ua-mobile", "?0")
    req.Header.Set("sec-ch-ua-platform", "\"Linux\"")
    req.Header.Set("sec-fetch-dest", "empty")
    req.Header.Set("sec-fetch-mode", "cors")
    req.Header.Set("sec-fetch-site", "same-site")
    req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")
    req.Header.Set("x-o3-adv-current-organisation", strconv.FormatInt(organizationID, 10))
    req.Header.Set("x-o3-app-name", "performance-sc")
    req.Header.Set("x-o3-company-id", strconv.FormatInt(companyID, 10))
    req.Header.Set("x-o3-language", "ru")

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New("wrong response status " + fmt.Sprint(resp.StatusCode))
    }

    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    return bodyBytes, nil
}

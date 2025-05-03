package analytics

type StartPayload struct {
    Filters    []interface{} `json:"filters"`
    Metrics    []string      `json:"metrics"`
    Dimensions []string      `json:"dimensions"`
    DateFrom   string        `json:"date_from"`
    DateTo     string        `json:"date_to"`
    IsAction   bool          `json:"is_action"`
}

type StartResponse struct {
    Code string `json:"code"`
}

type StatusPayload struct {
    Code string `json:"code"`
}

type StatusResponse struct {
    Status    string `json:"status"`
    ErrorCode int64  `json:"error_code"`
}

type DownloadPayload struct {
    CompanyID int64  `json:"company_id"`
    Code      string `json:"code"`
}

type DownloadResult struct {
    Result struct {
        ContentType string `json:"content_type"`
        FileName    string `json:"file_name"`
        FileContent string `json:"file_content"`
    } `json:"result"`
}

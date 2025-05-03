package prices

type StartPayload struct {
    Visibility           string   `json:"visibility"`
    PriceColorIndex      []string `json:"price_color_index"`
    IsSuperEconomEnabled bool     `json:"is_super_econom_enabled"`
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
    Progress  int64  `json:"progress"`
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

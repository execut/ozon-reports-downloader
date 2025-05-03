package orders

import "time"

type StartPayload struct {
    Filter    Filter `json:"filter"`
    Lang      string `json:"lang"`
    With      With   `json:"with"`
    CompanyID string `json:"company_id"`
    SortDir   string `json:"sort_dir"`
}
type Filter struct {
    ProcessedAtTo   time.Time `json:"processed_at_to"`
    ProcessedAtFrom time.Time `json:"processed_at_from"`
    DeliverySchema  string    `json:"delivery_schema"`
}

type With struct {
    AnalyticsData bool `json:"analytics_data"`
    JewelryCodes  bool `json:"jewelry_codes"`
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

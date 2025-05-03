package returns

import "time"

type StartPayloadRealFBS struct {
	Parameters     PayloadParameters `json:"rfbs_parameters"`
	SellerId       int64             `json:"seller_id"`
	TimeZoneOffset int64             `json:"time_zone_offset"`
}

type StartPayloadFBOS struct {
	Parameters     PayloadParameters `json:"return_parameters"`
	SellerId       int64             `json:"seller_id"`
	TimeZoneOffset int64             `json:"time_zone_offset"`
}

type PayloadParameters struct {
	PreFilterType int       `json:"pre_filter_type"`
	DateFrom      time.Time `json:"date_from"`
	DateTo        time.Time `json:"date_to"`
}

type StartResponse struct {
	Code string `json:"code"`
}

type StatusPayload struct {
	Code           string `json:"code"`
	SellerId       int64  `json:"seller_id"`
	TimeZoneOffset int64  `json:"time_zone_offset"`
}

type StatusResponse struct {
	Status    string `json:"status"`
	ErrorCode int64  `json:"error_code"`
	Link      string `json:"link"`
	FileName  string `json:"file_name"`
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

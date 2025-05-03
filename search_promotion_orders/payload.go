package search_promotion_orders

import "time"

type StartDownloadPayload struct {
    TimeBounds StartDownloadTimeBounds `json:"timeBounds"`
}

type StartDownloadTimeBounds struct {
    From time.Time `json:"from"`
    To   time.Time `json:"to"`
}

type StartResponse struct {
    Code string `json:"uuid"`
}

type ReportResponse struct {
    Items []struct {
        Name string `json:"name"`
        Meta struct {
            UUID      string    `json:"UUID"`
            State     string    `json:"state"`
            CreatedAt time.Time `json:"createdAt"`
            UpdatedAt time.Time `json:"updatedAt"`
            Request   struct {
                From time.Time `json:"from"`
                To   time.Time `json:"to"`
            } `json:"request"`
            Kind string `json:"kind"`
            Link string `json:"link,omitempty"`
        } `json:"meta"`
        Title string `json:"title"`
    } `json:"items"`
    Total string `json:"total"`
}

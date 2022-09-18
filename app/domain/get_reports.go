package domain

type GetReportsRequest struct{}

type GetReportsResponse struct {
	Count   *int     `json:"count,omitempty"`
	Reports []Report `json:"reports" json:"reports,omitempty"`
}

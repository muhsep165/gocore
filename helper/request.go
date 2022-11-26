package helper

type RequestSearchFilter struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	Search   string `json:"search"`
	DateFrom string `json:"from"`
	DateTo   string `json:"to"`
}

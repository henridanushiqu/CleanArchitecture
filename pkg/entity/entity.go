package entity

type Country struct {
	ID         int    `json:"country_id"`
	Country    string `json:"country"`
	LastUpdate string `json:"last_update"`
}

type Pagination struct {
	RowsNumber string `json:"size"`
	PageNumber string `json:"page"`
}

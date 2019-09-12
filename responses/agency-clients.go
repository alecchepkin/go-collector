package responses

type AgencyClientsResp struct {
	Count int `json:"count"`
	Items []struct {
		AccessType string `json:"access_type"`
		Status     string `json:"status"`
		User       User   `json:"user"`
	} `json:"items"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

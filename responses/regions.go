package responses

type RegionsResp struct {
	Count int      `json:"count"`
	Items []Region `json:"items"`
	Limit int      `json:"limit"`
}

type Region struct {
	Flags    []string `json:"flags"`
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	ParentID int      `json:"parent_id"`
}

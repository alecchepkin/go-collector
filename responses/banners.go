package responses

type BannersResp struct {
	Count  int      `json:"count"`
	Items  []Banner `json:"items"`
	Offset int      `json:"offset"`
}

type Banner struct {
	CampaignID       int    `json:"campaign_id"`
	ID               int64  `json:"id"`
	ModerationStatus string `json:"moderation_status"`
}

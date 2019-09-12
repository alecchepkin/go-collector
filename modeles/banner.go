package modeles

import "collector-toolkit"

var (
	_ collector_toolkit.IRow = (*Banner)(nil)
)

type Banner struct {
	TableName struct{} `sql:"banner,alias:t" pg:",discard_unknown_columns"`

	BannerID         int64  `sql:"banner_id"`
	CampaignID       int    `sql:"campaign_id"`
	ModerationStatus string `sql:"moderation_status"`
	CabinetID        int    `sql:"cabinet_id"`
}

func (Banner) UniqFields() []string {
	return []string{
		"BannerID",
	}
}

func (Banner) UpdateFields() []string {
	return []string{
		"BannerID",
		"CampaignID",
		"ModerationStatus",
		"CabinetID",
	}
}

func (t Banner) Struct() interface{} {
	return t
}

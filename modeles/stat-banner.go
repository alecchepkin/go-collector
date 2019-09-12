package modeles

import (
	"collector-toolkit"
)

var _ collector_toolkit.IRow = (*StatBanner)(nil)

type StatBanner struct {
	TableName struct{} `sql:"stat_banner,alias:t" pg:",discard_unknown_columns"`

	BannerID  int    `sql:"banner_id"`
	CabinetID int    `sql:"cabinet_id"`
	Date      string `sql:"date"`
	Clicks    int    `sql:"clicks"`
	Goals     int    `sql:"goals"`
	Increment int    `sql:"increment"`
	Reach     int    `sql:"reach"`
	Shows     int    `sql:"shows"`
	Spent     string `sql:"spent"`
	Total     int    `sql:"total"`
}

func (StatBanner) UniqFields() []string {
	return []string{
		"BannerID",
		"Date",
	}
}

func (StatBanner) UpdateFields() []string {

	return []string{
		"BannerID",
		"CabinetID",
		"Date",
		"Clicks",
		"Goals",
		"Increment",
		"Reach",
		"Shows",
		"Spent",
		"Total",
	}
}

func (t StatBanner) Struct() interface{} {
	return t
}

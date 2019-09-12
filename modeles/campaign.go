package modeles

import "collector-toolkit"

var (
	_ collector_toolkit.IRow = (*Campaign)(nil)
)

type Campaign struct {
	TableName struct{} `sql:"campaign,alias:t" pg:",discard_unknown_columns"`

	CampaignID     int64   `sql:"campaign_id,pk"`
	CabinetID      int     `sql:"cabinet_id"`
	Name           string  `sql:"name,notnull"`
	Status         string  `sql:"status,notnull"`
	Updated        string  `sql:"updated,notnull"`
	Created        string  `sql:"created,notnull"`
	BudgetLimit    float64 `sql:"budget_limit,notnull"`
	BudgetLimitDay float64 `sql:"budget_limit_day,notnull"`
}

func (Campaign) UniqFields() []string {
	return []string{
		"CampaignID",
		"CabinetID",
	}
}

func (Campaign) UpdateFields() []string {
	return []string{
		"CampaignID",
		"CabinetID",
		"Name",
		"Status",
		"Updated",
		"Created",
		"BudgetLimit",
		"BudgetLimitDay",
	}
}

func (t Campaign) Struct() interface{} {
	return t
}

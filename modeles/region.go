package modeles

import "collector-toolkit"

var (
	_ collector_toolkit.IRow = (*Region)(nil)
)

type Region struct {
	TableName struct{} `sql:"region,alias:t" pg:",discard_unknown_columns"`

	RegionID int    `sql:"region_id"`
	Flags    string `sql:"flags"`
	Name     string `sql:"name,notnull"`
	ParentID int    `sql:"parent_id,notnull"`
}

func (Region) UniqFields() []string {
	return []string{
		"RegionID",
	}
}

func (Region) UpdateFields() []string {
	return []string{
		"RegionID",
		"Flags",
		"Name",
		"ParentID",
	}
}

func (t Region) Struct() interface{} {
	return t
}

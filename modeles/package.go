package modeles

import "collector-toolkit"

var (
	_ collector_toolkit.IRow = (*Package)(nil)
)

type Package struct {
	TableName struct{} `sql:"package,alias:t" pg:",discard_unknown_columns"`

	CabinetID         int      `sql:"cabinet_id,notnull"`
	PackageID         int      `sql:"package_id,notnull"`
	BannerFormatID    int      `sql:"banner_format_id,notnull"`
	Created           string   `sql:"created,notnull"`
	Description       string   `sql:"description,notnull"`
	Flags             []string `sql:"flags,notnull"`
	MaxPricePerUnit   string   `sql:"max_price_per_unit,notnull"`
	MaxUniqShowsLimit int      `sql:"max_uniq_shows_limit,notnull"`
	Name              string   `sql:"name,notnull,notnull"`
	Objective         []string `sql:"objective,notnull"`
	PadsTreeID        int      `sql:"pads_tree_id,notnull"`
	PaidEventType     int      `sql:"paid_event_type,notnull"`
	Price             string   `sql:"price,notnull"`
	PricedEventType   int      `sql:"priced_event_type,notnull"`
	RelatedPackageIds []int    `sql:"related_package_ids,notnull"`
	Status            string   `sql:"status,notnull"`
	Updated           string   `sql:"updated,notnull"`
	UrlType           string   `sql:"url_type,notnull"`
	UrlTypes          string   `sql:"url_types,notnull"`
}

func (Package) UniqFields() []string {
	return []string{
		"CabinetID",
		"PackageID",
	}
}

func (Package) UpdateFields() []string {
	return []string{
		"CabinetID",
		"PackageID",
		"BannerFormatID",
		"Created",
		"Description",
		"Flags",
		"MaxPricePerUnit",
		"MaxUniqShowsLimit",
		"Name",
		"Objective",
		"PadsTreeID",
		"PaidEventType",
		"Price",
		"PricedEventType",
		"RelatedPackageIds",
		"Status",
		"Updated",
		"UrlType",
		"UrlTypes",
	}
}

func (t Package) Struct() interface{} {
	return t
}

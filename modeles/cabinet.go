package modeles

import (
	"collector-toolkit"
)

var (
	_ collector_toolkit.IRow = (*Cabinet)(nil)
)

type Cabinet struct {
	TableName struct{} `sql:"cabinet,alias:t" pg:",discard_unknown_columns"`

	CabinetID    int     `sql:"cabinet_id,notnull"`
	AccountID    int     `sql:"account_id,notnull"`
	Username     string  `sql:"username,notnull"`
	Name         string  `sql:"name,notnull"`
	Email        string  `sql:"email,notnull"`
	Balance      float64 `sql:"balance,notnull"`
	Currency     string  `sql:"currency,notnull"`
	IsExternal   bool    `sql:"is_external,notnull"`
	IsAgency     bool    `sql:"is_agency,notnull"`
	ParentID     int     `sql:"parent_id,notnull"`
	ClientID     string  `sql:"client_id,notnull"`
	ClientSecret string  `sql:"client_secret,notnull"`
	RefreshToken string  `sql:"refresh_token,notnull"`
	AccessToken  string  `sql:"access_token,notnull"`

	Parent *Cabinet `sql:"-"`
}

func (Cabinet) UniqFields() []string {
	return []string{
		"CabinetID",
	}
}

func (Cabinet) UpdateFields() []string {
	return []string{
		"CabinetID",
		"AccountID",
		"Username",
		"Name",
		"Email",
		"Balance",
		"Currency",
		"ParentID",
		"AccessToken",
		"RefreshToken",
	}
}

func (t Cabinet) Struct() interface{} {
	return t
}

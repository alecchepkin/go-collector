package modeles

import "time"

type ExtClient struct {
	TableName struct{} `sql:"token_collector,alias:t" pg:",discard_unknown_columns"`

	ID               int       `sql:"id,pk"`
	State            string    `sql:"state"`
	Scope            string    `sql:"scope"`
	ClientName       string    `sql:"client_name"`
	ClientCabinetId  int       `sql:"client_cabinet_id"`
	DailyAccessToken string    `sql:"daily_access_token"`
	RefreshToken     string    `sql:"refresh_token"`
	CreatedAt        time.Time `sql:""`
	UpdatedAt        time.Time `sql:""`
}

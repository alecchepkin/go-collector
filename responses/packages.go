package responses

type PackagesResp struct {
	Items []Package `json:"items"`
}

type Package struct {
	UrlTypes struct {
		Primary [][]string `json:"primary"`
	} `json:"url_types"`
	MaxUniqShowsLimit int      `json:"max_uniq_shows_limit"`
	PricedEventType   int      `json:"priced_event_type"`
	MaxPricePerUnit   string   `json:"max_price_per_unit"`
	Description       string   `json:"description"`
	Created           string   `json:"created"`
	Price             string   `json:"price"`
	Updated           string   `json:"updated"`
	PaidEventType     int      `json:"paid_event_type"`
	Status            string   `json:"status"`
	UrlType           string   `json:"url_type"`
	PadsTreeID        int      `json:"pads_tree_id"`
	Objective         []string `json:"objective"`
	BannerFormatID    int      `json:"banner_format_id"`
	RelatedPackageIds []int    `json:"related_package_ids"`
	Flags             []string `json:"flags"`
	ID                int      `json:"id"`
	Name              string   `json:"name"`
}

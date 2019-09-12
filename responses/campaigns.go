package responses

type CampaignsResp struct {
	Count  int        `json:"count"`
	Items  []Campaign `json:"items"`
	Offset int        `json:"offset"`
}

type Campaign struct {
	Created string `json:"created"`
	ID      int    `json:"id"`
	Issues  []struct {
		Arguments struct {
		} `json:"arguments"`
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"issues"`
	Name       string `json:"name"`
	Objective  string `json:"objective"`
	PackageID  int    `json:"package_id"`
	Price      string `json:"price"`
	Status     string `json:"status"`
	Targetings struct {
		Age struct {
			AgeList []int `json:"age_list"`
			Expand  bool  `json:"expand"`
		} `json:"age"`
		Fulltime struct {
			Flags []interface{} `json:"flags"`
			Fri   []int         `json:"fri"`
			Mon   []int         `json:"mon"`
			Sat   []int         `json:"sat"`
			Sun   []int         `json:"sun"`
			Thu   []int         `json:"thu"`
			Tue   []int         `json:"tue"`
			Wed   []int         `json:"wed"`
		} `json:"fulltime"`
		MobileOperationSystems []int    `json:"mobile_operation_systems"`
		MobileTypes            []string `json:"mobile_types"`
		Pads                   []int    `json:"pads"`
		Regions                []int    `json:"regions"`
		Sex                    []string `json:"sex"`
	} `json:"targetings"`
	Updated string `json:"updated"`
}

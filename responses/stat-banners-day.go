package responses

type StatBannersDayResp struct {
	Items []struct {
		ID   int `json:"id"`
		Rows []struct {
			Date string `json:"date"`
			Base struct {
				Shows  int     `json:"shows"`
				Clicks int     `json:"clicks"`
				Goals  int     `json:"goals"`
				Spent  string  `json:"spent"`
				Cpm    string  `json:"cpm"`
				Cpc    string  `json:"cpc"`
				Cpa    string  `json:"cpa"`
				Ctr    float64 `json:"ctr"`
				Cr     float64 `json:"cr"`
			} `json:"base"`
			Events struct {
				OpeningApp          int `json:"opening_app"`
				OpeningPost         int `json:"opening_post"`
				MovingIntoGroup     int `json:"moving_into_group"`
				ClicksOnExternalURL int `json:"clicks_on_external_url"`
				LaunchingVideo      int `json:"launching_video"`
				Comments            int `json:"comments"`
				Joinings            int `json:"joinings"`
				Likes               int `json:"likes"`
				Shares              int `json:"shares"`
				Votings             int `json:"votings"`
				SendingForm         int `json:"sending_form"`
			} `json:"events"`
			Uniques struct {
				Reach     int `json:"reach"`
				Total     int `json:"total"`
				Increment int `json:"increment"`
				Frequency int `json:"frequency"`
			} `json:"uniques"`
			Video struct {
				Started              int    `json:"started"`
				Paused               int    `json:"paused"`
				ResumedAfterPause    int    `json:"resumed_after_pause"`
				FullscreenOn         int    `json:"fullscreen_on"`
				FullscreenOff        int    `json:"fullscreen_off"`
				SoundTurnedOff       int    `json:"sound_turned_off"`
				SoundTurnedOn        int    `json:"sound_turned_on"`
				Viewed10Seconds      int    `json:"viewed_10_seconds"`
				Viewed25Percent      int    `json:"viewed_25_percent"`
				Viewed50Percent      int    `json:"viewed_50_percent"`
				Viewed75Percent      int    `json:"viewed_75_percent"`
				Viewed100Percent     int    `json:"viewed_100_percent"`
				Viewed10SecondsRate  int    `json:"viewed_10_seconds_rate"`
				Viewed25PercentRate  int    `json:"viewed_25_percent_rate"`
				Viewed50PercentRate  int    `json:"viewed_50_percent_rate"`
				Viewed75PercentRate  int    `json:"viewed_75_percent_rate"`
				Viewed100PercentRate int    `json:"viewed_100_percent_rate"`
				DepthOfView          int    `json:"depth_of_view"`
				Viewed10SecondsCost  string `json:"viewed_10_seconds_cost"`
				Viewed25PercentCost  string `json:"viewed_25_percent_cost"`
				Viewed50PercentCost  string `json:"viewed_50_percent_cost"`
				Viewed75PercentCost  string `json:"viewed_75_percent_cost"`
				Viewed100PercentCost string `json:"viewed_100_percent_cost"`
			} `json:"video"`
			Viral struct {
				Impressions         int `json:"impressions"`
				Reach               int `json:"reach"`
				Total               int `json:"total"`
				Increment           int `json:"increment"`
				Frequency           int `json:"frequency"`
				OpeningApp          int `json:"opening_app"`
				OpeningPost         int `json:"opening_post"`
				MovingIntoGroup     int `json:"moving_into_group"`
				ClicksOnExternalURL int `json:"clicks_on_external_url"`
				LaunchingVideo      int `json:"launching_video"`
				Comments            int `json:"comments"`
				Joinings            int `json:"joinings"`
				Likes               int `json:"likes"`
				Shares              int `json:"shares"`
				Votings             int `json:"votings"`
				SendingForm         int `json:"sending_form"`
			} `json:"viral"`
			Carousel struct {
				Slide1Clicks int `json:"slide_1_clicks"`
				Slide1Shows  int `json:"slide_1_shows"`
				Slide2Clicks int `json:"slide_2_clicks"`
				Slide2Shows  int `json:"slide_2_shows"`
				Slide3Clicks int `json:"slide_3_clicks"`
				Slide3Shows  int `json:"slide_3_shows"`
				Slide4Clicks int `json:"slide_4_clicks"`
				Slide4Shows  int `json:"slide_4_shows"`
				Slide5Clicks int `json:"slide_5_clicks"`
				Slide5Shows  int `json:"slide_5_shows"`
				Slide6Clicks int `json:"slide_6_clicks"`
				Slide6Shows  int `json:"slide_6_shows"`
				Slide1Ctr    int `json:"slide_1_ctr"`
				Slide2Ctr    int `json:"slide_2_ctr"`
				Slide3Ctr    int `json:"slide_3_ctr"`
				Slide4Ctr    int `json:"slide_4_ctr"`
				Slide5Ctr    int `json:"slide_5_ctr"`
				Slide6Ctr    int `json:"slide_6_ctr"`
			} `json:"carousel"`
			AdOffers struct {
				OfferPostponed   int `json:"offer_postponed"`
				UploadReceipt    int `json:"upload_receipt"`
				EarnOfferRewards int `json:"earn_offer_rewards"`
			} `json:"ad_offers"`
			Tps struct {
				Tps string `json:"tps"`
				Tpd string `json:"tpd"`
			} `json:"tps"`
			Moat struct {
				Impressions               int `json:"impressions"`
				InView                    int `json:"in_view"`
				NeverFocused              int `json:"never_focused"`
				NeverVisible              int `json:"never_visible"`
				Never50PercVisible        int `json:"never_50_perc_visible"`
				Never1SecVisible          int `json:"never_1_sec_visible"`
				HumanImpressions          int `json:"human_impressions"`
				ImpressionsAnalyzed       int `json:"impressions_analyzed"`
				InViewPercent             int `json:"in_view_percent"`
				HumanAndViewablePerc      int `json:"human_and_viewable_perc"`
				NeverFocusedPercent       int `json:"never_focused_percent"`
				NeverVisiblePercent       int `json:"never_visible_percent"`
				Never50PercVisiblePercent int `json:"never_50_perc_visible_percent"`
				Never1SecVisiblePercent   int `json:"never_1_sec_visible_percent"`
				InViewDiffPercent         int `json:"in_view_diff_percent"`
				ActiveInViewTime          int `json:"active_in_view_time"`
				AttentionQuality          int `json:"attention_quality"`
			} `json:"moat"`
			Romi struct {
				Value        string `json:"value"`
				Romi         int    `json:"romi"`
				AdvCostShare int    `json:"adv_cost_share"`
			} `json:"romi"`
		} `json:"rows"`
		Total struct {
			Base struct {
				Shows  int     `json:"shows"`
				Clicks int     `json:"clicks"`
				Goals  int     `json:"goals"`
				Spent  string  `json:"spent"`
				Cpm    string  `json:"cpm"`
				Cpc    string  `json:"cpc"`
				Cpa    string  `json:"cpa"`
				Ctr    float64 `json:"ctr"`
				Cr     float64 `json:"cr"`
			} `json:"base"`
			Events struct {
				OpeningApp          int `json:"opening_app"`
				OpeningPost         int `json:"opening_post"`
				MovingIntoGroup     int `json:"moving_into_group"`
				ClicksOnExternalURL int `json:"clicks_on_external_url"`
				LaunchingVideo      int `json:"launching_video"`
				Comments            int `json:"comments"`
				Joinings            int `json:"joinings"`
				Likes               int `json:"likes"`
				Shares              int `json:"shares"`
				Votings             int `json:"votings"`
				SendingForm         int `json:"sending_form"`
			} `json:"events"`
			Uniques struct {
				Reach          int `json:"reach"`
				Total          int `json:"total"`
				Increment      int `json:"increment"`
				Frequency      int `json:"frequency"`
				FrequencyTotal int `json:"frequency_total"`
			} `json:"uniques"`
			Video struct {
				Started              int    `json:"started"`
				Paused               int    `json:"paused"`
				ResumedAfterPause    int    `json:"resumed_after_pause"`
				FullscreenOn         int    `json:"fullscreen_on"`
				FullscreenOff        int    `json:"fullscreen_off"`
				SoundTurnedOff       int    `json:"sound_turned_off"`
				SoundTurnedOn        int    `json:"sound_turned_on"`
				Viewed10Seconds      int    `json:"viewed_10_seconds"`
				Viewed25Percent      int    `json:"viewed_25_percent"`
				Viewed50Percent      int    `json:"viewed_50_percent"`
				Viewed75Percent      int    `json:"viewed_75_percent"`
				Viewed100Percent     int    `json:"viewed_100_percent"`
				Viewed10SecondsRate  int    `json:"viewed_10_seconds_rate"`
				Viewed25PercentRate  int    `json:"viewed_25_percent_rate"`
				Viewed50PercentRate  int    `json:"viewed_50_percent_rate"`
				Viewed75PercentRate  int    `json:"viewed_75_percent_rate"`
				Viewed100PercentRate int    `json:"viewed_100_percent_rate"`
				DepthOfView          int    `json:"depth_of_view"`
				Viewed10SecondsCost  string `json:"viewed_10_seconds_cost"`
				Viewed25PercentCost  string `json:"viewed_25_percent_cost"`
				Viewed50PercentCost  string `json:"viewed_50_percent_cost"`
				Viewed75PercentCost  string `json:"viewed_75_percent_cost"`
				Viewed100PercentCost string `json:"viewed_100_percent_cost"`
			} `json:"video"`
			Viral struct {
				Impressions         int `json:"impressions"`
				Reach               int `json:"reach"`
				Total               int `json:"total"`
				Increment           int `json:"increment"`
				Frequency           int `json:"frequency"`
				OpeningApp          int `json:"opening_app"`
				OpeningPost         int `json:"opening_post"`
				MovingIntoGroup     int `json:"moving_into_group"`
				ClicksOnExternalURL int `json:"clicks_on_external_url"`
				LaunchingVideo      int `json:"launching_video"`
				Comments            int `json:"comments"`
				Joinings            int `json:"joinings"`
				Likes               int `json:"likes"`
				Shares              int `json:"shares"`
				Votings             int `json:"votings"`
				SendingForm         int `json:"sending_form"`
			} `json:"viral"`
			Carousel struct {
				Slide1Clicks int `json:"slide_1_clicks"`
				Slide1Shows  int `json:"slide_1_shows"`
				Slide2Clicks int `json:"slide_2_clicks"`
				Slide2Shows  int `json:"slide_2_shows"`
				Slide3Clicks int `json:"slide_3_clicks"`
				Slide3Shows  int `json:"slide_3_shows"`
				Slide4Clicks int `json:"slide_4_clicks"`
				Slide4Shows  int `json:"slide_4_shows"`
				Slide5Clicks int `json:"slide_5_clicks"`
				Slide5Shows  int `json:"slide_5_shows"`
				Slide6Clicks int `json:"slide_6_clicks"`
				Slide6Shows  int `json:"slide_6_shows"`
				Slide1Ctr    int `json:"slide_1_ctr"`
				Slide2Ctr    int `json:"slide_2_ctr"`
				Slide3Ctr    int `json:"slide_3_ctr"`
				Slide4Ctr    int `json:"slide_4_ctr"`
				Slide5Ctr    int `json:"slide_5_ctr"`
				Slide6Ctr    int `json:"slide_6_ctr"`
			} `json:"carousel"`
			AdOffers struct {
				OfferPostponed   int `json:"offer_postponed"`
				UploadReceipt    int `json:"upload_receipt"`
				EarnOfferRewards int `json:"earn_offer_rewards"`
			} `json:"ad_offers"`
			Tps struct {
				Tps string `json:"tps"`
				Tpd string `json:"tpd"`
			} `json:"tps"`
			Moat struct {
				Impressions               int `json:"impressions"`
				InView                    int `json:"in_view"`
				NeverFocused              int `json:"never_focused"`
				NeverVisible              int `json:"never_visible"`
				Never50PercVisible        int `json:"never_50_perc_visible"`
				Never1SecVisible          int `json:"never_1_sec_visible"`
				HumanImpressions          int `json:"human_impressions"`
				ImpressionsAnalyzed       int `json:"impressions_analyzed"`
				InViewPercent             int `json:"in_view_percent"`
				HumanAndViewablePerc      int `json:"human_and_viewable_perc"`
				NeverFocusedPercent       int `json:"never_focused_percent"`
				NeverVisiblePercent       int `json:"never_visible_percent"`
				Never50PercVisiblePercent int `json:"never_50_perc_visible_percent"`
				Never1SecVisiblePercent   int `json:"never_1_sec_visible_percent"`
				InViewDiffPercent         int `json:"in_view_diff_percent"`
				ActiveInViewTime          int `json:"active_in_view_time"`
				AttentionQuality          int `json:"attention_quality"`
			} `json:"moat"`
			Romi struct {
				Value        string `json:"value"`
				Romi         int    `json:"romi"`
				AdvCostShare int    `json:"adv_cost_share"`
			} `json:"romi"`
		} `json:"total"`
	} `json:"items"`
	Total struct {
		Base struct {
			Shows  int     `json:"shows"`
			Clicks int     `json:"clicks"`
			Goals  int     `json:"goals"`
			Spent  string  `json:"spent"`
			Cpm    string  `json:"cpm"`
			Cpc    string  `json:"cpc"`
			Cpa    string  `json:"cpa"`
			Ctr    float64 `json:"ctr"`
			Cr     float64 `json:"cr"`
		} `json:"base"`
		Events struct {
			OpeningApp          int `json:"opening_app"`
			OpeningPost         int `json:"opening_post"`
			MovingIntoGroup     int `json:"moving_into_group"`
			ClicksOnExternalURL int `json:"clicks_on_external_url"`
			LaunchingVideo      int `json:"launching_video"`
			Comments            int `json:"comments"`
			Joinings            int `json:"joinings"`
			Likes               int `json:"likes"`
			Shares              int `json:"shares"`
			Votings             int `json:"votings"`
			SendingForm         int `json:"sending_form"`
		} `json:"events"`
		Video struct {
			Started              int    `json:"started"`
			Paused               int    `json:"paused"`
			ResumedAfterPause    int    `json:"resumed_after_pause"`
			FullscreenOn         int    `json:"fullscreen_on"`
			FullscreenOff        int    `json:"fullscreen_off"`
			SoundTurnedOff       int    `json:"sound_turned_off"`
			SoundTurnedOn        int    `json:"sound_turned_on"`
			Viewed10Seconds      int    `json:"viewed_10_seconds"`
			Viewed25Percent      int    `json:"viewed_25_percent"`
			Viewed50Percent      int    `json:"viewed_50_percent"`
			Viewed75Percent      int    `json:"viewed_75_percent"`
			Viewed100Percent     int    `json:"viewed_100_percent"`
			Viewed10SecondsRate  int    `json:"viewed_10_seconds_rate"`
			Viewed25PercentRate  int    `json:"viewed_25_percent_rate"`
			Viewed50PercentRate  int    `json:"viewed_50_percent_rate"`
			Viewed75PercentRate  int    `json:"viewed_75_percent_rate"`
			Viewed100PercentRate int    `json:"viewed_100_percent_rate"`
			DepthOfView          int    `json:"depth_of_view"`
			Viewed10SecondsCost  string `json:"viewed_10_seconds_cost"`
			Viewed25PercentCost  string `json:"viewed_25_percent_cost"`
			Viewed50PercentCost  string `json:"viewed_50_percent_cost"`
			Viewed75PercentCost  string `json:"viewed_75_percent_cost"`
			Viewed100PercentCost string `json:"viewed_100_percent_cost"`
		} `json:"video"`
		Viral struct {
			Impressions         int `json:"impressions"`
			OpeningApp          int `json:"opening_app"`
			OpeningPost         int `json:"opening_post"`
			MovingIntoGroup     int `json:"moving_into_group"`
			ClicksOnExternalURL int `json:"clicks_on_external_url"`
			LaunchingVideo      int `json:"launching_video"`
			Comments            int `json:"comments"`
			Joinings            int `json:"joinings"`
			Likes               int `json:"likes"`
			Shares              int `json:"shares"`
			Votings             int `json:"votings"`
			SendingForm         int `json:"sending_form"`
		} `json:"viral"`
		Carousel struct {
			Slide1Clicks int `json:"slide_1_clicks"`
			Slide1Shows  int `json:"slide_1_shows"`
			Slide2Clicks int `json:"slide_2_clicks"`
			Slide2Shows  int `json:"slide_2_shows"`
			Slide3Clicks int `json:"slide_3_clicks"`
			Slide3Shows  int `json:"slide_3_shows"`
			Slide4Clicks int `json:"slide_4_clicks"`
			Slide4Shows  int `json:"slide_4_shows"`
			Slide5Clicks int `json:"slide_5_clicks"`
			Slide5Shows  int `json:"slide_5_shows"`
			Slide6Clicks int `json:"slide_6_clicks"`
			Slide6Shows  int `json:"slide_6_shows"`
			Slide1Ctr    int `json:"slide_1_ctr"`
			Slide2Ctr    int `json:"slide_2_ctr"`
			Slide3Ctr    int `json:"slide_3_ctr"`
			Slide4Ctr    int `json:"slide_4_ctr"`
			Slide5Ctr    int `json:"slide_5_ctr"`
			Slide6Ctr    int `json:"slide_6_ctr"`
		} `json:"carousel"`
		AdOffers struct {
			OfferPostponed   int `json:"offer_postponed"`
			UploadReceipt    int `json:"upload_receipt"`
			EarnOfferRewards int `json:"earn_offer_rewards"`
		} `json:"ad_offers"`
		Tps struct {
			Tps string `json:"tps"`
			Tpd string `json:"tpd"`
		} `json:"tps"`
		Moat struct {
			Impressions               int `json:"impressions"`
			InView                    int `json:"in_view"`
			NeverFocused              int `json:"never_focused"`
			NeverVisible              int `json:"never_visible"`
			Never50PercVisible        int `json:"never_50_perc_visible"`
			Never1SecVisible          int `json:"never_1_sec_visible"`
			HumanImpressions          int `json:"human_impressions"`
			ImpressionsAnalyzed       int `json:"impressions_analyzed"`
			InViewPercent             int `json:"in_view_percent"`
			HumanAndViewablePerc      int `json:"human_and_viewable_perc"`
			NeverFocusedPercent       int `json:"never_focused_percent"`
			NeverVisiblePercent       int `json:"never_visible_percent"`
			Never50PercVisiblePercent int `json:"never_50_perc_visible_percent"`
			Never1SecVisiblePercent   int `json:"never_1_sec_visible_percent"`
			InViewDiffPercent         int `json:"in_view_diff_percent"`
			ActiveInViewTime          int `json:"active_in_view_time"`
			AttentionQuality          int `json:"attention_quality"`
		} `json:"moat"`
		Romi struct {
			Value        string `json:"value"`
			Romi         int    `json:"romi"`
			AdvCostShare int    `json:"adv_cost_share"`
		} `json:"romi"`
	} `json:"total"`
}

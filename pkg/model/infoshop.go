package model

type DataInfoShopsCrawled struct {
	DataInfoShops []DataInfoShop `json:"data"`
}

type DataInfoShop struct {
	DataInfoShopDetails `json:"data"`
}

type DataInfoShopDetails struct {
	Shopid         int    `json:"shopid"`
	Userid         int    `json:"userid"`
	LastActiveTime int    `json:"last_active_time"`
	Vacation       bool   `json:"vacation"`
	Place          string `json:"place"`
	Account        struct {
		Portrait string `json:"portrait"`
		Username string `json:"username"`
		Status   int    `json:"status"`
	} `json:"account"`
	IsShopeeVerified      bool        `json:"is_shopee_verified"`
	IsPreferredPlusSeller bool        `json:"is_preferred_plus_seller"`
	IsOfficialShop        bool        `json:"is_official_shop"`
	ShopLocation          string      `json:"shop_location"`
	ItemCount             int         `json:"item_count"`
	RatingStar            float64     `json:"rating_star"`
	ResponseRate          int         `json:"response_rate"`
	SessionInfo           interface{} `json:"session_info"`
	Name                  string      `json:"name"`
	Ctime                 int         `json:"ctime"`
	ResponseTime          int         `json:"response_time"`
	FollowerCount         int         `json:"follower_count"`
	ShowOfficialShopLabel bool        `json:"show_official_shop_label"`
	RatingBad             int         `json:"rating_bad"`
	RatingGood            int         `json:"rating_good"`
	RatingNormal          int         `json:"rating_normal"`
	SessionInfos          interface{} `json:"session_infos"`
	ShowShopeeMission     bool        `json:"show_shopee_mission"`
	Status                int         `json:"status"`
	IsIndividualSeller    interface{} `json:"is_individual_seller"`
	Is3Pf                 bool        `json:"is_3pf"`
}

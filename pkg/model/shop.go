package model

// TotalOfficialShop

type DataShopCrawled struct {
	Data []Data `json:"data"`
}

type Data struct {
	Data DataShop `json:"data"`
}

type DataShop struct {
	Total         int            `json:"total"`
	OfficialShops []OfficialShop `json:"official_shops"`
}

type OfficialShop struct {
	Userid           int    `json:"userid"`
	Username         string `json:"username"`
	Shopid           int    `json:"shopid"`
	ShopName         string `json:"shop_name"`
	Logo             string `json:"logo"`
	LogoPc           string `json:"logo_pc"`
	ShopCollectionID int    `json:"shop_collection_id"`
	Ctime            int    `json:"ctime"`
	BrandLabel       int    `json:"brand_label"`
}

// ShopDetail

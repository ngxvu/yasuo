package model

type ShopeeCateData struct {
	Data DataCat `json:"data"`
}

type DataCat struct {
	CategoryList []CategoryList `json:"category_list"`
}

type CategoryList struct {
	Catid       int    `json:"catid"`
	ParentCatid int    `json:"parent_catid"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Image       string `json:"image"`
	Level       int    `json:"level"`
	Children    []struct {
		Catid       int         `json:"catid"`
		ParentCatid int         `json:"parent_catid"`
		Name        string      `json:"name"`
		DisplayName string      `json:"display_name"`
		Image       string      `json:"image"`
		Level       int         `json:"level"`
		Children    interface{} `json:"children"`
	}
}

type DataCatCrawled struct {
	Data DataCatYs `json:"data"`
}

type DataCatYs struct {
	CategoryList []struct {
		Catid       int    `json:"catid"`
		ParentCatid int    `json:"parent_catid"`
		Name        string `json:"name"`
		DisplayName string `json:"display_name"`
		Image       string `json:"image"`
		Level       int    `json:"level"`
		Children    []struct {
			Catid       int         `json:"catid"`
			ParentCatid int         `json:"parent_catid"`
			Name        string      `json:"name"`
			DisplayName string      `json:"display_name"`
			Image       string      `json:"image"`
			Level       int         `json:"level"`
			Children    interface{} `json:"children"`
		} `json:"children"`
	} `json:"category_list"`
}

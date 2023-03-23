package model

type Category struct {
	CategoryId   string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	CategoryUrl  string `json:"categoryUrl"`
}

type CategoryList struct {
	Large  []Category `json:"large"`
	Medium []Category `json:"medium"`
	Small  []Category `json:"small"`
}

type Result struct {
	Result CategoryList `json:"result"`
}

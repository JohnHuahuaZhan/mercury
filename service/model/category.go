package model

//category_id category_name 有唯一性约束不为空
type BasicCategory struct {
	CategoryId   uint64 `json:"category_id"`
	CategoryName string `json:"category_name"`
}

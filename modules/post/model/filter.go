package postmodel

type Filter struct {
	CategoryId string `json:"category_id,omitempty" form:"category_id"`
	UserId     string `json:"user_id,omitempty" form:"user_id"`
}

package common

type listResponse struct {
	ListData interface{} `json:"listData"`
	Paging   interface{} `json:"paging,omitempty"`
	Filter   interface{} `json:"filter,omitempty"`
}
type detailResponse struct {
	DetailData interface{} `json:"detailData"`
}
type actionResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func NewListResponse(listData, paging, filter interface{}) *listResponse {
	return &listResponse{ListData: listData, Paging: paging, Filter: filter}
}

func NewDetailResponse(detailData interface{}) *detailResponse {
	return &detailResponse{DetailData: detailData}
}

func NewActionResponse(message string, data interface{}) *actionResponse {
	return &actionResponse{Message: message, Data: data}
}

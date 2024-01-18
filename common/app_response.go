package common

type sucessRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessResponse(data, paging, filter interface{}) *sucessRes {
	return &sucessRes{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessReponse(data interface{}) *sucessRes {
	return NewSuccessResponse(data, nil, nil)
}

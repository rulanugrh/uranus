package web

type ResponseSuccess struct {
	Code    int    `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
	Data    any    `json:"data" form:"data"`
}

type ResponseFailure struct {
	Code    int    `json:"code" form:"code"`
	Message string `json:"message" form:"message"`
}
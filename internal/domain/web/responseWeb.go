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

type ResponseSuccessAuth struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Token   string `json:"token"`
}
type ValidationList struct {
	Field string      `json:"field"`
	Error interface{} `json:"error"`
}

type ValidationError struct {
	Message string           `json:"message"`
	Errors  []ValidationList `json:"error"`
}

func (err ValidationError) Error() string {
	return err.Message
}

type WebValidationError struct {
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

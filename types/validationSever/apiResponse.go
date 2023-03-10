package validationServer

type ApiResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Desc   string      `json:"desc"`
	Data   interface{} `json:"data"`
}

package global

type APIResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type APIDataResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}
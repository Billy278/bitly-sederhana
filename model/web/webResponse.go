package web

type WebResponse struct {
	Code   int         `validate:"required" json:"code"`
	Status string      `validate:"required" json:"status"`
	Data   interface{} `validate:"required" json:"data"`
}

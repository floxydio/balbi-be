package models

type Response struct {
	Status  int         `json:"err_no"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

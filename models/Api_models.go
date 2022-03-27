package models

type ApiError struct {
	Param   string
	Message string
}

type ApiResponse struct {
	Status  int         `json:"status"`
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

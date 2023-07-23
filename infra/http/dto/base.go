package dto

type BaseFailure struct {
	Success bool   `json:"success" default:"false"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

type BaseSuccess struct {
	Success bool        `json:"success" default:"true"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

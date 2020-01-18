package helper

import (
	"github.com/gin-gonic/gin"
)

//ResponseData : response format
type ResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

//ErrorField : 422 error validation format
type ErrorField struct {
	ID      string `json:"id"`
	Value   string `json:"value"`
	Caused  string `json:"caused"`
	Message string `json:"message"`
}

//RespondJSON : send a proper response json to the client
func RespondJSON(w *gin.Context, status int, message string, payload interface{}, errors interface{}) {
	var res ResponseData

	if status >= 200 && status < 300 {
		res.Success = true
	}

	if len(message) != 0 {
		res.Message = message
	}

	if payload != nil {
		res.Data = payload
	}

	if errors != nil {
		res.Error = errors
	}

	w.JSON(status, res)
}

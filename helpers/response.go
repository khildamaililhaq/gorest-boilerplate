package helpers

import "strings"

type Response struct {
	Status bool 	`json:"status"`
	Message string `json:"message"`
	Error 	interface{} `json:"errors"`
	Data 	interface{} `json:"data"`
}

type ResponseEmpty struct {

}

func ResponseSuccess(status bool, message string, data interface{}) Response {
	res := Response{
		Status: status,
		Message: message,
		Error: nil,
		Data: data,
	}
	return res
}

func ResponseError(message string, err string, data interface{}) Response {
	splitedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splitedError,
		Data:    data,
	}
	return res
}
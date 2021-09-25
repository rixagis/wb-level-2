package handlers

import (
	"encoding/json"

	"github.com/rixagis/wb-level-2/develop/dev11/internal/app/models"
)

// структура для запаковки ответа с ошибкой
type errorResponse struct {
	ErrorMessage string `json:"error"`
}

// структура для запаковки ответа со списком событий
type resultResponse struct {
	Result []models.Event `json:"result"`
}

// структура для запаковки ответа с текстовым сообщением
type postSuccessResultResponse struct {
	Result string `json:"result"`
}

// MakeJSONErrorResponse создает json для ответа с ошибкой
func MakeJSONErrorResponse(msg string) []byte {
	result, err := json.Marshal(errorResponse{msg})
	if err != nil {
		panic("json error while making error response: " + err.Error())
	}
	return result
}

// MakeJSONResultResponse создает json для ответа со списком событий
func MakeJSONResultResponse(events []models.Event) []byte {
	result, err := json.Marshal(resultResponse{events})
	if err != nil {
		panic("json error while making result response: " + err.Error())
	}
	return result
}

// MakeJSONPostSuccessResultResponse создает json для ответа с текстовым сообщением
func MakeJSONPostSuccessResultResponse(msg string) []byte {
	result, err := json.Marshal(postSuccessResultResponse{msg})
	if err != nil {
		panic("json error while making error response: " + err.Error())
	}
	return result
}
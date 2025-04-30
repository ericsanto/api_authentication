package customerrors

import (
	"net/http"
	"time"
)

func BadRequest() RestError {

	message := "body da requisição é invalido"
	code := http.StatusBadRequest
	time := time.Now().Format(time.RFC3339)

	rest := RestError{
		Message:   message,
		Code:      uint(code),
		TimeStamp: time,
	}

	return rest
}

func InternalServerError() RestError {

	message := "erro interno no servidor"
	code := http.StatusInternalServerError
	time := time.Now().Format(time.RFC3339)

	rest := RestError{
		Message:   message,
		Code:      uint(code),
		TimeStamp: time,
	}

	return rest
}

func UnprocessableEntity(message interface{}) RestError {

	code := http.StatusUnprocessableEntity
	time := time.Now().Format(time.RFC3339)

	rest := RestError{
		Message:   message,
		Code:      uint(code),
		TimeStamp: time,
	}

	return rest
}

func Conflict(message interface{}) RestError {

	code := http.StatusConflict
	time := time.Now().Format(time.RFC3339)

	rest := RestError{
		Message:   message,
		Code:      uint(code),
		TimeStamp: time,
	}

	return rest
}

func NotFound(message interface{}) RestError {

	code := http.StatusNotFound
	time := time.Now().Format(time.RFC3339)

	rest := RestError{
		Message:   message,
		Code:      uint(code),
		TimeStamp: time,
	}

	return rest
}

func Unauthorized(message interface{}) RestError {

	code := http.StatusUnauthorized
	time := time.Now().Format(time.RFC3339)

	rest := RestError{
		Message:   message,
		Code:      uint(code),
		TimeStamp: time,
	}

	return rest
}

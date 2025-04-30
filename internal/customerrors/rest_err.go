package customerrors

import (
	"fmt"
)

type RestError struct {
	Message   interface{} `json:"message"`
	Code      uint        `json:"code"`
	TimeStamp string      `json:"timestamp"`
}

func (re *RestError) Error() string {
	return fmt.Sprintf("erro: %s, %d, %s", re.Message, re.Code, re.TimeStamp)
}

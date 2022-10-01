package messages

import "encoding/json"

type ApiMessage struct {
	ErrorCode int    `json:"ErrorCode"`
	Message   string `json:"Message"`
	Data      string `json:"Data"`
}

func GetApiMessageString(message ApiMessage) string {
	res, err := json.Marshal(message)
	if err != nil {
		return SMessageUnknownException
	}
	return string(res)
}

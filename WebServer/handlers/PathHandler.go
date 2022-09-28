package handlers

import (
	"WebServer/messages"
	"encoding/json"
	"fmt"
	"net/http"
)

type PathHandler struct {
}

func (h *PathHandler) InvalidPath(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(messages.OMessageInvalidPath)
	var res string
	if err != nil {
		res = messages.SMessageUnknownException
	} else {
		res = string(response)
	}
	fmt.Fprintf(w, res)
}

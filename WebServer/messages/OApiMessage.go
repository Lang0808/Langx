package messages

import "WebServer/errors"

var OMessageInvalidPath ApiMessage = ApiMessage{
	ErrorCode: errors.UNKNOWN_PATH,
	Message:   "Invalid Path",
	Data:      "",
}

var SMessageUnknownException string = "{'ErrorCode':1,\n'Message':'Unknown Exception',\n'Data':''}"

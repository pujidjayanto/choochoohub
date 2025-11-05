package delivery

type SuccessResponse struct {
	Message  string `json:"message"`
	Data     any    `json:"data"`
	Metadata any    `json:"metadata,omitempty"`
}

type ErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"errorCode"`
	Trace     string `json:"trace,omitempty"`
}

var SuccessMessage = "ok"

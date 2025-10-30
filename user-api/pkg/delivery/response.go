package delivery

type SuccessResponse struct {
	Message  string `json:"message"`
	Data     any    `json:"data"`
	Metadata any    `json:"metadata,omitempty"`
}

type ErrorResponse struct {
	ErrorCode int    `json:"errorCode"`
	Error     string `json:"error"`
	Trace     string `json:"trace,omitempty"`
}

var SuccessMessage = "ok"

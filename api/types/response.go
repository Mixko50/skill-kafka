package types

type Response struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func ErrorResponse(message string) Response {
	return Response{
		Status:  "error",
		Message: message,
	}
}

func MessageResponse(message string) Response {
	return Response{
		Status:  "success",
		Message: message,
	}
}

func SuccessResponse(data any) Response {
	return Response{
		Status: "success",
		Data:   data,
	}
}

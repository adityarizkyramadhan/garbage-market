package utils

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

func ResponseWhenFail(message string, body interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Body:    body,
	}
}

func ResponseWhenSuccess(message string, body interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Body:    body,
	}
}

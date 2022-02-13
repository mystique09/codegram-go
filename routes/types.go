package routes

type (
	Error struct {
		Message string `json:"message"`
	}
	Response struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

func NewError(message string) Error {
	return Error{Message: message}
}

func NewResponse(success bool, message string, data interface{}) Response {
	return Response{Success: success, Message: message, Data: data}
}

package response

type Response struct {
	Status  int         `json:"Status"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

type ResponseLogin struct {
	Token string `json:"token"`
}

func HandleSuccessResponse(data interface{}, message string, status int) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
func HandleFailedResponse(message string, status int) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    nil,
	}
}

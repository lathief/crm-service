package payload

type Response struct {
	ResponseMeta
	Data any `json:"Data"`
}
type ResponseMeta struct {
	Status  int    `json:"Status"`
	Message string `json:"Message"`
}
type ResponseLogin struct {
	Token string `json:"token"`
}

func HandleSuccessResponse(data interface{}, message string, status int) Response {
	return Response{
		ResponseMeta: ResponseMeta{
			Status:  status,
			Message: message,
		},
		Data: data,
	}
}
func HandleFailedResponse(message string, status int) Response {
	return Response{
		ResponseMeta: ResponseMeta{
			Status:  status,
			Message: message,
		},
	}
}

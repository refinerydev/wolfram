package helper

func HttpResponseError(statusCode int, errMsg interface{}) Response {
	response := ResponseFormatter(
		statusCode,
		"fail",
		errMsg,
		nil,
	)
	return response
}

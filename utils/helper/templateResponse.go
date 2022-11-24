package helper

func SuccessDataResponseHelper(httpError int, msg, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  httpError,
		"message": msg,
		"data":    data,
	}
}

func FailedResponseHelper(httpError int, msg interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  httpError,
		"message": msg,
	}
}

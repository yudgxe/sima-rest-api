package handler

type errorResponce struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

func newErrorResponce(status string, message interface{}) *errorResponce {
	return &errorResponce{
		Status:  status,
		Message: message,
	}
}

type successResponce struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func newSuccessResponce(status string, data interface{}) *successResponce {
	return &successResponce{
		Status: status,
		Data:   data,
	}
}

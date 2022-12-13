package http_response

import "encoding/json"

type Status struct {
	Code    int
	Message string
}

type CustomReponseSingle struct {
	Status *Status
	Data   *ResponseItemJson
}

func MapResponse(code int, message string) ([]byte, error) {
	httpResponse := &CustomReponseSingle{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: nil,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

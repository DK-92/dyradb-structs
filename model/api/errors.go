package api

import "github.com/go-playground/validator/v10"

type returnMessage struct {
	Information []string        `json:"information,omitempty"`
	Errors      []*errorMessage `json:"errors,omitempty"`
}

type errorMessage struct {
	Field   string `json:"field"`
	Type    string `json:"type"`
	Generic string `json:"message,omitempty"`
}

func DetailedErrorMessage(e validator.ValidationErrors) *returnMessage {
	var eArr []*errorMessage
	for _, err := range e {
		message := errorMessage{
			Field: err.Field(),
			Type:  err.Tag(),
		}
		eArr = append(eArr, &message)
	}

	return &returnMessage{Errors: eArr}
}

func GenericErrorMessage(e error) *returnMessage {
	message := []*errorMessage{{
		Generic: e.Error(),
	}}

	return &returnMessage{
		Errors: message,
	}
}

func InformationMessage(message ...string) *returnMessage {
	return &returnMessage{
		Information: message,
	}
}

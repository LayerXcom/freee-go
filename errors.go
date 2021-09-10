package freee

import "encoding/json"

const (
	UnauthorizedCodeInvalidAccessToken      = "invalid_access_token"
	UnauthorizedCodeExpiredAccessToken      = "expired_access_token"
	UnauthorizedCodeUserDoNotHavePermission = "user_do_not_have_permission"
	UnauthorizedCodeCompanyNotFound         = "company_not_found"
	UnauthorizedCodeFreeePlanLimit          = "freee_plan_limit"
	UnauthorizedCodeSourceIPAddressLimit    = "source_ip_address_limit"
)

type UnauthorizedError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type Error struct {
	StatusCode              int
	RawError                string
	IsAuthorizationRequired bool
}

func (e *Error) Error() string {
	return e.RawError
}

func (e *Error) Messages() []string {
	messages, _ := ExtractFreeeErrorMessage(e.RawError)
	return messages
}

type FreeErrorMessageDetail struct {
	Messages []string `json:"messages"`
}
type FreeeErrorMessage struct {
	ErrorDescription string                   `json:"error_description"`
	Message          string                   `json:"message"`
	Messages         []string                 `json:"messages"`
	ErrorDetails     []FreeErrorMessageDetail `json:"errors"`
}

func ExtractFreeeErrorMessage(errorString string) ([]string, error) {
	var messages []string

	var errorMessage FreeeErrorMessage
	if err := json.Unmarshal([]byte(errorString), &errorMessage); err != nil {
		return messages, err
	}

	if errorMessage.ErrorDescription != "" {
		messages = append(messages, errorMessage.ErrorDescription)
	}
	if errorMessage.Message != "" {
		messages = append(messages, errorMessage.Message)
	}
	for _, msg := range errorMessage.Messages {
		messages = append(messages, msg)
	}
	for _, errorDetail := range errorMessage.ErrorDetails {
		if len(errorDetail.Messages) > 0 {
			messages = append(messages, errorDetail.Messages...)
		}
	}

	return messages, nil
}

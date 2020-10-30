package freee

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

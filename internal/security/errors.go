package security

type TokenTimeoutError struct {
	Message string
}

type InvalidTokenError struct {
	Message string
}

func (e *InvalidTokenError) Error() string {
	return ""
}

func (e *TokenTimeoutError) Error() string {
	return ""
}

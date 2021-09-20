package session

type Service interface {
	// GetSession get session or nil
	GetSession(token string) *string
	// SetSession set session or return error
	SetSession(username string) (string, error)
}

package utils

const (
	ErrorAuthorizationMission = "authorization header missing"
	ErrorInvalidTokenFormat   = "invalid token format"
	ErrorInvalidToken         = "invalid token"
)

const (
	PasswordSecurityDefaul = "123mamaestapresa"
	TimeSessionMinute      = 15
)

const (
	CtxAccountID = iota + 1
	CtxAccountEmail
	CtxAccountName
)

package utils

const (
	//ACCOUNT
	ErrorPasswordEquals = "different password"
	ErrorEmailOrUser    = "user or email not found"

	//SECURITY
	ErrorAuthorizationMission = "authorization header missing"
	ErrorInvalidTokenFormat   = "invalid token format"
	ErrorInvalidToken         = "invalid token"
)

const (
	PasswordSecurityDefaul = "123mamaestapresa"
	TimeSessionMinute      = 15
)

const (
	NameCookieToken = "auth_token"
)

const (
	CtxAccountID = iota + 1
	CtxAccountEmail
	CtxAccountName
)

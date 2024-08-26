package utils

const (
	LimitRecordFive     = 5
	TimeCheckInfoServer = 1 //minutes
)
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
	NameAuthAPI     = "AuthorizationAPI"
)

const (
	CtxAccountID = iota + 1
	CtxAccountEmail
	CtxAccountName
	CtxTypeAccount
	CtxClaim
	CtxNavWeb
)

const (
	// Encryp passwd token
	UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCase = "abcdefghijklmnopqrstuvwxyz"
	Digits    = "0123456789"
	Special   = "!@#$%^&*()_-+=[]{}|;:,.<>?~"
)

const (
	IconNewsTicketCommunity uint8 = iota + 1
	IconNewsTicketDevelopment
	IconNewsTicketSupport
	IconNewsTicketTechnical
)

// Path Icon Target
const (
	PathIconNewsTicketCommunity   = "newsicon_community_small.png"
	PathIconNewsTicketDevelopment = "newsicon_development_small.png"
	PathIconNewsTicketSupport     = "newsicon_support_small.png"
	PathIconNewsTicketTechnical   = "newsicon_technical_small.png"
)

// Privileges User
// player == account privileges ID
const (
	UserPlayer int = iota + 1
	UserTutor
	UserSeniorTutor
	UserGameMaster
	UserCommunityManager
	UserGod
)

// privbi
const ()

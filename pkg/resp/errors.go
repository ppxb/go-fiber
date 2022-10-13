package resp

const (
	Ok                  = 201
	NotOk               = 405
	Unauthorized        = 401
	Forbidden           = 403
	InternalServerError = 500
)

const (
	OkMsg                      = "success"
	NotOkMsg                   = "failed"
	UnauthorizedMsg            = "login expired, please login again"
	InvalidParameterMsg        = "invalid parameter"
	IllegalParameterMsg        = "illegal parameter"
	LoginCheckErrorMsg         = "wrong username or password"
	ForbiddenMsg               = "no permission to access this resource"
	InternalServerErrorMsg     = "server internal error"
	IdempotenceTokenEmptyMsg   = "idempotent token is empty"
	IdempotenceTokenInvalidMsg = "idempotent token expired"
	UserDisabledMsg            = "the account has been disabled"
	WeakPassword               = "the password is too weak"
	UserLockedMsg              = "the account has been locked"
	InvalidCaptchaMsg          = "the verification code is invalid or expired"
	InvalidSignIdMsg           = "invalid app id"
	IllegalSignIdMsg           = "illegal app id"
	InvalidSignTokenMsg        = "invalid token"
	IllegalSignTokenMsg        = "illegal token"
	InvalidSignTimestampMsg    = "invalid timestamp"
	InvalidSignScopeMsg        = "invalid scope"
)

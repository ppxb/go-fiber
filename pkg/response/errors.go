package response

const (
	Ok                  = 200
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
	LoginCheckErrorMsg         = "错误的用户名或密码"
	AssetImportErrorMsg        = "数据导入失败"
	ForbiddenMsg               = "no permission to access this resource"
	InternalServerErrorMsg     = "服务器内部错误"
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

var ErrMsg = map[int]string{
	Ok:                  OkMsg,
	InternalServerError: InternalServerErrorMsg,
}

package constants

// success

const (
	SuccessCode = "0000"
)

// Authentication/Authorization errors (1XXX)
const (
	ErrCodeUnauthorized      = "1001"
	ErrCodeForbidden         = "1002"
	ErrCodeInvalidToken      = "1003"
	ErrCodeTokenExpired      = "1004"
	ErrCodeInvalidCredential = "1005"
)

// Validation errors (2XXX)
const (
	ErrCodeInvalidInput  = "2001"
	ErrCodeRequiredField = "2002"
	ErrCodeInvalidValue  = "2005"
)

// Resource errors (3XXX)
const (
	ErrCodeResourceNotFound    = "3001"
	ErrCodeResourceExists      = "3002"
	ErrCodeResourceUnavailable = "3003"
)

// Request errors (4XXX)
const (
	ErrCodeBadRequest           = "4001"
	ErrCodeMethodNotAllowed     = "4002"
	ErrCodeRequestTimeout       = "4003"
	ErrCodeTooManyRequests      = "4004"
	ErrCodeUnsupportedMediaType = "4005"
)

// Server errors (5XXX)
const (
	ErrCodeInternalServer     = "5001"
	ErrCodeServiceUnavailable = "5002"
	ErrCodeDatabaseError      = "5003"
	ErrCodeUnexpectedError    = "5004"
)

// Business logic errors (6XXX)
const (
	ErrCodeBusinessLogic     = "6001"
	ErrCodeInsufficientFunds = "6002"
	ErrCodeOperationFailed   = "6003"
)

// External service errors (7XXX)
const (
	ErrCodeExternalService     = "7001"
	ErrCodeExternalTimeout     = "7002"
	ErrCodeExternalUnavailable = "7003"
)

// Error code structure
type ErrorCode struct {
	Code       string
	CodeLabel  map[string]string // lang -> label
	HTTPStatus int               // Corresponding HTTP status (optional)
}

// Message map with descriptions in Thai and English
var MessageMap = map[string]ErrorCode{
	// Authentication/Authorization errors
	ErrCodeUnauthorized: {
		Code: ErrCodeUnauthorized,
		CodeLabel: map[string]string{
			LangEN: "Unauthorized access",
			LangTH: "ไม่ได้รับอนุญาตให้เข้าถึง",
		},
		HTTPStatus: 401,
	},
	ErrCodeForbidden: {
		Code: ErrCodeForbidden,
		CodeLabel: map[string]string{
			LangEN: "Access forbidden",
			LangTH: "การเข้าถึงถูกปฏิเสธ",
		},
		HTTPStatus: 403,
	},
	ErrCodeInvalidToken: {
		Code: ErrCodeInvalidToken,
		CodeLabel: map[string]string{
			LangEN: "Invalid authentication token",
			LangTH: "โทเค็นการยืนยันตัวตนไม่ถูกต้อง",
		},
		HTTPStatus: 401,
	},
	ErrCodeTokenExpired: {
		Code: ErrCodeTokenExpired,
		CodeLabel: map[string]string{
			LangEN: "Authentication token expired",
			LangTH: "โทเค็นการยืนยันตัวตนหมดอายุ",
		},
		HTTPStatus: 401,
	},
	ErrCodeInvalidCredential: {
		Code: ErrCodeInvalidCredential,
		CodeLabel: map[string]string{
			LangEN: "Invalid username or password",
			LangTH: "ชื่อผู้ใช้หรือรหัสผ่านไม่ถูกต้อง",
		},
		HTTPStatus: 401,
	},

	// Validation errors
	ErrCodeInvalidInput: {
		Code: ErrCodeInvalidInput,
		CodeLabel: map[string]string{
			LangEN: "Invalid input data",
			LangTH: "ข้อมูลที่ป้อนไม่ถูกต้อง",
		},
		HTTPStatus: 400,
	},
	ErrCodeRequiredField: {
		Code: ErrCodeRequiredField,
		CodeLabel: map[string]string{
			LangEN: "Required field missing",
			LangTH: "ข้อมูลที่จำเป็นไม่ครบถ้วน",
		},
		HTTPStatus: 400,
	},

	// Resource errors
	ErrCodeResourceNotFound: {
		Code: ErrCodeResourceNotFound,
		CodeLabel: map[string]string{
			LangEN: "Resource not found",
			LangTH: "ไม่พบข้อมูลที่ต้องการ",
		},
		HTTPStatus: 404,
	},
	ErrCodeResourceExists: {
		Code: ErrCodeResourceExists,
		CodeLabel: map[string]string{
			LangEN: "Resource already exists",
			LangTH: "ข้อมูลนี้มีอยู่แล้ว",
		},
		HTTPStatus: 409,
	},

	// Request errors
	ErrCodeBadRequest: {
		Code: ErrCodeBadRequest,
		CodeLabel: map[string]string{
			LangEN: "Bad request",
			LangTH: "คำขอไม่ถูกต้อง",
		},
		HTTPStatus: 400,
	},
	ErrCodeTooManyRequests: {
		Code: ErrCodeTooManyRequests,
		CodeLabel: map[string]string{
			LangEN: "Too many requests",
			LangTH: "คำขอมากเกินไป",
		},
		HTTPStatus: 429,
	},

	// Server errors
	ErrCodeInternalServer: {
		Code: ErrCodeInternalServer,
		CodeLabel: map[string]string{
			LangEN: "Internal server error",
			LangTH: "เซิร์ฟเวอร์ผิดพลาด",
		},
		HTTPStatus: 500,
	},
	ErrCodeServiceUnavailable: {
		Code: ErrCodeServiceUnavailable,
		CodeLabel: map[string]string{
			LangEN: "Service unavailable",
			LangTH: "บริการไม่พร้อมใช้งาน",
		},
		HTTPStatus: 503,
	},
	ErrCodeDatabaseError: {
		Code: ErrCodeDatabaseError,
		CodeLabel: map[string]string{
			LangEN: "Database error",
			LangTH: "ฐานข้อมูลผิดพลาด",
		},
		HTTPStatus: 500,
	},

	// Business logic errors
	ErrCodeBusinessLogic: {
		Code: ErrCodeBusinessLogic,
		CodeLabel: map[string]string{
			LangEN: "Business logic error",
			LangTH: "ข้อผิดพลาดทางธุรกิจ",
		},
		HTTPStatus: 422,
	},
	ErrCodeInsufficientFunds: {
		Code: ErrCodeInsufficientFunds,
		CodeLabel: map[string]string{
			LangEN: "Insufficient funds",
			LangTH: "ยอดเงินไม่เพียงพอ",
		},
		HTTPStatus: 422,
	},
}

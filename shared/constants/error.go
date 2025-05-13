package constants

const (
	ErrCode0001 = "0001"
	ErrCode0002 = "0002"
	ErrCode5000 = "5000"
)

type ErrorCode struct {
	Code      string
	CodeLabel map[string]string // lang -> label
}

var ErrMap = map[string]ErrorCode{
	ErrCode0001: {
		Code: ErrCode0001,
		CodeLabel: map[string]string{
			LangEN: "Error 0001",
			LangTH: "ผิดพลาด 0001",
		},
	},
	ErrCode0002: {
		Code: ErrCode0002,
		CodeLabel: map[string]string{
			LangEN: "Error 0002",
			LangTH: "ผิดพลาด 0002",
		},
	},
	ErrCode5000: {
		Code: ErrCode5000,
		CodeLabel: map[string]string{
			LangEN: "Internal server error",
			LangTH: "ระบบผิดพลาด",
		},
	},
}

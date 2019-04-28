package base

import "fmt"

type ErrorCode int

const (
	FGEBase = 10000
)

const (
	Failed  ErrorCode = -1
	Success ErrorCode = 200
)

const (
	EInternalError ErrorCode = 5*FGEBase + iota
)

const (
	ErrClientError ErrorCode = 4*FGEBase + iota
)

var evtDesc = map[ErrorCode]string{
	Failed:  "Failed",
	Success: "Success",

	EInternalError: "系统内部错误",

	ErrClientError: "客户端错误",
}

func (e ErrorCode) String() string {
	if value, ok := evtDesc[e]; ok {
		return value
	}
	return "Unknow"
}

func (e ErrorCode) Error() string {
	return fmt.Sprintf("%d:%s", int(e), e.String())
}

package resultcode

type Code int

const (
	OK               Code = 0
	InternalError    Code = 10000 // 系统错误
	InvalidArgument  Code = 10001 // 参数错误
	Unauthorized     Code = 10002 // 访问未授权
	PermissionDenied Code = 10003 // 权限错误
)

var msg = map[Code]string{
	OK:               "OK",
	InternalError:    "内部错误，请稍后重试",
	InvalidArgument:  "参数错误",
	Unauthorized:     "访问未授权",
	PermissionDenied: "权限不足",
}

func (c Code) String() string {
	return msg[c]
}

package response

type ErrorCode int

const (
	SystemError   ErrorCode = 7
	ParamError    ErrorCode = 1
	NetworkError  ErrorCode = 2
	NotFoundError ErrorCode = 404
	IdNotEmpty    ErrorCode = 50001
)

func (e ErrorCode) String() string {
	switch e {
	case SystemError:
		return "系统错误"
	case ParamError:
		return "参数错误"
	case NetworkError:
		return "网络错误"
	case NotFoundError:
		return "404"
	case IdNotEmpty:
		return "ID不能为空"
	default:
		return "未知错误"
	}
}

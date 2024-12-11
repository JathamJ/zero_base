package errx

// 通用错误码
const (
	Success         = 200   //成功
	InvalidParam    = 10001 //参数错误
	DbError         = 10002 //数据库错误
	RedisError      = 10003 //Redis异常
	ApiError        = 10004 //依赖接口异常
	SignError       = 10005 //签名错误
	PermissionError = 10006 //权限不足
	RequestNotAllow = 10007 //限制访问
	RecordNotFound  = 10008 //记录不存在
	FileNotExist    = 10009 //文件不存在
	SystemError     = 10010 //系统异常
	BusinessError   = 10011 //业务异常
)

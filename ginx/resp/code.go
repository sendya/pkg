package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorCode int

// 业务异常状态码标识
const (
	ok              ErrorCode = 0   // 请求正常状态码
	errBadRequest   ErrorCode = -1  // 请求参数错误
	errUnauthorized ErrorCode = -2  // 未授权
	errForbidden    ErrorCode = -3  // 访问拒绝
	errServer       ErrorCode = -10 // 服务异常
)

// 服务正常
func OK(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, NewResp(int(ok), "OK"))
}

// 服务正常 并返回 JSON 数据结构
func JSON(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, NewResp(int(ok), "OK").WithData(v))
}

// 服务正常 并返回分页数据，分页对象需要实现 resp.RespPage interface{}
func WithPage(ctx *gin.Context, v RespPage) {
	ctx.JSON(http.StatusOK, NewResp(int(ok), "").
		WithPage(v.GetData(), v.GetPageSize(), v.GetCurrent(), int(v.GetTotal())),
	)
}

// 服务正常 并返回分页数据结构
func JSONPage(ctx *gin.Context, v interface{}, pageSize, current int, total int64) {
	ctx.JSON(http.StatusOK, NewResp(int(ok), "OK").WithPage(v, pageSize, current, int(total)))
}

// 服务异常
func ERR(ctx *gin.Context, httpCode int, code int, arg interface{}) {
	var msg string
	switch arg := arg.(type) {
	case error:
		msg = arg.Error()
	case string:
		msg = arg
	default:
		msg = ""
	}
	ctx.JSON(httpCode, NewResp(code, msg))
}

func ERRServer(ctx *gin.Context, err error) bool {
	if err != nil {
		ERR(ctx, http.StatusInternalServerError, int(errServer), err.Error())
		return true
	}
	return false
}

func BadRequest(ctx *gin.Context, err error) bool {
	if err != nil {
		ERR(ctx, http.StatusBadRequest, int(errBadRequest), err.Error())
		return true
	}
	return false
}

func Unauthorized(ctx *gin.Context, err error) {
	ERR(ctx, http.StatusUnauthorized, int(errUnauthorized), err)
	ctx.Abort()
}

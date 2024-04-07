package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jimu-server/common/resp"
	"github.com/jimu-server/logger"
	"runtime/debug"
	"strings"
)

type argsError struct {
	Msg string `json:"msg"`
}

func (e argsError) Error() string {
	return e.Msg
}

func ArgsErr(msg ...string) argsError {
	return argsError{Msg: strings.Join(msg, " ")}
}

func GlobalException() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				switch e := err.(type) {
				case argsError:
					ctx.JSON(500, resp.Error(e, resp.Msg("参数错误 "), resp.Code(resp.WebArgsErr)))
				default:
					ctx.JSON(500, resp.Error(nil, resp.Msg("系统错误 ", err.(error).Error()), resp.Code(resp.WebErr)))
				}
				logger.Logger.Error(string(debug.Stack()))
			}
		}()
		ctx.Next()
	}
}

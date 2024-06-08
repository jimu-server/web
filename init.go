package web

import (
	"github.com/gin-gonic/gin"
	"github.com/jimu-server/logger"
	"net/http"
)

var Engine = gin.New()

func init() {
	Engine.Use(logger.GinLogger(), GlobalException(), Cors())
}

// Cors 跨域处理
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin) // 可将将 * 替换为指定的域名
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization,Orgid")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
		c.Next()
	}
}

func BindJSON(ctx *gin.Context, bind any) {
	if err := ctx.BindJSON(bind); err != nil {
		panic(ArgsErr(err.Error()))
	}
}

func ShouldJSON(ctx *gin.Context, bind any) {
	if err := ctx.ShouldBind(bind); err != nil {
		panic(ArgsErr(err.Error()))
	}
}
func ShouldBindUri(ctx *gin.Context, bind any) {
	if err := ctx.ShouldBindUri(bind); err != nil {
		panic(ArgsErr(err.Error()))
	}
}

package router

import (
	"context"
	"net/http"
	"seeder-app/config"
	"seeder-app/controller"
	"seeder-app/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func NewRouter(a *controller.App) *gin.Engine {
	r := gin.New()
	cnf, _ := config.NewConfig(context.Background())
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			cnf.AppURL,
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"Origin",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	r.Use(requestid.New()) 
	r.Use(gin.LoggerWithConfig(logger.CustomLogger())) 
	r.Use(gin.Recovery())

	r.POST("/prompt", func(c *gin.Context) {
		a.Prompt(c)
	})
	r.GET("/", hello)
	return r
}

func hello(c *gin.Context) {
	response := gin.H{
		"message": "Hello World!",
	}
	c.JSON(http.StatusOK, response)
}

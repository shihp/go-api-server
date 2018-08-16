package router

import (
	_ "apiserver/docs"
	"apiserver/handler/sd"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	//g.Use(gin.Recovery())
	//g.Use(middleware.NoCache)
	//g.Use(middleware.Options)
	//g.Use(middleware.Secure)
	//g.Use(mw...)

	g.GET("/reload", func(context *gin.Context) {

	})

	g.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
		//log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(context)})
	})

	// swagger api docs
	//g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 404 Handler.
	//g.NoRoute(func(c *gin.Context) {
	//	c.String(http.StatusNotFound, "The incorrect API route.")
	//})

	// Middlewares.
	//pprof.Register(g)

	//g.POST("/login", user.Login)

	//u := g.Group("/v1/user")
	//u.Use(middleware.AuthMiddleware())
	//{
	//	u.POST("", user.Create)       // 创建用户
	//	u.DELETE("/:id", user.Delete) // 删除用户
	//	u.PUT("/:id", user.Update)    // 更新用户
	//	u.GET("", user.List)          // 用户列表
	//	u.GET("/:username", user.Get) // 获取指定用户的详细信息...
	//}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}

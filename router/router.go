package router


import (
	"controller/login"
	"dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Init() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.GET("/login", login.Get)
	router.POST("/login", login.Post)
	router.Use(checkToken)
	api := router.Group("/api", checkToken)
	{
		api.GET("/index", func(context *gin.Context) {
			context.JSON(http.StatusOK, 123)
		})
		api.POST("/cmd", )
	}

	router.Run(":8002")
}

func checkToken(c *gin.Context) {
	return
	token := c.GetHeader("Token")
	// 跳往登录
	if dto.NewToken().GetToken() == "" ||
		(dto.NewToken().GetToken() != "" && dto.NewToken().GetToken() != token) ||
		time.Now().Unix() - dto.NewToken().GetTime().Unix() > 30 * 60 {
		// 单体应用是跳转
		// c.Redirect(http.StatusMovedPermanently, "/login")
		// api是返回重新登录
		c.JSON(http.StatusOK, dto.Resp{Code:-1,Message:"not login"})
		c.Abort()
	}
	// 更新token时间
	dto.NewToken().Set(token, time.Now(), nil)
}

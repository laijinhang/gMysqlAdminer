package cmd

import (
	"conf"
	"dto"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Cmd(c *gin.Context) {
	var sReq dto.SqlReq
	var sResp dto.SqlResp
	if err := c.BindJSON(&sReq);err != nil {
		logrus.Error("Cmd Req Error: ", err)
		sResp.Code = -1
		sResp.Message = err.Error()
		c.JSON(http.StatusOK, sResp)
		return
	}
	conf.DB.Exec("")
}
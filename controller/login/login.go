package login

import (
	"conf"
	"dto"
	"fmt"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"time"
)

func Get(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Post(c *gin.Context) {
	// 1、统一处理请求
	var lReq dto.LoginReq
	var lResp dto.LoginResp
	if err := c.BindJSON(&lReq);err != nil {
		logrus.Errorf("Post login error: ", err)
		lResp.Code = -1
		lResp.Message = err.Error()
		c.JSON(http.StatusOK, &lResp)
		return
	}

	// 2、处理
	if err := login(&lReq);err != nil {
		lResp.Code = -1
		lResp.Message = err.Error()
		c.JSON(http.StatusOK, lResp)
		return
	}
	var err error

	//conf.DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
	//	lReq.Name, lReq.Pwd, lReq.Addr, lReq.DBName))
	conf.DB, err = sql.Open(lReq.DBDriver, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			lReq.Name, lReq.Pwd, lReq.Addr, lReq.DBName))
	res, err := conf.DB.Query("show tables")
	var i string
	for res.Next() {
		res.Scan(&i)
		fmt.Println(i)
	}
	fmt.Println(conf.DB.Exec("show tables"))
	//sql.Open("mysql", "")
	if err != nil {
		lResp.Code = -1
		lResp.Message = err.Error()
		logrus.Error(conf.DB, err)
		c.JSON(http.StatusOK, lResp)
		return
	}

	// 3、响应
	token := fmt.Sprintf("%016v", rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000000000000))
	lResp.Data = map[string]interface{} {
		"token": token,
	}
	dto.NewToken().Set(token, time.Now(), &dto.DBInfo{
		Name:lReq.Name,
		Pwd:lReq.Pwd,
		DBDriver:lReq.DBDriver,
		Addr:lReq.Addr,
		DBName:lReq.DBName,
	})
	c.JSON(http.StatusOK, lResp)
}

func login(lReq *dto.LoginReq) error {
	if lReq.DBDriver != "mysql" {
		return errors.New("DB driver not mysql!")
	}

	return nil
}
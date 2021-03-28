package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shen.com/ginvue/common"
	"shen.com/ginvue/model"
	"shen.com/ginvue/utils"
)

func Register(c *gin.Context) {
	// 引入db
	db := common.GetDB()
	name, _ := c.GetPostForm("name")
	password, _ := c.GetPostForm("password")
	phone, _ := c.GetPostForm("phone")
	// 校验数据的安全性
	log.Println(name, password, phone)
	if len(name) == 0 {
		name = utils.GenerateRand(6)
	}
	if len(password) == 0 {
		password = utils.GenerateRand(10)
	}
	if len(phone) != 11 {
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"message": "手机号不符合要求，请输入11位手机号",
		})
	}
	// 校验数据的安全性
	log.Println(name, password, phone)
	var u = &model.User{}
	db.Where("name = ?", name).First(u)
	// 判断是否查到了此用户信息
	if u.ID != 0 {
		log.Println("此用户信息已存在，不能注册")
		c.JSON(http.StatusUnsupportedMediaType, gin.H{
			"message": "此用户信息已存在，不能注册",
		})
		return
	}
	u.Name = name
	u.Password = password
	u.Phone = phone
	db.Create(u)
	c.JSON(http.StatusUnsupportedMediaType, gin.H{
		"message": "恭喜你注册成功",
	})
}

func Login(c *gin.Context) {
	// 引入db
	db := common.GetDB()
	name, _ := c.GetPostForm("name")
	password, _ := c.GetPostForm("password")
	// 校验数据的安全性
	log.Println(name, password)
	// 查询数据库验证用户名和密码的用户
	var user = new(model.User)
	db.Where("name = ? and password = ?", name, password).First(user)
	if user.ID == 0 {
		log.Println("用户名或密码错误")
		c.JSON(http.StatusOK, gin.H{
			"message": "用户名或密码错误",
		})
		return
	}
	token, err := common.ReleaseToken(user)
	if err != nil {
		log.Println("生成token出错")
	}
	// 登录成功后发放token和cookie
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功，欢迎回来",
		"token":   token,
	})
}

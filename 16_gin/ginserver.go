package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
	"reflect"
	"strings"
	"time"
)

var trans ut.Translator

type LoginForm struct {
	Username string `json:"username" binding:"required,min=3,max=10"`
	Password string `json:"password" binding:"required"`
}

type SignUpForm struct {
	Age        uint8  `json:"age" binding:"gte=18,lte=130"`
	Name       string `json:"name" binding:"required,min=3"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		//注册一个获取jsontag的方法
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT)
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}

		switch locale {
		case "en":
			en_translations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, trans)
		default:
			zh_translations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func timeUseMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		//提前停止函数调用,不再调用后面的函数
		//context.Abort()
		context.Next()
		end := time.Since(t)
		fmt.Println("耗时", end)
		status := context.Writer.Status()
		fmt.Println("状态码", status)

	}
}

func main() {

	if err := InitTrans("zh"); err != nil {
		fmt.Println("初始化编译器错误")
		return
	}
	//使用new 则默认不会开启任何插件
	//r : = gin.New()
	//使用Default会默认开启俩个插件 1. logger 2. recover
	r := gin.Default()
	r.Use(timeUseMiddleware())

	testGroup := r.Group("test")
	{
		testGroup.GET("ping", Pong)
		testGroup.GET("hello", func(context *gin.Context) {
			context.String(http.StatusOK, "hello")
		})
	}

	goodsGroup := r.Group("goods")
	{
		goodsGroup.GET("", goodsList)
		goodsGroup.GET("/:id", goodsDetail)
		goodsGroup.POST("", createGoods)
		goodsGroup.POST("postForm", postForm)
	}

	userGroup := r.Group("user")
	{
		userGroup.POST("/login", loginForm)
		userGroup.POST("/signup", signUpForm)
	}

	r.Run(":8090")
}

func removeTopSruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func signUpForm(context *gin.Context) {
	var signUpForm SignUpForm
	err := context.ShouldBind(&signUpForm)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"error": removeTopSruct(errs.Translate(trans)),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
}

func loginForm(context *gin.Context) {
	var loginForm LoginForm
	err := context.ShouldBind(&loginForm)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
		context.JSON(http.StatusBadRequest, gin.H{
			"error": removeTopSruct(errs.Translate(trans)),
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": "登录成功",
	})
}

func postForm(context *gin.Context) {
	goodsName := context.PostForm("goodsName")
	goodsType := context.DefaultPostForm("goodsType", "1")
	context.JSON(http.StatusOK, gin.H{
		"goodsName": goodsName,
		"goodsType": goodsType,
	})
}

func createGoods(context *gin.Context) {

}

func goodsList(context *gin.Context) {
	goodsName := context.Query("goodsName")
	goodsType := context.DefaultQuery("goodsType", "1")
	context.JSON(http.StatusOK, gin.H{
		"goodsName": goodsName,
		"goodsType": goodsType,
	})
}

func goodsDetail(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func Pong(context *gin.Context) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("error: %v\n", err)
			}
		}()
		var point *LoginForm
		fmt.Println(point.Username)
	}()
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

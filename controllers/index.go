package controllers

import (
	interfaces "authentication/interface"
	"authentication/models"
	"fmt"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService interfaces.UserService
}

func New(userservice interfaces.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}
func (uc *UserController) Signup(ctx *gin.Context) {
	var user models.Credentials
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.Signup(&user)
	if err != nil {
		
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var user models.Credentials
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	 err := uc.UserService.Login(&user,ctx)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "bad credential"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message":"successful" })

}
func (uc *UserController) Home(ctx*gin.Context){
	var us models.Credentials
	ctx.ShouldBindJSON(&us)
	tokenCookie, err := ctx.Request.Cookie("token")
	
	if err != nil {
		ctx.String(http.StatusBadRequest, "Token cookie not found")
		return
	}
	response, err := uc.UserService.Home(tokenCookie.Value)
	if err!=nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"message":"tokenexpired"})
	}
	
    if response!="invalid token"{
		ctx.JSON(http.StatusOK,gin.H{"message":"hello "+response})
	}
}


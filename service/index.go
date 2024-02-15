package services

import (
	interfaces "authentication/interface"
	"authentication/models"
	"errors"
	"fmt"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserServ(usercollection *mongo.Collection, ctx context.Context) interfaces.UserService {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) Signup(user *models.Credentials) error {
	password := user.Password
	//fmt.Println(user.Password)
	var duplicate *models.Credentials
	er := u.usercollection.FindOne(u.ctx, bson.M{"username": user.Username}).Decode(&duplicate)
	fmt.Println(er)
	// if er !=nil{
	// 	return er
	// }
	if duplicate!=nil&&duplicate.Username==user.Username{
		return errors.New("user_name_already_exist")
	}
	if er !=nil{
	hashedPassword, err1 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err1 == nil {

		user.Password = string(hashedPassword)
	}
}

	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err

}

func (u *UserServiceImpl) Login(user *models.Credentials, ctx *gin.Context) error {
	var storeduser *models.Credentials
	var claims *models.Claims
	err := u.usercollection.FindOne(u.ctx, bson.M{"username": user.Username}).Decode(&storeduser)
	fmt.Println(err)
	err1 := bcrypt.CompareHashAndPassword([]byte(storeduser.Password), []byte(user.Password))
	if err1 != nil {

		return err1
	}
	expirationTime := time.Now().UTC().Add(time.Minute * 5)
	claims = &models.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret-key"))
	ctx.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "", false, true)
	
	return err
}

func (u *UserServiceImpl)Home(token string)(string,error){
    claims:=&models.Claims{}
	tkn,err:=jwt.ParseWithClaims(token,claims,func(token *jwt.Token) (interface{}, error) {
		return []byte("secret-key"), nil
	})
	
if err!=nil{
	return  "invalid token",err
}
if !tkn.Valid{
    return  "invalid token",err
}
return claims.Username,nil
}






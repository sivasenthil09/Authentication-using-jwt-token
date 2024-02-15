package interfaces

import (
	"authentication/models"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Signup(*models.Credentials) error
	Login(*models.Credentials,*gin.Context) error
	Home( string)(string,error)
	// Refresh(*http.Cookie)(*http.Cookie,error)
}

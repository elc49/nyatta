package interfaces

import (
	"github.com/3dw1nM0535/nyatta/graph/model"
	jwt "github.com/dgrijalva/jwt-go"
)

type UserService interface {
	CreateUser(user *model.NewUser) (*model.User, error)
	FindById(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	SignIn(user *model.NewUser) (*string, error)
	ValidateToken(token *string) (*jwt.Token, error)
	ServiceName() string
}
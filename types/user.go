package types

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `bson:"_id,omitempty" json:"id,omitempty"`
	Email             string `gorm:"unique" bson:"email" json:"email,omitempty"`
	EncryptedPassword string `bson:"encryptedPassword" json:"_,omitempty"`
	IsAdmin           bool   `bson:"isAdmin" json:"isAdmin,omitempty"`
}

const NBSecretPassword = "A String Very Very Very Strong!!@##$!@#$"

// A Util function to generate jwt_token which can be used in the request header
func GenToken(id string) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwt_token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(NBSecretPassword))
	return token
}

func NewAdminUser(email, password string) (*User, error) {
	user, err := NewUser(email, password)
	if err != nil {
		return nil, err
	}
	user.IsAdmin = true
	return user, nil
}

func NewUser(email, password string) (*User, error) {
	epw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Email:             email,
		EncryptedPassword: string(epw),
	}, nil
}

func (u *User) ValidatePassword(pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(pw))
	return err == nil
}

type UserLoginResponse struct {
	Email string `json:"email,omitempty"`
	Token string `json:"token,omitempty"`
}

func (u *User) ToLoginResponse() *UserLoginResponse {
	return &UserLoginResponse{
		Email: u.Email,
		Token: GenToken(strconv.Itoa(u.ID)),
	}
}

type UserResponse struct {
	Email string
	ID    int
}

func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		Email: u.Email,
		ID:    u.ID,
	}
}

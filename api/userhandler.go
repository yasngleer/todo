package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yasngleer/todo/store"
	"github.com/yasngleer/todo/types"
)

type UserHandler struct {
	store *store.GormUserStore
}

func NewUserHandler(store *store.GormUserStore) *UserHandler {
	return &UserHandler{store: store}
}

type UserRegisterRequest struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *UserHandler) UserRegister(c *gin.Context) {
	uregreq := &UserRegisterRequest{}
	err := c.BindJSON(uregreq)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	user, _ := types.NewUser(uregreq.Email, uregreq.Password)
	err = h.store.Insert(user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(200, user.ToLoginResponse())
}
func (u *UserHandler) UserLogin(c *gin.Context) {
	uregreq := &UserRegisterRequest{}
	err := c.BindJSON(uregreq)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	usr, _ := u.store.GetByMail(uregreq.Email)
	if usr.ValidatePassword(uregreq.Password) {
		c.JSON(200, usr.ToLoginResponse())
		return
	}
	c.String(200, "zort")

}

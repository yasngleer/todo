package store

import (
	"github.com/yasngleer/todo/types"
)

type UserStore interface {
	Insert(user *types.User) error
	GetById(id int) (*types.User, error)
	GetByMail(mail string) (*types.User, error)
}

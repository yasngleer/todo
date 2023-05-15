package store

import (
	"fmt"

	"github.com/yasngleer/todo/types"
	"gorm.io/gorm"
)

type GormUserStore struct {
	db *gorm.DB
}

func NewGormUserStore(db *gorm.DB) *GormUserStore {
	db.AutoMigrate(types.User{})
	return &GormUserStore{db: db}
}

//s *GormUserStore
func (s *GormUserStore) Insert(user *types.User) error {
	err := s.db.Create(user).Error
	fmt.Println(err)
	return err
}

func (s *GormUserStore) GetById(id int) (*types.User, error) {
	user := &types.User{}
	err := s.db.First(user, types.User{ID: id}).Error
	return user, err
}

func (s *GormUserStore) GetByMail(mail string) (*types.User, error) {
	user := &types.User{}
	err := s.db.First(user, types.User{Email: mail}).Error
	return user, err
}

package store

import (
	"github.com/yasngleer/todo/types"
	"gorm.io/gorm"
)

type GormTodoStore struct {
	db *gorm.DB
}

func NewGormTodoStore(db *gorm.DB) *GormTodoStore {
	db.AutoMigrate(types.TodoList{})
	db.AutoMigrate(types.TodoStep{})

	return &GormTodoStore{db: db}
}

func (s *GormTodoStore) InsertTodo(todo *types.TodoList) error {
	err := s.db.Create(todo).Error
	return err
}

func (s *GormTodoStore) InsertStep(todostep *types.TodoStep) error {
	err := s.db.Create(todostep).Error
	return err
}

func (s *GormTodoStore) GetTodoByID(id int, showdeleted bool) (types.TodoList, error) {
	todo := types.TodoList{}
	if showdeleted {
		err := s.db.Where(&types.TodoList{ID: id}).
			Preload("Steps").
			Preload("User").Find(&todo).Error
		return todo, err

	} else {
		err := s.db.Where(&types.TodoList{ID: id}).
			Preload("Steps", "deleted_at is NULL").
			Preload("User").Find(&todo).Error
		return todo, err
	}
}
func (s *GormTodoStore) GetTodosByUserid(userid int, showdeleted bool) ([]types.TodoList, error) {
	todos := []types.TodoList{}
	if showdeleted {
		err := s.db.Where(&types.TodoList{UserID: userid}).Preload("Steps").Preload("User").Find(&todos).Error
		return todos, err
	} else {
		err := s.db.Where(&types.TodoList{UserID: userid}).Preload("Steps", "deleted_at is NULL").Preload("User").Find(&todos).Error
		return todos, err
	}
}
func (s *GormTodoStore) GetStepByID(id int, showdeleted bool) (types.TodoStep, error) {
	step := types.TodoStep{}
	if showdeleted {
		err := s.db.Where(&types.TodoStep{ID: id}).Find(&step).Error
		return step, err
	} else {
		err := s.db.Where(&types.TodoStep{ID: id, DeletedAt: nil}).Find(&step).Error
		return step, err
	}
}
func (s *GormTodoStore) UpdateTodo(todo *types.TodoList) error {
	err := s.db.Save(todo).Error
	return err
}
func (s *GormTodoStore) UpdateStep(step *types.TodoStep) error {
	err := s.db.Save(step).Error
	return err
}

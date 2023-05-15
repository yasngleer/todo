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
func (s *GormTodoStore) GetTodoPercentageByID(id int) (int64, error) {
	var all int64 = 0
	var done int64 = 0
	s.db.Model(&types.TodoStep{}).Where(&types.TodoStep{TodoListID: id}).Count(&all)
	s.db.Model(&types.TodoStep{}).Where(&types.TodoStep{TodoListID: id, Completed: true}).Count(&done)
	percentage := float64(done) / float64(all)

	return int64(percentage * 100), nil
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

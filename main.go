package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yasngleer/todo/api"
	"github.com/yasngleer/todo/store"
	"github.com/yasngleer/todo/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api")
	db, _ := gorm.Open(sqlite.Open("test_file.db"), &gorm.Config{})
	userstore := store.NewGormUserStore(db)

	userhandler := api.NewUserHandler(userstore)
	v1.POST("/users", userhandler.UserRegister)
	v1.POST("/users/login", userhandler.UserLogin)

	todostore := store.NewGormTodoStore(db)
	todohandler := api.NewTodoHandler(userstore, todostore)
	v1.Use(api.AuthMiddleware(true))
	v1.POST("/todo", todohandler.Create)
	v1.GET("/todo/:id", todohandler.GetOne)
	v1.GET("/todo/", todohandler.GetAll)
	v1.PUT("/todo/:todoid", todohandler.UpdateTodo)
	v1.DELETE("/todo/:todoid", todohandler.DeleteTodo)
	v1.POST("/todo/:id/step", todohandler.AddStep)
	v1.PUT("/todo/:todoid/step/:stepid", todohandler.UpdateStep)
	v1.DELETE("/todo/:todoid/step/:stepid", todohandler.DeleteStep)

	//create admin user. Not a good idea ,demo only
	ausr, _ := types.NewAdminUser("admin@admin.com", "admin")
	userstore.Insert(ausr)

	r.Run()
}

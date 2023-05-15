package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yasngleer/todo/store"
	"github.com/yasngleer/todo/types"
)

type TodoListCreateRequest struct {
	Name string
}

type TodoHandler struct {
	userStore store.UserStore
	todoStore *store.GormTodoStore
}

func NewTodoHandler(us store.UserStore, ts *store.GormTodoStore) *TodoHandler {
	return &TodoHandler{
		userStore: us,
		todoStore: ts,
	}
}

func (h *TodoHandler) Create(c *gin.Context) {
	value, exists := c.Get("user_id")
	if !exists {
		c.Status(401)
		return
	}
	user, _ := h.userStore.GetById(value.(int))
	req := &TodoListCreateRequest{}
	err := c.BindJSON(req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	todo := &types.TodoList{
		Name:   req.Name,
		UserID: user.ID,
	}
	h.todoStore.InsertTodo(todo)
	c.JSON(200, todo)
}

func (h *TodoHandler) DeleteStep(c *gin.Context) {

	userid, exists := c.Get("user_id")
	if !exists {
		c.Status(401)
		return
	}
	stepid := c.Param("stepid")
	stepidint, _ := strconv.Atoi(stepid)

	user, _ := h.userStore.GetById(userid.(int))
	step, _ := h.todoStore.GetStepByID(stepidint, user.IsAdmin)
	todo, _ := h.todoStore.GetTodoByID(step.TodoListID, user.IsAdmin)
	if !user.IsAdmin && user.ID != todo.UserID {
		c.Status(401)
		return
	}
	currenttime := time.Now()
	step.DeletedAt = &currenttime
	h.todoStore.UpdateStep(&step)
	c.JSON(200, todo.ToResponse())
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {

	userid, exists := c.Get("user_id")
	if !exists {
		c.Status(401)
		return
	}
	todoid, _ := strconv.Atoi(c.Param("todoid"))

	user, _ := h.userStore.GetById(userid.(int))
	todo, _ := h.todoStore.GetTodoByID(todoid, user.IsAdmin)
	if !user.IsAdmin && user.ID != todo.UserID {
		c.Status(401)
		return
	}
	currenttime := time.Now()
	todo.DeletedAt = &currenttime
	h.todoStore.UpdateTodo(&todo)
	c.JSON(200, todo.ToResponse())
}

func (h *TodoHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	value, exists := c.Get("user_id")
	if !exists {
		c.Status(401)
		return
	}
	user, _ := h.userStore.GetById(value.(int))
	idint, _ := strconv.Atoi(id)
	todo, _ := h.todoStore.GetTodoByID(idint, user.IsAdmin)
	percentage, _ := h.todoStore.GetTodoPercentageByID(idint)
	todo.PercentOfCompletion = int(percentage)
	if !user.IsAdmin && user.ID != todo.UserID {
		c.Status(401)
		return
	}
	c.JSON(200, todo.ToResponse())
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	userid := c.Query("userid")
	value, exists := c.Get("user_id")
	if !exists {
		c.Status(401)
		return
	}
	user, _ := h.userStore.GetById(value.(int))
	// set userid to users id if it is not given
	var useridint int
	if userid == "" {
		useridint = user.ID
	} else {
		useridint, _ = strconv.Atoi(userid)
	}
	// if the user is not admin it should match the userid param
	if !user.IsAdmin && useridint != user.ID {
		c.Status(401)
		return
	}
	var todos []types.TodoList

	todos, _ = h.todoStore.GetTodosByUserid(useridint, user.IsAdmin)

	c.JSON(200, types.TodoListsToResponse(todos))
}

type StepCreateRequest struct {
	Context string
}

func (h *TodoHandler) AddStep(c *gin.Context) {
	req := StepCreateRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	todoid := c.Param("id")
	value, exists := c.Get("user_id")
	if !exists {
		c.Status(401)
		return
	}
	user, _ := h.userStore.GetById(value.(int))
	todoidint, _ := strconv.Atoi(todoid)
	todo, _ := h.todoStore.GetTodoByID(todoidint, user.IsAdmin)
	if !user.IsAdmin && user.ID != todo.UserID {
		c.Status(401)
		return
	}
	step := types.TodoStep{
		Context:    req.Context,
		TodoListID: todoidint,
	}
	h.todoStore.InsertStep(&step)
	c.JSON(200, step)
}

type StepUpdateRequest struct {
	Context   *string
	Completed *bool
}

func (h *TodoHandler) UpdateStep(c *gin.Context) {
	req := StepUpdateRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	stepid, _ := strconv.Atoi(c.Param("stepid"))

	//check auth
	value, exists := c.Get("user_id")
	if !exists {
		c.Status(401)
		return
	}
	user, _ := h.userStore.GetById(value.(int))

	step, _ := h.todoStore.GetStepByID(stepid, user.IsAdmin)
	todo, _ := h.todoStore.GetTodoByID(step.TodoListID, user.IsAdmin)
	if !user.IsAdmin && user.ID != todo.UserID {
		c.Status(401)
		return
	}

	if req.Context != nil {
		step.Context = *req.Context
	}

	if req.Completed != nil {
		step.Completed = *req.Completed
	}
	h.todoStore.UpdateStep(&step)
	c.JSON(200, step)
}

type TodoUpdateRequest struct {
	Name *string
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	req := TodoUpdateRequest{}
	err := c.BindJSON(&req)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	todoid, _ := strconv.Atoi(c.Param("todoid"))

	//check auth
	value, exists := c.Get("user_id")
	if !exists {
		c.Status(401)
		return
	}
	user, _ := h.userStore.GetById(value.(int))

	todo, _ := h.todoStore.GetTodoByID(todoid, user.IsAdmin)
	if !user.IsAdmin && user.ID != todo.UserID {
		c.Status(401)
		return
	}
	if req.Name != nil {
		todo.Name = *req.Name

	}

	h.todoStore.UpdateTodo(&todo)
	c.JSON(200, todo)
}

package controllers

import (
	"github.com/balkin/todolist/todo"
	"github.com/gin-gonic/gin"
)

// TodoCountItems godoc
// @Summary Count todo items without a parent
// @Produce  json
// @Success 200 {object} controllers.CountStruct
// @Router /todo/count [get]
func TodoCountItems(ctx *gin.Context) {
	if c, err := todo.CountRootTodoItems(); err == nil {
		ctx.JSON(200, CountStruct{Count: c})
	} else {
		ctx.Error(err)
	}
}

// TodoListItems godoc
// @Summary List todo items
// @Produce  json
// @Success 200 {array} todo.TodoItem
// @Router /todo/item [get]
func TodoListItems(ctx *gin.Context) {
	todo_items := []todo.TodoItem{}
	ctx.JSON(200, todo_items)
}

// TodoShowItem godoc
// @Summary Show todo item including subitems
// @Produce  json
// @Param id path int true "TodoItem ID"
// @Success 200 {array} todo.TodoItem
// @Router /todo/item/{id} [get]
func TodoShowItem(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"id": ctx.Param("id")})
}

// TodoAddItem godoc
// @Summary Add todo item (root)
// @Param item body controllers.SimpleTodoItem true "Add Todo item"
// @Produce  json
// @Success 200 {array} todo.TodoItem
// @Router /todo/item/ [post]
func TodoAddItem(ctx *gin.Context) {
	ctx.JSON(200, nil)
}

// TodoAddSubItem godoc
// @Summary Add todo sub item
// @Produce  json
// @Param id path int true "TodoItem ID"
// @Success 200 {array} todo.TodoItem
// @Router /todo/item/{id} [post]
func TodoAddSubItem(ctx *gin.Context) {
	ctx.JSON(200, nil)
}

// TodoDeleteItem godoc
// @Summary Delete todo item including subitems
// @Param id path int true "TodoItem ID"
// @Success 200
// @Router /todo/item/{id} [delete]
func TodoDeleteItem(ctx *gin.Context) {
	ctx.JSON(200, nil)
}

package controllers

import (
	"github.com/balkin/todolist/todo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: "Failed to count todo items"})
	}
}

// TodoCountItems godoc
// @Summary Count todo items without a parent
// @Produce  json
// @Success 200 {object} controllers.CountStruct
// @Router /todo/countall [get]
func TodoCountAllItems(ctx *gin.Context) {
	if c, err := todo.CountTodoItems(); err == nil {
		ctx.JSON(200, CountStruct{Count: c})
	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: "Failed to count todo items"})
	}
}

// TodoListItems godoc
// @Summary List todo items
// @Produce  json
// @Success 200 {array} todo.TodoItem
// @Router /todo/item/ [get]
func TodoListItems(ctx *gin.Context) {
	if todo_items, err := todo.ListTodoItems(); err == nil {
		ctx.JSON(200, todo_items)
	} else {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: "Failed to list root todo items"})
	}
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
// @Param item body todo.SimpleTodoItem true "Add TodoItem"
// @Produce  json
// @Success 200 {object} todo.TodoItem
// @Router /todo/item/ [post]
func TodoAddItem(ctx *gin.Context) {
	var item todo.SimpleTodoItem
	if err := ctx.BindJSON(&item); err == nil {
		var res *todo.TodoItem
		if res, err = todo.AddTodoItem(item.Name); err == nil {
			ctx.JSON(http.StatusOK, res)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: "Failed to add todo item"})
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorStruct{Error: "Failed to bind JSON fields"})
	}
}

// TodoAddSubItem godoc
// @Summary Add todo sub item
// @Produce  json
// @Param id path int true "TodoItem ID"
// @Param item body todo.SimpleTodoItem true "Add sub TodoItem"
// @Success 200 {object} todo.TodoItem
// @Router /todo/item/{id} [post]
func TodoAddSubItem(ctx *gin.Context) {
	var item todo.SimpleTodoItem
	if err := ctx.BindJSON(&item); err == nil {
		var res *todo.TodoItem
		if parent_id, err := strconv.Atoi(ctx.Param("id")); err == nil {
			if res, err = todo.AddTodoSubitem(parent_id, item.Name); err == nil {
				ctx.JSON(http.StatusOK, res)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: "Failed to add todo sub item"})
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorStruct{Error: "Failed to find parent todo item"})
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorStruct{Error: "Failed to bind JSON fields"})
	}
}

// TodoDeleteItem godoc
// @Summary Delete todo item including subitems
// @Param id path int true "TodoItem ID"
// @Success 200
// @Router /todo/item/{id} [delete]
func TodoDeleteItem(ctx *gin.Context) {
	ctx.JSON(200, nil)
}

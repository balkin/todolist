package controllers

import (
	"github.com/balkin/todolist/db"
	"github.com/balkin/todolist/todo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// TodoCountItems godoc
// @Summary Count todo items without a parent
// @Produce  json
// @Success 200 {object} controllers.CountStruct
// @Router /todo/count [get]
func TodoCountItems(ctx *gin.Context) {
	if c, err := db.CountRootTodoItems(); err == nil {
		ctx.JSON(http.StatusOK, CountStruct{Count: c})
	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: FailedToCountTodoItems})
	}
}

// TodoCountItems godoc
// @Summary Count todo items without a parent
// @Produce  json
// @Success 200 {object} controllers.CountStruct
// @Router /todo/countall [get]
func TodoCountAllItems(ctx *gin.Context) {
	if c, err := db.CountTodoItems(); err == nil {
		ctx.JSON(http.StatusOK, CountStruct{Count: c})
	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: FailedToCountTodoItems})
	}
}

// TodoListItems godoc
// @Summary List todo items
// @Produce  json
// @Success 200 {array} todo.TodoItem
// @Router /todo/item/ [get]
func TodoListItems(ctx *gin.Context) {
	if todo_items, err := db.ListTodoItems(); err == nil {
		ctx.JSON(http.StatusOK, todo_items)
	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: FailedToListRootTodoItems})
	}
}

// TodoShowItem godoc
// @Summary Show todo item including subitems
// @Produce  json
// @Param id path int true "TodoItem ID"
// @Success 200 {array} todo.TodoItem
// @Router /todo/item/{id} [get]
func TodoShowItem(ctx *gin.Context) {
	if id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		if items, err := db.ShowTodoItem(id); err == nil {
			ctx.JSON(http.StatusOK, items)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: FailedToGetOrSerializeTodoItem})
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusNotFound, ErrorStruct{Error: FailedToFindTodoItem})
	}
}

// TodoAddItem godoc
// @Summary Add todo item (root)
// @Param item body todo.SimpleTodoItem true "Add TodoItem"
// @Produce  json
// @Success 200 {object} todo.TodoItem
// @Router /todo/item/ [post]
func TodoAddItem(ctx *gin.Context) {
	item := new(todo.SimpleTodoItem)
	if err := ctx.BindJSON(item); err == nil {
		if res, err := db.AddTodoItem(item.Name); err == nil {
			ctx.JSON(http.StatusOK, res)
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: FailedToAddTodoItem})
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorStruct{Error: FailedToBindJsonFields})
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
		if parent_id, err := strconv.Atoi(ctx.Param("id")); err == nil {
			if res, err := db.AddTodoSubitem(parent_id, item.Name); err == nil {
				ctx.JSON(http.StatusOK, res)
			} else {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: FailedToAddTodoSubitem})
			}
		} else {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorStruct{Error: FailedToFindParentTodoItem})
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ErrorStruct{Error: FailedToBindJsonFields})
	}
}

// TodoDeleteItem godoc
// @Summary Delete todo item including subitems
// @Param id path int true "TodoItem ID"
// @Success 200
// @Router /todo/item/{id} [delete]
func TodoDeleteItem(ctx *gin.Context) {
	if id, err := strconv.Atoi(ctx.Param("id")); err == nil {
		if err := db.DeleteTodoItem(id); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"ok": true})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, ErrorStruct{Error: FailedToDeleteTodoItem})
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusNotFound, ErrorStruct{Error: FailedToFindTodoItem})
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"github.com/balkin/todolist/controllers"
	"github.com/balkin/todolist/db"
	"github.com/balkin/todolist/todo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const todoItemName = "SimpleTodo"

func TestListTodoWorks(t *testing.T) {
	router := SetupTestRouter()
	db.ConnectToTestDatabase()
	w := MakeGinRequest(router, "GET", "/api/v1/todo/item/")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")

	var items []todo.TodoItem
	if err := json.Unmarshal(w.Body.Bytes(), &items); err != nil {
		assert.Error(t, err)
	}
}

func TestCountTodoWorks(t *testing.T) {
	router := SetupTestRouter()
	db.ConnectToTestDatabase()
	w := MakeGinRequest(router, "GET", "/api/v1/todo/count")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")

	var countStruct controllers.CountStruct
	if err := json.Unmarshal(w.Body.Bytes(), &countStruct); err != nil {
		assert.Error(t, err)
	}
}

func TestCountAllTodoWorks(t *testing.T) {
	router := SetupTestRouter()
	db.ConnectToTestDatabase()
	w := MakeGinRequest(router, "GET", "/api/v1/todo/countall")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")

	var countStruct controllers.CountStruct
	if err := json.Unmarshal(w.Body.Bytes(), &countStruct); err != nil {
		assert.Error(t, err)
	}
}

func TestPostNewTodo(t *testing.T) {
	router := SetupTestRouter()
	db.ConnectToTestDatabase()
	if jsonBytes, err := json.Marshal(todo.SimpleTodoItem{Name: todoItemName}); err == nil {
		w := MakeGinReaderRequest(router, "POST", "/api/v1/todo/item/", bytes.NewReader(jsonBytes))
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Header().Get("Content-Type"), "application/json")
		var item todo.TodoItem
		if err := json.Unmarshal(w.Body.Bytes(), &item); err != nil {
			assert.Error(t, err)
		}
		// Should return same name
		assert.Equal(t, todoItemName, item.Name)
		// No parent
		assert.Equal(t, 0, item.ParentId)
		// Should have real id
		assert.True(t, item.Id > 0)
	} else {
		assert.Error(t, err)
	}
}

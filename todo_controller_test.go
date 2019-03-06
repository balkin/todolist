package main

import (
	"encoding/json"
	"github.com/balkin/todolist/controllers"
	"github.com/balkin/todolist/todo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestListTodoWorks(t *testing.T) {
	router := SetupTestRouter()
	todo.ConnectToTestDatabase()
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
	todo.ConnectToTestDatabase()
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
	todo.ConnectToTestDatabase()
	w := MakeGinRequest(router, "GET", "/api/v1/todo/countall")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Header().Get("Content-Type"), "application/json")

	var countStruct controllers.CountStruct
	if err := json.Unmarshal(w.Body.Bytes(), &countStruct); err != nil {
		assert.Error(t, err)
	}
}

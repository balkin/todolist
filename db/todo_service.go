package db

import (
	. "github.com/balkin/todolist/todo"
)

func AddTodoItem(name string) (*TodoItem, error) {
	item := TodoItem{Name: name}
	err := Db.Insert(&item)
	return &item, err
}

func ListTodoItems() ([]TodoItem, error) {
	var todoItems = []TodoItem{}
	_, err := Db.Query(&todoItems, "SELECT * FROM todo_items WHERE parent_id IS NULL")
	return todoItems, err
}

func CountTodoItems() (int, error) {
	return Db.Model((*TodoItem)(nil)).Count()
}

func CountRootTodoItems() (int, error) {
	return Db.Model((*TodoItem)(nil)).Where("parent_id IS NULL").Count()
}

func ShowTodoItem(id int) ([]TodoItem, error) {
	var todoItems = []TodoItem{}
	_, err := Db.Query(&todoItems, `
WITH RECURSIVE r AS (
  SELECT id, parent_id, name FROM todo_items WHERE id = ?
  UNION
  SELECT todo_items.id, todo_items.parent_id, todo_items.name FROM todo_items JOIN r ON todo_items.parent_id = r.id
)
SELECT * FROM r;
`, id)
	return todoItems, err
}

func AddTodoSubitem(id int, name string) (*TodoItem, error) {
	parent := TodoItem{Id: id}
	if err := Db.Select(&parent); err == nil {
		item := TodoItem{Name: name, ParentId: id}
		err := Db.Insert(&item)
		return &item, err
	} else {
		return nil, err
	}
}

func DeleteTodoItem(id int) error {
	return Db.Delete(&TodoItem{Id: id})
}

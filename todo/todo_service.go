package todo

import (
	"github.com/go-pg/pg"
)

var Db *pg.DB

func ConnectToDatabase() {
	Db = pg.Connect(&pg.Options{User: "root", Password: "Passw0rd"})
}

func DisconnectDatabase() {
	if Db != nil {
		_ = Db.Close()
		Db = nil
	}
}

func AddTodoItem(name string) error {
	return Db.Insert(TodoItem{Name: name})
}

func ListTodoItems() ([]TodoItem, error) {
	var todos []TodoItem
	_, err := Db.Query(&todos, "SELECT * FROM todos WHERE parent_id IS NULL")
	return todos, err
}

func CountTodoItems() (int, error) {
	return Db.Model((*TodoItem)(nil)).Count()
}

func CountRootTodoItems() (int, error) {
	return Db.Model((*TodoItem)(nil)).Where("parent_id IS NULL").Count()
}

func ShowTodoItem(id int) ([]TodoItem, error) {
	var todos []TodoItem
	_, err := Db.Query(&todos, `
WITH RECURSIVE r AS (
  SELECT id, parent_id, name FROM todos WHERE parent_id = ?
  UNION
  SELECT todos.id, todos.parent_id, todos.name FROM todos JOIN r ON todos.parent_id = r.id
)
SELECT * FROM r;
`, id)
	return todos, err
}

func AddTodoSubitem(id int, name string) error {
	item := TodoItem{Name: name, ParentId: id}
	return Db.Insert(item)
}

func DeleteTodoItem(id int) error {
	return Db.Delete(TodoItem{Id: id})
}

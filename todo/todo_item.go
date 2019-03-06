package todo

type SimpleTodoItem struct {
	Name string
}

type TodoItem struct {
	Id       int
	Name     string
	ParentId int       `sql:"on_delete:CASCADE"`
	Parent   *TodoItem `json:"-",sql:"on_delete:CASCADE",pg:"fk:parent_id"`
}

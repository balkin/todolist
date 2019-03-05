package todo

type SimpleTodoItem struct {
	Name string
}

type TodoItem struct {
	Id       int
	Name     string
	ParentId int
}

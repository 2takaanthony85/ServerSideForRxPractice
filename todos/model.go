package todos

//Todo todo item model
type Todo struct {
	ID       int    `json:"todo-id"`
	UUID     string `json:"todo-uuid"`
	Title    string `json:"todo-title"`
	Contents string `json:"todo-contents"`
}

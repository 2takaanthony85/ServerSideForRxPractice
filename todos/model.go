package todos

//Todo todo item model
type Todo struct {
	ID       string `json:"todo-id"`
	Title    string `json:"todo-title"`
	Contents string `json:"todo-contents"`
}

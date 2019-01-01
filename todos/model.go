package todos

type Todo struct {
	ID       int    `json:"todo-id"`
	Title    string `json:"todo-title"`
	Contents string `json:"todo-contents"`
}

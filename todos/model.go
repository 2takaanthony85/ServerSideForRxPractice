package todos

//Todos todo item model
type Todos struct {
	ID       int    `json:"todo-id"`
	UID      string `json:"todo-uid"`
	Title    string `json:"todo-title"`
	Contents string `json:"todo-contents"`
}

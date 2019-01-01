package users

//User user info
type User struct {
	ID       int    `json:"user-id"`
	UUID     string `json:"user-uuid"`
	Name     string `json:"user-name"`
	Email    string `json:"user-email"`
	Password string `json:"user-password"`
}

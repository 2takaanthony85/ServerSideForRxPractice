package users

//User user info
type User struct {
	ID       string `json:"user-id"`
	Name     string `json:"user-name"`
	Email    string `json:"user-email"`
	Password string `json:"user-password"`
}

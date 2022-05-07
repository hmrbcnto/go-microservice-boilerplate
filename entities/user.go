package entities

type User struct {
	ID       string `json:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

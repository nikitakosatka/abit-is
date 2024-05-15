package schema

type User struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}

type UID struct {
	ID string `json:"id"`
}

package api

type RegisterRequest struct {
	Username  string `json:"username" db:"username" binding:"required"`
	FirstName string `json:"first_name" db:"first_name" binding:"required"`
	LastName  string `json:"last_name" db:"last_name" binding:"required"`
	Password  string `json:"password" db:"hashed_password" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

package dto

type CreateUserDto struct {
	Name     string `json:"name" binding:"required" validate:"required"`
	Email    string `json:"email" binding:"required" validate:"required,email"`
	Password string `json:"password" binding:"required" validate:"required,gte=8"`
	Active   bool   `json:"active"`
	Image    string `json:"image"`
}

type UpdateUserDto struct {
	Name     string `json:"name"`
	Active   bool   `json:"active"`
	Image    string `json:"image"`
	Password string `json:"password" validate:"omitempty,gte=8"`
}

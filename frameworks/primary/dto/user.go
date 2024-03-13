package dto

type CreateUserDto struct {
	Name   string `json:"name,omitempty" binding:"required" validate:"required"`
	Age    int    `json:"age,omitempty" binding:"required" validate:"required,gte=18"`
	Active bool   `json:"active,omitempty" binding:"required" validate:"required"`
}

type UpdateUserDto struct {
	Name   string `json:"name,omitempty"`
	Age    int    `json:"age,omitempty" validate:"omitempty,gte=18"`
	Active bool   `json:"active,omitempty"`
}

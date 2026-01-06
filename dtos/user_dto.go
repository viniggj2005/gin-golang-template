package dtos

type UserDto struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*"`
	Age      int    `json:"age" binding:"required,gte=18"`
}

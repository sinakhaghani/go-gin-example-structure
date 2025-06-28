package validations

type CreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,unique=users.email"`
	Password string `json:"password" binding:"required,min=6"`
}

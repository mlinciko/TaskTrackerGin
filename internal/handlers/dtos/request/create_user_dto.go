package requests

type CreateUserRequestDto struct {
	FirstName string `json:"firstName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

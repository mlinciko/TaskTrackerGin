package requests

type PutUserRequestDto struct {
	FirstName string `json:"firstName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

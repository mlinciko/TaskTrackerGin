package requests

type CreateUserRequestDto struct {
	FirstName string `json:"firstName" binding:"required"`
}

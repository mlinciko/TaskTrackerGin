package requests

type PutUserRequestDto struct {
	FirstName string `json:"firstName" binding:"required"`
}

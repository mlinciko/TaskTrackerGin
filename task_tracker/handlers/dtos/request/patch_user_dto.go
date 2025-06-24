package requests

type PatchUserRequestDto struct {
	FirstName string `json:"firstName"`
	Email     string `json:"email" binding:"email"`
}

package responses

type Response struct {
	Data         any    `json:"data"`
	Code         int    `json:"code"`
	ErrorMessage string `json:"errorMessage"`
}

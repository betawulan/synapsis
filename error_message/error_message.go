package error_message

type Unauthorized struct {
	Message string `json:"message" example:"unauthorized"`
}

func (e Unauthorized) Error() string {
	return e.Message
}

package events

type UserCreatedEvent struct {
	ID    string `json:"id"`
	Email string `json:"email""`
	Token string `json:"token"`
	Host  string `json:"host"`
}

type PasswordReset struct {
	ID    string `json:"id"`
	Email string `json:"email""`
	Token string `json:"token"`
	Host  string `json:"host"`
}

func (e *UserCreatedEvent) EventName() string {
	return "user.created"
}

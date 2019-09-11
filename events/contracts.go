package events

type UserCreatedEvent struct {
	ID    string `json:"id"`
	Email string `json:"email""`
	Token string `json:"token"`
}

func (e *UserCreatedEvent) EventName() string {
	return "user.created"
}

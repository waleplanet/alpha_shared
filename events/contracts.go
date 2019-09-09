package events

type UserCreatedEvent struct {
	ID    string `json:"id"`
	Email string `json:email`
}

func (e *UserCreatedEvent) EventName() string {
	return "user.created"
}

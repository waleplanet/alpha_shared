package events

type UserCreatedEvent struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Host  string `json:"host"`
}
type WelcomeUserEvent struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Host     string `json:"host"`
	Username string `json:"username"`
}

type PasswordReset struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Host  string `json:"host"`
}

type OTPCreated struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
	Host  string `json:"host"`
}

func (e *UserCreatedEvent) EventName() string {
	return "user.created"
}
func (e *PasswordReset) EventName() string {
	return "user.reset_password"
}

func (e *OTPCreated) EventName() string {
	return "otp.created"
}

func (e *WelcomeUserEvent) EventName() string {
	return "user.welcome"
}

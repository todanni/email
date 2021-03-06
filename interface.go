package email

type Recipient struct {
	Email    string
	FullName string
}

type Service interface {
	// SendPasswordResetEmail -
	SendPasswordResetEmail(recipient Recipient) error

	// SendVerificationEmail -
	SendVerificationEmail(code string, recipient Recipient) error
}

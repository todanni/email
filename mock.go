package email

import "github.com/stretchr/testify/mock"

type MockEmail struct {
	mock.Mock
}

func (e *MockEmail) SendPasswordResetEmail(code string, recipient Recipient) error {
	args := e.Called(code, recipient)
	return args.Error(0)
}

func (e *MockEmail) SendVerificationEmail(code string, recipient Recipient) error {
	args := e.Called(code, recipient)
	return args.Error(0)
}

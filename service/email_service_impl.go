package service

import (
	"fmt"
)

type emailServiceImpl struct{}

func (e *emailServiceImpl) SendEmail(to string, subject string, body string) error {
	fmt.Println("=== EMAIL SENT ===")
	fmt.Println("To:", to)
	fmt.Println("Subject:", subject)
	fmt.Println("Body:", body)
	fmt.Println("==================")

	return nil
}

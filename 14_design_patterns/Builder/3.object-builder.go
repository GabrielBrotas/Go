package main

import "strings"

type email struct {
	from, to, subject, body string
}

type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	// validations
	if !strings.Contains(from, "@")  {
		panic("email should container @")
	}

	b.email.from = from
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	// validations ...
	b.email.to = to
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	// validations ...
	b.email.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	// validations ...
	b.email.body = body
	return b
}

func SendMailImpl(b *EmailBuilder) {
	
}

func main() {
	my_email := EmailBuilder{}
	my_email.
		From("me@gmail.com").
		To("him@gmail.com").
		Subject("news about builder pattern").
		Body("Very useful")

	SendMailImpl(&my_email)
}
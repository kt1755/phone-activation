package tiki

import "time"

type Usage struct {
	ActivationDate   time.Time
	DeactivationDate time.Time
}

type PhoneUsage struct {
	Phone    string
	Duration Usage
}

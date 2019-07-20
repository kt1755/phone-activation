package tiki

import "time"

type Usage struct {
	ActivationDate *time.Time
	Deactivation   *time.Time
}

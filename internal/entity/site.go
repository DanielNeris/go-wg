package entity

import "time"

type Site struct {
	URL     string
	Status  int
	Latency time.Duration
	Err     error
}

func (s Site) OK() bool {
	return s.Err == nil && s.Status >= 200 && s.Status < 400
}

package echotest

import "testing"

type Echotest struct {
	t *testing.T
}

func New(t *testing.T) *Echotest {
	return &Echotest{t: t}
}

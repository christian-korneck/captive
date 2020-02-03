package main

import (
	"testing"
)

func TestIscaptive(t *testing.T) {
	c, err := Iscaptive()
	if c != true || err != nil {
		t.Errorf("captive status offline or errored with: '%s' - are you really online?", err)
	}
}

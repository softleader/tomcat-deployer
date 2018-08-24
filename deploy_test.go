package main

import "testing"

func TestStop(t *testing.T) {
	err := stop()
	if err != nil {
		t.Error(err)
	}
}

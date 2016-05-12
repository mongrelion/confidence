package main

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	v := GetVersion()
	if v != "1.0" {
		t.Errorf("expected 1.0; got %s", v)
	}
}

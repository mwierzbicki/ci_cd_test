package main

import (
	"strings"
	"testing"
)

func TestWelcomeString(t *testing.T) {
	output := welcomeString()
	if !strings.Contains(output, "Hello Go!") {
		t.Fatalf("welcomeString() does not output a greeting")
	}
}

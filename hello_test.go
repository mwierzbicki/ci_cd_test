package main

import (
	"strings"
	"testing"
)

func TestNonEmptyString(t *testing.T) {
	output := welcomeString()
	if output == "" {
		t.Fatalf("App outputs nothing!")
	}
}

func TestWelcomeString(t *testing.T) {
	output := welcomeString()
	if !strings.Contains(output, "Hello Go!") {
		t.Fatalf("App does not output a greeting")
	}
}

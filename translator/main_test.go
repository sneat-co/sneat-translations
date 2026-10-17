package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	oldFunc := translateMissingFunc
	oldStderr := osStderr
	defer func() {
		translateMissingFunc = oldFunc
		osStderr = oldStderr
	}()

	// 1. Success
	called := false
	translateMissingFunc = func(baseLocale string) error {
		called = true
		if baseLocale != "en-UK" {
			t.Fatalf("expected en-UK, got %s", baseLocale)
		}
		return nil
	}
	main()
	if !called {
		t.Fatal("expected translateMissingFunc to be called")
	}

	// 2. Error
	var buf bytes.Buffer
	osStderr = &buf
	translateMissingFunc = func(baseLocale string) error {
		return errors.New("boom error")
	}
	main()
	if !strings.Contains(buf.String(), "boom error") {
		t.Fatalf("expected stderr to contain 'boom error', got %q", buf.String())
	}
}

package main

import (
	"context"
	"strings"
	"testing"

	"golang.org/x/text/language"
)

func TestTranslateMulti(t *testing.T) {
	targetLang, err := language.Parse("ru-RU")
	if err != nil {
		t.Errorf("language.Parse: %v", err)
	}
	sourceLang, err := language.Parse("en-UK")
	if err != nil {
		t.Errorf("language.Parse: %v", err)
		return
	}

	ctx := context.Background()
	translations, err := translateMulti(ctx, sourceLang, targetLang, []string{"Today is %s", "Today is {DATE}", "Today is:\n\t{DATE}"})
	if err != nil {
		t.Errorf("translateMulti: %v", err)
		return
	}
	t.Log("Translated into:\n" + strings.Join(translations, "\n\t"))
}

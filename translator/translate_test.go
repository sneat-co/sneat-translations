package main

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"cloud.google.com/go/translate"
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

	translations, err := translateMulti(sourceLang, targetLang, []string{"Today is %s", "Today is {DATE}", "Today is:\n\t{DATE}"})
	if err != nil {
		t.Errorf("translateMulti: %v", err)
		return
	}
	t.Log("Translated into:\n" + strings.Join(translations, "\n\t"))
}

// translateText translates the given text into the specified targetLanguage. sourceLanguage
// is optional. If empty, the API will attempt to detect the source language automatically.
// targetLanguage and sourceLanguage should follow ISO 639 language code of the input text
// (e.g., 'fr' for French)
//
// Find a list of supported languages and codes here:
// https://cloud.google.com/translate/docs/languages#nmt
func translateMulti(sourceLanguage, targetLanguage language.Tag, textsToTranslate []string) (translatedTexts []string, err error) {
	ctx := context.Background()

	reVarName := regexp.MustCompile(`{\w+}`)
	reVarIndex := regexp.MustCompile(`{\d+}`)

	var client *translate.Client
	if client, err = translate.NewClient(ctx); err != nil {
		err = fmt.Errorf("translate.NewClient error: %w", err)
		return
	}
	defer func() {
		_ = client.Close()
	}()

	options := &translate.Options{
		Source: sourceLanguage,
	}

	type source struct {
		text string
		vars []string
	}
	var sources = make([]source, len(textsToTranslate))
	for i, text := range textsToTranslate {
		s := source{text: text}
		if strings.Contains(text, "{") && strings.Contains(text, "}") {
			varIndex := 0
			text = string(reVarName.ReplaceAllFunc([]byte(s.text), func(sub []byte) []byte {
				s.vars = append(s.vars, string(sub[1:len(sub)-1]))
				result := []byte(fmt.Sprintf("{%d}", varIndex))
				varIndex++
				return result
			}))
		}
		text = strings.ReplaceAll(text, "\n", `{n}`)
		text = strings.ReplaceAll(text, "\n", `{t}`)
		sources[i] = s
		textsToTranslate[i] = text
	}

	var resp []translate.Translation
	// Find more information about translate function here:
	// https://pkg.go.dev/cloud.google.com/go/translate#Client.Translate
	if resp, err = client.Translate(ctx, textsToTranslate, targetLanguage, options); err != nil {
		err = fmt.Errorf("client.Translate error: %w", err)
		return
	}
	if len(resp) == 0 {
		err = fmt.Errorf("client.Translate returned empty response to texts: %s", strings.Join(textsToTranslate, "\n"))
		return
	}
	translatedTexts = make([]string, len(resp))
	for i, trans := range resp {
		var varIndex int
		s := sources[i]
		text := trans.Text
		text = strings.ReplaceAll(text, `{n}`, "\n")
		text = strings.ReplaceAll(text, `{t}`, "\n")

		translatedTexts[i] = string(reVarIndex.ReplaceAllFunc([]byte(text), func(sub []byte) []byte {
			if varIndex, err = strconv.Atoi(string(sub[1 : len(sub)-1])); err != nil {
				return sub
			}
			result := []byte(fmt.Sprintf("{%s}", s.vars[varIndex]))
			varIndex++
			return result
		}))
	}
	return
}

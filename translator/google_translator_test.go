package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func TestTranslateMulti_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"translations": []map[string]string{
					{"translatedText": "Сегодня %s"},
					{"translatedText": "Сегодня {0}"},
					{"translatedText": "Сегодня:{n}{t}{0}"},
					{"translatedText": "Plain"},
				},
			},
		})
	}))
	defer ts.Close()

	targetLang := language.MustParse("ru-RU")
	sourceLang := language.MustParse("en-UK")
	ctx := context.Background()

	inputs := []string{
		"Today is %s",
		"Today is {DATE}",
		"Today is:\n\t{{.DATE}}",
		"Plain",
	}

	translations, err := translateMulti(ctx, sourceLang, targetLang, inputs,
		option.WithoutAuthentication(),
		option.WithEndpoint(ts.URL),
	)
	if err != nil {
		t.Fatalf("translateMulti failed: %v", err)
	}
	if len(translations) != 4 {
		t.Fatalf("expected 4 translations, got %d", len(translations))
	}
	if translations[1] != "Сегодня {DATE}" {
		t.Fatalf("expected 'Сегодня {DATE}', got %q", translations[1])
	}
	if translations[2] != "Сегодня:\n\t{{.DATE}}" {
		t.Fatalf("expected 'Сегодня:\\n\\t{{.DATE}}', got %q", translations[2])
	}
}

func TestTranslateMulti_Errors(t *testing.T) {
	ctx := context.Background()
	targetLang := language.MustParse("ru-RU")
	sourceLang := language.MustParse("en-UK")

	// 1. Client creation error
	oldClient := newTranslateClient
	newTranslateClient = func(ctx context.Context, opts ...option.ClientOption) (*translate.Client, error) {
		return nil, errors.New("client creation error")
	}
	if _, err := translateMulti(ctx, sourceLang, targetLang, []string{"hi"}); err == nil {
		t.Fatal("expected error on client creation failure")
	}
	newTranslateClient = oldClient

	// 2. Translate call error (HTTP 500)
	tsErr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer tsErr.Close()

	if _, err := translateMulti(ctx, sourceLang, targetLang, []string{"hi"},
		option.WithoutAuthentication(),
		option.WithEndpoint(tsErr.URL),
	); err == nil {
		t.Fatal("expected error on 500 status")
	}

	// 3. Empty response
	tsEmpty := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"translations": []map[string]string{},
			},
		})
	}))
	defer tsEmpty.Close()

	if _, err := translateMulti(ctx, sourceLang, targetLang, []string{"hi"},
		option.WithoutAuthentication(),
		option.WithEndpoint(tsEmpty.URL),
	); err == nil {
		t.Fatal("expected error on empty response")
	}

	// 4. strconv.Atoi failure in reVarIndex (e.g. index overflows int)
	tsOverflow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"data": map[string]any{
				"translations": []map[string]string{
					{"translatedText": "overflow {9999999999999999999999999999999999999999999999999999999999999999}"},
				},
			},
		})
	}))
	defer tsOverflow.Close()

	if _, err := translateMulti(ctx, sourceLang, targetLang, []string{"{test}"},
		option.WithoutAuthentication(),
		option.WithEndpoint(tsOverflow.URL),
	); err == nil {
		t.Fatal("expected error on overflow index")
	}
}

package main

import (
	"bytes"
	"context"
	"errors"
	"go/ast"
	"go/token"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func TestLocaleKeyToCode(t *testing.T) {
	if got := localeKeyToCode(`"en"`); got != "en-EN" {
		t.Fatalf("expected en-EN, got %s", got)
	}
	if got := localeKeyToCode(`"en-UK"`); got != "en-UK" {
		t.Fatalf("expected en-UK, got %s", got)
	}
	if got := localeKeyToCode(`"enUK"`); got != "en-UK" {
		t.Fatalf("expected en-UK, got %s", got)
	}

	assertPanics(t, "empty locale code", func() {
		localeKeyToCode(`""`)
	})
	assertPanics(t, "unsupported locale:toolong", func() {
		localeKeyToCode(`"toolong"`)
	})
}

func TestLocaleCodeToKey(t *testing.T) {
	if got := localeCodeToKey("en-UK"); got != "enUK" {
		t.Fatalf("expected enUK, got %s", got)
	}
}

func TestTranslateLocale(t *testing.T) {
	oldMock := useMockTranslator
	defer func() { useMockTranslator = oldMock }()

	useMockTranslator = true
	sourceLang := language.MustParse("en")
	texts := []string{"hello"}
	got, err := translateLocale(sourceLang, "ru-RU", texts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 1 || got[0] != "TRANSLATED: hello" {
		t.Fatalf("unexpected result: %v", got)
	}

	useMockTranslator = false
	// Invalid locale -> translateByGoogle returns error
	if _, err := translateLocale(sourceLang, "---invalid---", texts); err == nil {
		t.Fatal("expected error on invalid locale")
	}
}

func TestTranslateByGoogle(t *testing.T) {
	oldClient := newTranslateClient
	defer func() { newTranslateClient = oldClient }()

	sourceLang := language.MustParse("en")

	// 1. Invalid target language
	if _, err := translateByGoogle(sourceLang, "---invalid---", []string{"hi"}); err == nil {
		t.Fatal("expected error on invalid target language")
	}

	// 2. Batch of 25 items with httptest server (exercises both cap(batch) and trailing batch)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":{"translations":[{"translatedText":"trans"}]}}`))
	}))
	defer ts.Close()

	newTranslateClient = func(ctx context.Context, opts ...option.ClientOption) (*translate.Client, error) {
		return translate.NewClient(ctx, option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	}

	texts := make([]string, 25)
	for i := range texts {
		texts[i] = "text"
	}
	oldStdout := osStdout
	var outBuf bytes.Buffer
	osStdout = os.Stdout
	defer func() { osStdout = oldStdout }()

	res, err := translateByGoogle(sourceLang, "es", texts)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(res) == 0 {
		t.Fatal("expected non-empty translated texts")
	}

	// 3. Error in translateBatch when batch == cap(batch)
	newTranslateClient = func(ctx context.Context, opts ...option.ClientOption) (*translate.Client, error) {
		return nil, errors.New("batch error")
	}
	osStdout = &outBuf
	if _, err := translateByGoogle(sourceLang, "es", texts); err == nil {
		t.Fatal("expected error when batch == cap(batch)")
	}

	// 4. Error in trailing batch when batch > 0
	if _, err := translateByGoogle(sourceLang, "es", []string{"short"}); err == nil {
		t.Fatal("expected error when trailing batch > 0")
	}
}

func TestTranslateMock(t *testing.T) {
	in := []string{"foo", "bar"}
	out := translateMock(in)
	if len(out) != 2 || out[0] != "TRANSLATED: foo" || out[1] != "TRANSLATED: bar" {
		t.Fatalf("unexpected mock translations: %v", out)
	}
}

func TestGetKey_GetValue_GetKeyValue(t *testing.T) {
	// 1. BasicLit key and value
	kvLit := &ast.KeyValueExpr{
		Key:   &ast.BasicLit{Kind: token.STRING, Value: `"key1"`},
		Value: &ast.BasicLit{Kind: token.STRING, Value: `"val1"`},
	}
	if k := getKey(kvLit); k != `"key1"` {
		t.Fatalf("expected \"key1\", got %s", k)
	}
	if v, _ := getValue(kvLit); v != "val1" {
		t.Fatalf("expected val1, got %s", v)
	}
	k, v := getKeyValue(kvLit)
	if k != `"key1"` || v != "val1" {
		t.Fatalf("unexpected getKeyValue: %s, %s", k, v)
	}

	// 2. Ident key
	kvIdent := &ast.KeyValueExpr{
		Key:   &ast.Ident{Name: "IdentKey"},
		Value: &ast.BasicLit{Kind: token.STRING, Value: "`raw`"},
	}
	if k := getKey(kvIdent); k != "IdentKey" {
		t.Fatalf("expected IdentKey, got %s", k)
	}
	if v, _ := getValue(kvIdent); v != "raw" {
		t.Fatalf("expected raw, got %s", v)
	}

	// 3. BinaryExpr value
	binExpr := &ast.BinaryExpr{
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"hello "`},
		Op: token.ADD,
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"world"`},
	}
	kvBin := &ast.KeyValueExpr{
		Key:   &ast.Ident{Name: "BinKey"},
		Value: binExpr,
	}
	if v, _ := getValue(kvBin); v != "hello world" {
		t.Fatalf("expected 'hello world', got %s", v)
	}

	// 4. Panics
	assertPanics(t, "unexpected key type", func() {
		getKey(&ast.KeyValueExpr{Key: &ast.CompositeLit{}})
	})
	assertPanics(t, "unexpected value type", func() {
		getValue(&ast.KeyValueExpr{Value: &ast.CompositeLit{}})
	})
}

func TestGetBinaryExprValue(t *testing.T) {
	// Nested BinaryExpr on both sides
	leftBin := &ast.BinaryExpr{
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"a"`},
		Op: token.ADD,
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"b"`},
	}
	rightBin := &ast.BinaryExpr{
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"c"`},
		Op: token.ADD,
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"d"`},
	}
	nested := &ast.BinaryExpr{
		X:  leftBin,
		Op: token.ADD,
		Y:  rightBin,
	}
	if got := getBinaryExprValue(nested); got != "abcd" {
		t.Fatalf("expected abcd, got %s", got)
	}

	// Unsupported op
	badOp := &ast.BinaryExpr{
		X:  &ast.BasicLit{Kind: token.STRING, Value: `"a"`},
		Op: token.SUB,
		Y:  &ast.BasicLit{Kind: token.STRING, Value: `"b"`},
	}
	assertPanics(t, "unsupported binary expressions", func() {
		getBinaryExprValue(badOp)
	})
}

func TestGetTextToTranslate_GetIsAlreadyPresent_GetCurrentValue(t *testing.T) {
	lit := &ast.CompositeLit{
		Elts: []ast.Expr{
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "enUK"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: `"Hello"`},
			},
			&ast.KeyValueExpr{
				Key:   &ast.Ident{Name: "ruRU"},
				Value: &ast.BasicLit{Kind: token.STRING, Value: "`Привет`"},
			},
			&ast.BasicLit{Kind: token.STRING, Value: "not-a-kv"},
		},
	}

	txt, isQuoted := getTextToTranslate("en-UK", lit)
	if txt != "Hello" || !isQuoted {
		t.Fatalf("expected 'Hello', true; got %q, %v", txt, isQuoted)
	}

	txt2, isQuoted2 := getTextToTranslate("ru-RU", lit)
	if txt2 != "Привет" || isQuoted2 {
		t.Fatalf("expected 'Привет', false; got %q, %v", txt2, isQuoted2)
	}

	txt3, _ := getTextToTranslate("es-ES", lit)
	if txt3 != "" {
		t.Fatalf("expected empty string, got %s", txt3)
	}

	if !getIsAlreadyPresent(lit, "en-UK") {
		t.Fatal("expected getIsAlreadyPresent true for en-UK")
	}
	if getIsAlreadyPresent(lit, "es-ES") {
		t.Fatal("expected getIsAlreadyPresent false for es-ES")
	}

	val, expr := getCurrentValue(lit, "en-UK")
	if val != "Hello" || expr == nil {
		t.Fatalf("expected Hello and non-nil expr, got %q, %v", val, expr)
	}
	val2, expr2 := getCurrentValue(lit, "es-ES")
	if val2 != "" || expr2 != nil {
		t.Fatalf("expected empty string and nil expr, got %q, %v", val2, expr2)
	}
}

func TestHasTranslatableTextParts(t *testing.T) {
	if !hasTranslatableTextParts("Hello World") {
		t.Fatal("expected true for 'Hello World'")
	}
	if hasTranslatableTextParts("{{.VAR}} %s <b> </b>") {
		t.Fatal("expected false for variables/tags only")
	}
}

type errWriter struct{}

func (errWriter) WriteString(s string) (int, error) { return 0, errors.New("write error") }
func (errWriter) Close() error                      { return nil }

func TestTranslateMissing(t *testing.T) {
	oldPath := translationsFilePath
	oldMock := useMockTranslator
	oldParse := parseLanguage
	oldFormat := formatNode
	oldCreate := createOutFile
	defer func() {
		translationsFilePath = oldPath
		useMockTranslator = oldMock
		parseLanguage = oldParse
		formatNode = oldFormat
		createOutFile = oldCreate
	}()

	// 1. File not found
	translationsFilePath = "nonexistent-translations.go"
	if err := translateMissing("en-UK"); err == nil {
		t.Fatal("expected error on nonexistent file")
	}

	// 2. Parse error on invalid Go file
	tmpDir := t.TempDir()
	badFile := filepath.Join(tmpDir, "bad.go")
	_ = os.WriteFile(badFile, []byte("package trans\ninvalid syntax !!!"), 0644)
	translationsFilePath = badFile
	if err := translateMissing("en-UK"); err == nil {
		t.Fatal("expected error on bad syntax")
	}

	// 3. Valid translations file with missing translations (includes elements testing !ok continues)
	sampleSource := `package trans

var Translations = map[string]map[string]string{
	"HELLO": {
		"enUK": "Hello",
	},
	"NOT_TRANSLATABLE": {
		"enUK": "{{.VAR}}",
	},
	"NOT_TRANSLATABLE_RAW": {
		"enUK": ` + "`{{.RAW}}`" + `,
	},
	"SINGLE_ONLY_S1NGL": {
		"enUK": "Single only",
	},
	"RAW_STRING": {
		"enUK": ` + "`Raw string`" + `,
	},
	"NON_COMPOSITE_VALUE": "just string",
	"non_kv_elem",
}
`
	goodFile := filepath.Join(tmpDir, "translations.go")
	if err := os.WriteFile(goodFile, []byte(sampleSource), 0644); err != nil {
		t.Fatal(err)
	}

	translationsFilePath = goodFile
	useMockTranslator = true

	if err := translateMissing("en-UK"); err != nil {
		t.Fatalf("translateMissing failed: %v", err)
	}

	// Verify file was written and contains translated entries
	content, err := os.ReadFile(goodFile)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(content), "TRANSLATED: Hello") {
		t.Fatal("expected file to contain TRANSLATED: Hello")
	}

	// 4. Test error from translateLocale
	useMockTranslator = false
	oldClient := newTranslateClient
	newTranslateClient = func(ctx context.Context, opts ...option.ClientOption) (*translate.Client, error) {
		return nil, errors.New("translate client error")
	}
	defer func() { newTranslateClient = oldClient }()
	_ = os.WriteFile(goodFile, []byte(sampleSource), 0644)
	if err := translateMissing("en-UK"); err == nil {
		t.Fatal("expected error when translation fails")
	}

	// 5. Test panic on empty string with non-nil expr (line 235)
	emptyStrSource := `package trans

var Translations = map[string]map[string]string{
	"EMPTY_TEST": {
		"enUK": "{{.VAR}}",
		"ruRU": "",
	},
}
`
	emptyStrFile := filepath.Join(tmpDir, "empty_str.go")
	_ = os.WriteFile(emptyStrFile, []byte(emptyStrSource), 0644)
	translationsFilePath = emptyStrFile
	useMockTranslator = true
	assertPanics(t, "is empty string", func() {
		_ = translateMissing("en-UK")
	})

	// 6. Test parseLanguage error
	translationsFilePath = goodFile
	useMockTranslator = true
	parseLanguage = func(s string) (language.Tag, error) {
		return language.Tag{}, errors.New("parse error")
	}
	if err := translateMissing("en-UK"); err == nil {
		t.Fatal("expected error on parseLanguage error")
	}
	parseLanguage = oldParse

	// 7. Test formatNode error
	_ = os.WriteFile(goodFile, []byte(sampleSource), 0644)
	formatNode = func(dst io.Writer, fset *token.FileSet, node any) error {
		return errors.New("format error")
	}
	if err := translateMissing("en-UK"); err == nil {
		t.Fatal("expected error on formatNode error")
	}
	formatNode = oldFormat

	// 8. Test outFile.WriteString error
	_ = os.WriteFile(goodFile, []byte(sampleSource), 0644)
	createOutFile = func(name string) (fileWriter, error) {
		return errWriter{}, nil
	}
	if err := translateMissing("en-UK"); err == nil {
		t.Fatal("expected error on WriteString error")
	}
	createOutFile = oldCreate

	// 9. Test createOutFile error
	_ = os.WriteFile(goodFile, []byte(sampleSource), 0644)
	createOutFile = func(name string) (fileWriter, error) {
		return nil, errors.New("create error")
	}
	if err := translateMissing("en-UK"); err == nil {
		t.Fatal("expected error on createOutFile error")
	}
}

func assertPanics(t *testing.T, expectedSubstring string, fn func()) {
	t.Helper()
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic, got nil")
		}
		msg := ""
		switch v := r.(type) {
		case string:
			msg = v
		case error:
			msg = v.Error()
		}
		if !strings.Contains(msg, expectedSubstring) {
			t.Fatalf("expected panic containing %q, got %q", expectedSubstring, msg)
		}
	}()
	fn()
}

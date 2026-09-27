package trans

import (
	"testing"
)

func TestCommands(t *testing.T) {
	testKey := "__test_commands_key__"
	TRANS[testKey] = map[string]string{
		"1": "",
		"2": "/help",
		"3": "/help",
		"4": "x",
		"5": "hello world",
	}
	defer delete(TRANS, testKey)

	got := Commands(testKey, "extra1", "extra1")
	if len(got) == 0 {
		t.Fatal("expected non-empty commands")
	}

	has := func(target string) bool {
		for _, s := range got {
			if s == target {
				return true
			}
		}
		return false
	}

	for _, expected := range []string{"help", "/help", "x", "hello world", "extra1"} {
		if !has(expected) {
			t.Errorf("Commands() missing expected %q, got: %v", expected, got)
		}
	}
}

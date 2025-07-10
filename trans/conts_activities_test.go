package trans

import "testing"

func TestGetActivityCode(t *testing.T) {
	if actualCode := GetActivityCode("unknown activity"); actualCode != "" {
		t.Errorf("GetActivityCode(unknown activity) = %s, want empty string", actualCode)
	}

	for _, code := range AllActivityCodes {
		for lang, translation := range TRANS[code] {
			if actualCode := GetActivityCode(translation); actualCode != code {
				t.Errorf("GetActivityCode(%s:%s) = %s, want %s", translation, lang, actualCode, code)
			}
		}
	}
}

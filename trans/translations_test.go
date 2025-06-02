package trans

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

var (
	reVars  = regexp.MustCompile(`%[svd]|\{\{\..+?}}`)
	reWords = regexp.MustCompile(`\w+|%[svd]`)
)

func TestTRANS(t *testing.T) {
	var wordsCount int

	var missingRequired = make(map[string][]string, len(TRANS))
	var missingSupported = make(map[string][]string, len(TRANS))

	for key, vals := range TRANS {
		countsByLang := make(map[string]map[string]int)

		for _, requiredLocale := range RequiredLocales {
			if _, ok := vals[requiredLocale]; !ok {
				missingRequired[key] = append(missingRequired[key], requiredLocale)
			}
		}

		for _, supportedLocale := range SupportedLocales {
			if _, ok := vals[supportedLocale]; !ok {
				missingSupported[key] = append(missingSupported[key], supportedLocale)
			}
		}

		for lang, val := range vals {
			if !isSupportedLang(lang) {
				t.Errorf("Key %v has unsupported language: %v", key, lang)
				continue
			}
			if strings.Contains(val, "https: ") || strings.Contains(val, "http: ") {
				t.Logf("Invalid http(s): link: %v=%v", key, val)
			}
			vars := reVars.FindAllString(val, -1)
			counts := make(map[string]int, len(vars))
			for _, v := range vars {
				counts[v] += 1
			}
			countsByLang[lang] = counts
		}
		enCounts, ok := countsByLang[enUK]
		if !ok {
			t.Errorf("Key %v missing en-UK trnaslation", key)
			continue
		}
		wordsCount += len(reWords.FindAllString(vals[enUK], -1))
		reported := make(map[string]int)
		for lang, counts := range countsByLang {
			if lang == enUK {
				continue
			}
			for enV, enCount := range enCounts {
				if c := counts[enV]; c != enCount {
					if c != 0 || vals[lang] != "" {
						t.Errorf("%v:%v has %d of '%v' while en-US has %d", key, lang, counts[enV], enV, enCount)
						reported[enV] = enCount
					}
				}
			}

			for langV, langCount := range counts {
				if enCounts[langV] != langCount {
					if _, ok := reported[langV]; !ok {
						t.Errorf("%v:%v has %d of '%v' while en-US has %d", key, lang, langCount, langV, enCounts[langV])
					}
				}
			}
		}
	}

	if len(missingSupported) > 0 {
		s := reportMissingTranslations("supported", missingSupported, RequiredLocales)
		t.Log(s)
	}
	if len(missingRequired) > 0 {
		s := reportMissingTranslations("required", missingRequired, RequiredLocales)
		t.Errorf(s)
	}

	t.Logf("English words count: %d", wordsCount)
}

func reportMissingTranslations(missingType string, missing map[string][]string, expectedLocales []string) string {
	var s = make([]string, 0, len(missing)*len(expectedLocales)+1)
	missingTranslationsCount := 0
	for key, missingLocales := range missing {
		missingTranslationsCount += len(missingLocales)
		s = append(s, "\t"+key+": "+strings.Join(missingLocales, ","))
	}
	return strings.Join(append([]string{
		fmt.Sprintf("Missing %d %s translations for %d translation IDs:",
			missingTranslationsCount, missingType, len(missing),
		)}, s...), "\n")
}

func TestHtmlTags(t *testing.T) {
	for code, vals := range TRANS {
		for lang, val := range vals {
			for _, tag := range []string{"b", "i"} {
				openTag, closeTag := fmt.Sprintf("<%v>", tag), fmt.Sprintf("</%v>", tag)
				if openCount, closeCount := strings.Count(val, openTag), strings.Count(val, closeTag); openCount != closeCount {
					t.Errorf("%v[%v]: %v != %v: %v != %v: %v", code, lang, openTag, closeTag, openCount, closeCount, val)
				}
			}
			if openCount, closeCount := strings.Count(val, "<a"), strings.Count(val, "</a>"); openCount != closeCount {
				t.Errorf("%v[%v]: %v != %v: %v != %v: %v", code, lang, "<a>", "</a>", openCount, closeCount, val)
			}
		}
	}
}

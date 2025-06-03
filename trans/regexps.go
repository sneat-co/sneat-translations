package trans

import "regexp"

var (
	ReVars  = regexp.MustCompile(`%[svd]|\{\{\..+?}}`)
	ReTags  = regexp.MustCompile(`</?\s*\w+\s*>`)
	ReWords = regexp.MustCompile(`\w+|%[svd]`)
)

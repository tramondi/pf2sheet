package common

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AlekSi/pointer"
	strip "github.com/grokify/html-strip-tags-go"
)

func ClearText(text string) string {
	return strip.StripTags(strings.TrimSpace(text))
}

func ClearTextPtr(text *string) *string {
	if text == nil {
		return nil
	}

	out := ClearText(*text)

	return &out
}

func EscapeScripts(text string) string {
	scriptRe := regexp.MustCompile(`(?s)<script.*<\/script>`)
	endScriptRe := regexp.MustCompile(`</script>`)

	result := scriptRe.ReplaceAllStringFunc(text, func(s string) string {
		index := endScriptRe.FindStringIndex(s)
		return strings.TrimSpace(s[index[1]:])
	})

	return result
}

func EscapeScriptsPtr(text *string) *string {
	if text == nil {
		return nil
	}

	out := EscapeScripts(*text)

	return &out
}

func MakeErr(format string, args ...any) *string {
	return pointer.ToString(fmt.Sprintf(format, args...))
}

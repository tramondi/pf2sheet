package common

import (
	"fmt"
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

func MakeErr(format string, args ...any) *string {
	return pointer.ToString(fmt.Sprintf(format, args...))
}

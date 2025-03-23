package charset

import (
	"strings"

	"github.com/R4yL-dev/gogopwd/internal/argsparse"
	"github.com/R4yL-dev/gogopwd/internal/config"
)

func Generate(pwdConfig argsparse.Config) string {
	charsets := []struct {
		condition bool
		value     string
	}{
		{pwdConfig.HasLower, config.CHARSET_LOWER},
		{pwdConfig.HasUpper, config.CHARSET_UPPER},
		{pwdConfig.HasDigit, config.CHARSET_DIGIT},
		{pwdConfig.HasPunctuation, config.CHARSET_PUNCTUATION},
	}

	var charsetBuilder strings.Builder

	for _, cs := range charsets {
		if cs.condition {
			charsetBuilder.WriteString(cs.value)
		}
	}
	charset := charsetBuilder.String()
	if len(pwdConfig.Exclude) > 0 {
		charset = excludeFromCharset(charset, pwdConfig.Exclude)
	}
	return charset
}

func excludeFromCharset(charset string, exclude string) string {
	exclude = removeDuplicates(exclude)
	var result strings.Builder
	for _, char := range charset {
		if !strings.ContainsRune(exclude, char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func removeDuplicates(s string) string {
	seen := make(map[rune]bool)
	var result strings.Builder

	for _, char := range s {
		if !seen[char] {
			seen[char] = true
			result.WriteRune(char)
		}
	}
	return result.String()
}

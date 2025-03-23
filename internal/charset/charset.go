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

	return charsetBuilder.String()
}

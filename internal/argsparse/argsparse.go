package argsparse

import (
	"errors"
	"flag"

	"github.com/R4yL-dev/gogopwd/internal/config"
)

type Config struct {
	Length         int
	NumberOf       int
	HasLower       bool
	HasUpper       bool
	HasDigit       bool
	HasPunctuation bool
	Exclude        string
}

func Parse() (Config, error) {
	var pwdConfig Config

	flag.IntVar(&pwdConfig.Length, "size", config.DEFAULT_LENGTH, "Size of the password")
	flag.IntVar(&pwdConfig.Length, "s", config.DEFAULT_LENGTH, "Alias for len")

	flag.IntVar(&pwdConfig.NumberOf, "num", config.DEFAULT_NUMBEROF, "Number of password")
	flag.IntVar(&pwdConfig.NumberOf, "n", pwdConfig.NumberOf, "Alias for num")

	flag.BoolVar(&pwdConfig.HasLower, "lower", config.DEFAULT_LOWER, "Has lowercase ?")
	flag.BoolVar(&pwdConfig.HasLower, "l", config.DEFAULT_LOWER, "Alias for lowercase")

	flag.BoolVar(&pwdConfig.HasUpper, "upper", config.DEFAULT_UPPER, "Has uppercase ?")
	flag.BoolVar(&pwdConfig.HasUpper, "u", config.DEFAULT_UPPER, "Alias for uppercase")

	flag.BoolVar(&pwdConfig.HasDigit, "digit", config.DEFAULT_DIGIT, "Has digit ?")
	flag.BoolVar(&pwdConfig.HasDigit, "d", config.DEFAULT_DIGIT, "Alias for digit")

	flag.BoolVar(&pwdConfig.HasPunctuation, "punctuation", config.DEFAULT_PUNCTUATION, "Has punctuation ?")
	flag.BoolVar(&pwdConfig.HasPunctuation, "p", config.DEFAULT_PUNCTUATION, "Alias for special characters")

	flag.StringVar(&pwdConfig.Exclude, "exclude", "", "Characters to exclude")
	flag.StringVar(&pwdConfig.Exclude, "e", "", "Alias for exclude")

	flag.Parse()

	if err := validate(pwdConfig); err != nil {
		return Config{}, err
	}

	return pwdConfig, nil
}

func validate(pwdConfig Config) error {
	if pwdConfig.Length < 1 {
		return errors.New("password size muste be at least 1")
	}
	if !pwdConfig.HasLower && !pwdConfig.HasUpper && !pwdConfig.HasDigit && !pwdConfig.HasPunctuation {
		return errors.New("at least one character type must be enabled")
	}
	return nil
}

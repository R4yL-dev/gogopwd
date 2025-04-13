package main

import (
	"fmt"
	"os"

	"github.com/R4yL-dev/gogopwd/internal/argsparse"
	"github.com/R4yL-dev/gogopwd/internal/charset"
	"github.com/R4yL-dev/gogopwd/internal/pwdgen"
)

func main() {
	pwdConfig, err := argsparse.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	pwdCharset, err := charset.Generate(pwdConfig)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(2)
    }
	for i := 0; i < pwdConfig.NumberOf; i++ {
		pwd := pwdgen.Generate(pwdCharset, pwdConfig.Length)
		fmt.Println(pwd)
	}

}

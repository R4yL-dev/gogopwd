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

	pwdCharset := charset.Generate(pwdConfig)
	pwd := pwdgen.Generate(pwdCharset, pwdConfig.Length)
	fmt.Println(pwd)
}

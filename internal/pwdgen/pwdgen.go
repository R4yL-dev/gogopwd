package pwdgen

import (
	"math/rand"
	"strings"
	"time"
)

func Generate(charset string, pwdLength int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	charset_size := len(charset)
	var pwd strings.Builder

	for i := 0; i < pwdLength; i++ {
		index := r.Intn(charset_size)
		pwd.WriteByte(charset[index])
	}
	return pwd.String()
}

package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func GenerateRandCode() string {
	rand.Seed(time.Now().UnixNano())

	code := &strings.Builder{}
	for i := 0; i != 6; i++ {
		code.WriteString(fmt.Sprintf("%d", rand.Intn(10)))
	}
	return code.String()
}

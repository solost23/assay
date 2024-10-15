package util

import (
	"bytes"
	"io"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func UTF8ToGBK(b []byte) ([]byte, error) {
	encoder := simplifiedchinese.GBK.NewEncoder()
	reader := transform.NewReader(bytes.NewReader(b), encoder)
	return io.ReadAll(reader)
}

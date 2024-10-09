package servants

import (
	"assay/infra/global"
	"assay/infra/util/utf8togbk"
	"fmt"
	"strings"
)

func SendLoginShortMessage(phone string, message string) error {
	return Send([]string{phone}, message)
}

func Send(phones []string, message string) error {
	data := fmt.Sprintf("#%s#%s#", strings.Join(phones, ","), message)
	b, err := utf8togbk.UTF8ToGBK([]byte(data))
	if err != nil {
		return err
	}

	client := global.Cat
	if _, err = client.Write(b); err != nil {
		return err
	}
	return nil
}

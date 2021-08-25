package iconv

import "github.com/djimenez/iconv-go"

func Iconv(input string, fromEncoding string, toEncoding string) (output string, err error) {
	output, err = iconv.ConvertString(input, fromEncoding, toEncoding)
	return
}

package utils

import (
	"crypto/md5"
	"errors"
	"io"
)

// GeneratorMD5 生成md5
func GeneratorMD5(text string) (result string, _ error) {
	if text == "" {
		return "", errors.New("参数不能为空！")
	}
	hash := md5.New()
	_, err := io.WriteString(hash, text)
	if err != nil {
		return "", errors.New("编码失败！")
	}
	return string(hash.Sum(nil)), nil
}

func DecodeMD5(text string) {

}

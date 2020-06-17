package utils

import (
	"encoding/base64"
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/png"
)

//判断base64编码的图片格式，返回格式后缀和解码结果，支持jpg、png两种
func GetPictureFormat(imageData string) (string, []byte, error) {
	picBytes, err := base64.StdEncoding.DecodeString(imageData)
	if err != nil {
		return "", nil, err
	}

	buffer := bytes.NewBuffer(picBytes)
	_, format, err := image.Decode(buffer)
	if err != nil {
		return "", nil, err
	}
	if format[0] != '.' {
		format = "." + format
	}
	return format, picBytes, nil
}

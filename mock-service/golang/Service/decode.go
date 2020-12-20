package Service

import (
	"github.com/axgle/mahonia"
	"github.com/saintfish/chardet"
)

func Encode(output []byte) string {
	detector := chardet.NewTextDetector()
	charset, _ := detector.DetectBest(output)
	decoder := mahonia.NewDecoder(charset.Charset)
	outputString := decoder.ConvertString(string(output))
	return outputString
}

func EncodeWithError(output []byte) (string, error) {
	detector := chardet.NewTextDetector()
	charset, err := detector.DetectBest(output)
	if err != nil {
		return "", err
	}
	decoder := mahonia.NewDecoder(charset.Charset)
	outputString := decoder.ConvertString(string(output))
	return outputString, nil
}

package utilx

import (
	"bytes"
	"compress/gzip"
)

// GzipCompress 压缩gzip数据
func GzipCompress(gzData string) (string, error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write([]byte(gzData)); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}

	return b.String(), nil
}

// GzipUnCompress 解压gzip数据
func GzipUnCompress(gzData string) (string, error) {
	var b bytes.Buffer
	gz, err := gzip.NewReader(bytes.NewBuffer([]byte(gzData)))
	if err != nil {
		return "", err
	}
	if _, err := b.ReadFrom(gz); err != nil {
		return "", err
	}
	if err := gz.Close(); err != nil {
		return "", err
	}
	return b.String(), nil
}

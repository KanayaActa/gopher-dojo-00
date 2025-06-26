package imgconv

import (
	"fmt"
	"path/filepath"
	"strings"
)

// Format は画像形式を表します。
type Format string

const (
	FormatJPG Format = "jpg" // JPEG形式
	FormatPNG Format = "png" // PNG形式
)

// Converter は画像変換の設定を持ちます。
type Converter struct {
	InputFormat  Format
	OutputFormat Format
}

// NewConverter はConverterを生成します。
func NewConverter(in, out Format) *Converter {
	return &Converter{
		InputFormat:  in,
		OutputFormat: out,
	}
}

// ConvertError は変換時のエラーを表します。
type ConvertError struct {
	Path string
	Msg  string
}

func (e *ConvertError) Error() string {
	return fmt.Sprintf("error: %s: %s", e.Path, e.Msg)
}

// IsImageFile はpathが指定形式の画像ファイルか判定します。
func IsImageFile(path string, format Format) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch format {
	case FormatJPG:
		return ext == ".jpg" || ext == ".jpeg"
	case FormatPNG:
		return ext == ".png"
	default:
		return false
	}
}

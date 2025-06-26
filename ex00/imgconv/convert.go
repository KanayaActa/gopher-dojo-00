package imgconv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ConvertDir はdir配下の画像ファイルを変換します。
func (c *Converter) ConvertDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return &ConvertError{Path: dir, Msg: "no such file or directory"}
	}
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if IsImageFile(path, c.InputFormat) {
			return c.ConvertFile(path)
		}
		if isImageFile(path) {
			return &ConvertError{Path: path, Msg: "is not a valid file"}
		}
		return nil
	})
}

// ConvertFile は画像ファイルを変換します。
func (c *Converter) ConvertFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return &ConvertError{Path: path, Msg: err.Error()}
	}
	defer f.Close()

	var img image.Image
	switch c.InputFormat {
	case FormatJPG:
		img, err = jpeg.Decode(f)
	case FormatPNG:
		img, err = png.Decode(f)
	default:
		return &ConvertError{Path: path, Msg: "unsupported input format"}
	}
	if err != nil {
		return &ConvertError{Path: path, Msg: "failed to decode image"}
	}

	outPath := generateOutputPath(path, c.OutputFormat)
	out, err := os.Create(outPath)
	if err != nil {
		return &ConvertError{Path: outPath, Msg: err.Error()}
	}
	defer out.Close()

	switch c.OutputFormat {
	case FormatJPG:
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 90})
	case FormatPNG:
		err = png.Encode(out, img)
	default:
		return &ConvertError{Path: outPath, Msg: "unsupported output format"}
	}
	if err != nil {
		return &ConvertError{Path: outPath, Msg: "failed to encode image"}
	}
	return nil
}

// generateOutputPath は出力ファイルパスを生成します。
func generateOutputPath(inputPath string, outputFormat Format) string {
	dir := filepath.Dir(inputPath)
	name := strings.TrimSuffix(filepath.Base(inputPath), filepath.Ext(inputPath))
	return filepath.Join(dir, fmt.Sprintf("%s.%s", name, outputFormat))
}

// isImageFile は画像拡張子か判定します。
func isImageFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".jpg" || ext == ".jpeg" || ext == ".png"
}

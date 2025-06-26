package main

import (
	"flag"
	"fmt"
	"os"

	"imgconv/imgconv"
)

func main() {
	// フラグの定義
	inputFormat := flag.String("i", "jpg", "入力画像形式 (例: jpg, png)")
	outputFormat := flag.String("o", "png", "出力画像形式 (例: jpg, png)")
	flag.Parse()

	// 残りの引数（ディレクトリ）
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "error: invalid argument")
		os.Exit(1)
	}
	dir := args[0]

	converter := imgconv.NewConverter(imgconv.Format(*inputFormat), imgconv.Format(*outputFormat))
	if err := converter.ConvertDir(dir); err != nil {
		if convertErr, ok := err.(*imgconv.ConvertError); ok {
			fmt.Fprintln(os.Stderr, convertErr.Error())
		} else {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		os.Exit(1)
	}
}

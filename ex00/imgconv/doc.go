// Package imgconv は JPEG ⇄ PNG などの画像形式を変換するユーティリティです.
//
// # 使い方
//
//	conv := imgconv.NewConverter(imgconv.FormatJPG, imgconv.FormatPNG)
//	if err := conv.ConvertDir("./images"); err != nil {
//	    log.Fatal(err)
//	}
//
// 変換に失敗した場合は *imgconv.ConvertError が返ります。
// エラーメッセージ形式は  "error: <path>: <reason>" です。
package imgconv

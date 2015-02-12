package gopinyin

import "github.com/axgle/mahonia"

var enc = mahonia.NewEncoder("gbk")

func UTF8ToGBK(s string) string {
	return enc.ConvertString(s)
}

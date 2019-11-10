package tool

import (
	"github.com/axgle/mahonia"
	"net/url"
)

// GBK 编码短信内容， 需带上签名， 如果不需要签名，可以传 "" (目前只有验证码签名写死在内容中，传的签名为 "" )
func SmsContentEncode(signature, content string) string {
	enc := mahonia.NewEncoder("gbk")
	res := url.QueryEscape(enc.ConvertString(signature + content))
	return res
}

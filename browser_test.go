package user_agent

import (
	"fmt"
	"testing"
)

func TestWswUA(t *testing.T) {
	// The "New" function will create a new UserAgent object and it will parse
	// the given string. If you need to parse more strings, you can re-use
	// this object and call: ua.Parse("another string")
	ua := New("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36 NetType/WIFI MicroMessenger/7.0.20.1781(0x6700143B) WindowsWechat(0x63090621) XWEB/8287 Flue ")

	name, version := ua.Browser()
	fmt.Printf("browser name:%v\n", name)       // => Googlebot
	fmt.Printf("browser version:%v\n", version) // => 2.1
}

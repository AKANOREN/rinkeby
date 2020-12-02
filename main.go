package main

import (
	"fmt"
	"time"

	"github.com/AKANOREN/go-utils/format" // 大文字にする必要ある
	// "github.com/akanoren/go-utils/format"
	/* 以下のようなエラー
	 * module declares its path as: github.com/AKANOREN/go-utils
	 *        but was required as: github.com/akanoren/go-utils
	 */)

func main() {
	fmt.Println("hello workd")

	dateTime := time.Now()
	formattedDate := format.DateFormat(dateTime, "%Y年%m月%d日 %H時%M分%S秒")
	fmt.Println(formattedDate)
}

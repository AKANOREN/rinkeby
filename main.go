package main

import (
	"fmt"
	"time"

	"github.com/AKANOREN/go-utils/format"
)

func main() {
	fmt.Println("hello workd")

	dateTime := time.Now()
	formattedDate := format.DateFormat(dateTime, "%Y年%m月%d日 %H時%M分%S秒")
	fmt.Println(formattedDate)
}

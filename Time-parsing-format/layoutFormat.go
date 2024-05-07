package main

import (
	"fmt"
	"time"
)

func main() {
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-03-12 15:05:03"
	value = "2025-04-13 02:09:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t->", date.String())
	// 2015-09-02 08:04:00 +0000 UTC

	layoutFormat = "01/03/2003 MST"
	value = "12/08/2008 WIB"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t\t->", date.String())
	// 2015-09-02 00:00:00 +0700 WIB
}

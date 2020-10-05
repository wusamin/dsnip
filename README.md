# dsnip

## Overview
dsnip is snippets of date on Golang.  
A collection of date manuplations that I often do that  
I find wasy to forget.

## Getting Started
```
import (
	"time"

	"github.com/wusamin/dsnip"
)

func main() {
	now := time.Date(2020, 10, 5, 13, 0, 0, 0, time.Local)
	after := time.Date(2020, 10, 6, 12, 0, 0, 0, time.Local)

	println(dsnip.AddDay(now, 1).String()) // 2020-10-06 13:00:00 +0900 JST
	println(dsnip.GetDiffHour(after, now)) // 23
}
```
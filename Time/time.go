package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func main() {
	now := time.Now()
	pl(now.Year(), ":", now.Month(), ":", now.Day(), ":", now.Hour(), ":", now.Minute(), ":", now.Second())
}

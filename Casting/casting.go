// convert float to int
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	cV1 := 1.54
	cV2 := int(cV1)
	pl(cV2)
}

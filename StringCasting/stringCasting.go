// Convert string to Integer
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	cV1 := "500"
	cV2, err := strconv.Atoi(cV1)
	pl(cV2, err, reflect.TypeOf(cV2))

}

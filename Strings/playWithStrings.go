// This code demonstrates playing wit strings libary and its different functions :
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var pl = fmt.Println

func replaceme(myvalue string) {
	replacer := strings.NewReplacer(myvalue, "Gaurav Pande")
	replacedValue := replacer.Replace(myvalue)
	pl("Original value:", myvalue)
	pl("Replaced value :", replacedValue)

}

func main() {
	pl("Enter the string to play with")
	reader := bufio.NewReader(os.Stdin)
	mystring, err := reader.ReadString('\n')
	if err == nil {
		replaceme(mystring)
	}
}

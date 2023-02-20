/*
	This code demonstrates playing wit strings libary and its different functions like
replace
trimspace
length
UpperCase
*/

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

func trimspace(myvalue string) {
	trimValue := strings.TrimSpace(myvalue)
	pl("Trimmed value :", trimValue)
}

func getlen(myvalue string) {
	length := len(myvalue)
	pl("Length :", length)
}

func toupper(myvalue string) {
	value := strings.ToUpper(myvalue)
	pl("UPPERCASE IS :", value)
}

func main() {
	pl("Enter the string to play with")
	reader := bufio.NewReader(os.Stdin)
	mystring, err := reader.ReadString('\n')
	if err == nil {
		replaceme(mystring)
		trimspace(mystring)
		getlen(mystring)
		toupper(mystring)
	}
}

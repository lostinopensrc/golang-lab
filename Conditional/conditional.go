// if else condition to check Senior Citizen .
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	pl("What is your Birth year")
	reader := bufio.NewReader(os.Stdin)
	age, err1 := reader.ReadString('\n')
	if err1 == nil {
		ageTrim := strings.TrimSpace(age)      // Trims leading and traling spaces from a string
		yourAge, err2 := strconv.Atoi(ageTrim) // Converts string to a int
		if (err2 == nil) && (yourAge >= 1) && (yourAge < 60) {
			pl("Inside second if")
			pl("You are not senior Citizen")
		} else {
			pl("You are Senior Citizen")
		}
	} else {
		pl("Enter Age Correctly")
	}
}

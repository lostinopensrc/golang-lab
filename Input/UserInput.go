package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Creation of Alias pl(print line) as variable which will be used everytime in place of fmt.Println
var pl = fmt.Println

func main() {
	pl("What is your name ?")
	read := bufio.NewReader(os.Stdin)
	name, err := read.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}
}

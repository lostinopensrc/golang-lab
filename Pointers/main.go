package main

import "fmt"

func main() {
	s := "Hello Pointer address"
	p := &s                            // holds address of s
	fmt.Println("Value of S:", s)      //prints s value
	fmt.Println("p gives address:", p) // prints address
	fmt.Println("*p gives value of what p address is pointing to:", *p)

	//Dereference
	*p = "Hello Grophers this is example of derefernce and s value is changed now"
	fmt.Println(s)
}

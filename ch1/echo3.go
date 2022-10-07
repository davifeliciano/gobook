// echo2 prints its command line arguments, but
// using a the Join function from the strings package
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

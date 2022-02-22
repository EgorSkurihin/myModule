package mymodule

import "fmt"

func Version() string {
	return "The version of package is 1"
}

func main() {
	fmt.Println(Version())
}

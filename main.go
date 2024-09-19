package main

import (
	"fmt"
	"os"

	"github.com/danielmiessler/fabric/cli"
)

func main() {
	fmt.Println("Hello, World!")
	// for i, arg := range os.Args {
	// 	fmt.Printf("Argument %d: %s\n", i, arg)
	// }
	_, err := cli.Cli()
	// fmt.Println("first is: ", x)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

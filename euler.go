package main

import (
	"fmt"
	"os"

	"github.com/choodur/euler/go"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: ./euler <problem_number>")
	} else {
		problem := euler.Problems[fmt.Sprintf("problem%s", os.Args[1])]

		if problem == nil {
			fmt.Printf("Problem %s doesn't exist yet.", os.Args[1])
			return
		}

		problem()
	}
}

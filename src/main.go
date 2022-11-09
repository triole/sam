package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	parseArgs()

	args := strings.Join(CLI.Args, " ")
	if args == "" {
		args = getStdin()
	}

	res, err := Call(CLI.Command, args)
	if err != nil {
		log.Fatalf("Error calling command: %+v\n", err.Error())
	}
	if res != nil {
		fmt.Printf("%s\n", res)
	} else {
		ListFunctions()
		if CLI.List == false && CLI.ListShort == false {
			fmt.Printf("%s\n\n",
				"String transformation command not found. "+
					"Please use one of the above.",
			)
		}
	}
}

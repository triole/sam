package main

import (
	"fmt"
	"sam/src/conf"
)

func main() {
	parseArgs()
	conf := conf.Init(CLI)
	fmt.Printf("%+v\n", conf)

	// var tcli transform.TransformCLI
	// tcli.Command = CLI.Command
	// tcli.List = CLI.List
	// tcli.ListShort = CLI.ListShort
	// tcli.Args = CLI.Args

	// tr := transform.Init(tcli)
	// res, err := tr.Call()
	// if err != nil {
	// 	log.Fatalf("Error calling command: %+v\n", err.Error())
	// }
	// if res != nil {
	// 	fmt.Printf("%s\n", res)
	// } else {
	// 	tr.ListFunctions()
	// 	if !CLI.List && !CLI.ListShort {
	// 		fmt.Printf("%s\n\n",
	// 			"String transformation command not found. "+
	// 				"Please use one of the above.",
	// 		)
	// 	}
	// }
}

package main

import (
	"sam/src/conf"
	"sam/src/transform"
)

func main() {
	parseArgs()
	conf := conf.Init(CLI)
	tr := transform.Init(conf)
	tr.Run()
}

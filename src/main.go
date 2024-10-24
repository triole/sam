package main

import (
	"sam/src/conf"
	"sam/src/implant"
	"sam/src/transform"
)

func main() {
	impl := implant.Init()
	parseArgs(impl)
	conf := conf.Init(CLI)
	tr := transform.Init(conf, impl)
	tr.Run()
}

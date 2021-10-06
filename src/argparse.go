package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	// BUILDTAGS are injected ld flags during build
	BUILDTAGS      string
	appName        = "sam"
	appDescription = "string alteration machine"
	appMainversion = "0.1"
)

var CLI struct {
	Command           string   `help:"string transformation command" arg optional`
	StringToTransform []string `help:"string to process" arg optional passthrough`
	List              bool     `help:"list the available template functions" short:l`
	VersionFlag       bool     `help:"display version" short:V`
}

func parseArgs() {
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:      true,
			Summary:      true,
			NoAppSummary: true,
			FlagsLast:    false,
		}),
	)
	_ = ctx.Run()

	if CLI.VersionFlag == true {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	// ctx.FatalIfErrorf(err)
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "version: "+appMainversion+".", -1)
	arr := strings.Split(s, "\n")
	fmt.Printf("\n%s, %s\n", appName, appDescription)
	for _, el := range arr {
		if el != "" {
			fmt.Printf("%s\n", strings.TrimSpace(el))
		}
	}
}

func alnum(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile("[^a-z0-9_-]")
	return re.ReplaceAllString(s, "-")
}

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/alecthomas/kong"
)

var (
	// BUILDTAGS are injected ld flags during build
	BUILDTAGS      string
	appName        = "sam"
	appDescription = "a string alteration machine to ease string processing in shell scripts"
	appMainversion = "0.2"
)

var CLI struct {
	SubCommand string `kong:"-"`

	// keep-sorted start block=yes newline_separated=yes
	Align struct {
		Args   []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Length int      `help:"target string length" short:"l" required:""`
		Target string   `help:"where to align string to, can be: [${enum}]" enum:"left, right, l, r" short:"t" default:"left"`
	} `cmd:"" help:"align string"`

	Bool struct {
		Args []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
	} `cmd:"" help:"return bool value; returns 'true' on: 1, enable, enabled, on, true; returns 'false' on everything else; case insensitive"`

	Calc struct {
		Args      []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Precision int      `help:"amount of decimal places" short:"p" default:"-1"`
	} `cmd:"" help:"evaluate mathematical expressions"`

	Case struct {
		Args     []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Target   string   `help:"target case, can be: [${enum}]" enum:"lower, upper, camel, snake, title" short:"t" default:"lower"`
		Language string   `help:"set language" short:"l" default:"english"`
	} `cmd:"" help:"convert string case"`

	Color struct {
		Args []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
	} `cmd:"" help:"display color code list, input can be hex or rgb"`

	Encode struct {
		Args    []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Target  string   `help:"encode target, can be: [${enum}]" enum:"url, base64" short:"t" default:"base64"`
		Reverse bool     `help:"convert the other way round" short:"r"`
	} `cmd:"" help:"encode string to"`

	Hash struct {
		Args   []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Length int      `help:"hash length if hash type suports it" short:"l" default:"1024"`
		Target string   `help:"target case, can be: [${enum}]" enum:"md5, sha1, sha256, sha384, sha512, blake3, rake, whirlpool" short:"t" default:"sha512"`
		File   string   `help:"calculate hash for a file" short:"f" type:"existingFile"`
	} `cmd:"" help:"calculate hash of a string"`

	Path struct {
		Args   []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Target string   `help:"which part to get, can be: [${enum}]" enum:"dir, bn, ext" short:"t" default:"dir"`
	} `cmd:"" help:"get parts of a file path"`

	Tidy struct {
		Args   []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Target string   `help:"target characters groups to tidy, can be: [${enum}]" enum:"spaces,pseps" short:"t" default:"spaces"`
	} `cmd:"" help:"tidy string, replace multiple occurences of spaces or path separators by a single one"`

	Trim struct {
		Args       []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Substring  string   `help:"substring to trim" short:"s" required:""`
		Target     string   `help:"which string part to trim, can be: [${enum}]" enum:"prefix, suffix, both" short:"t" default:"both"`
		Aggressive bool     `help:"aggressive mode, remove multiple occurences of the match" short:"a"`
	} `cmd:"" help:"remove part of a string"`

	Version struct{} `cmd:"" help:"display version"`
	// keep-sorted end
}

func parseArgs() {
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:   true,
			Summary:   true,
			FlagsLast: false,
		}),
	)
	_ = ctx.Run()

	if ctx.Command() == "version" {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	// ctx.FatalIfErrorf(err)

	CLI.SubCommand = ctx.Command()
}

type tPrinter []tPrinterEl
type tPrinterEl struct {
	Key string
	Val string
}

func printBuildTags(buildtags string) {
	regexp, _ := regexp.Compile(`({|}|,)`)
	s := regexp.ReplaceAllString(buildtags, "\n")
	s = strings.Replace(s, "_subversion: ", "version: "+appMainversion+".", -1)
	fmt.Printf("\n%s\n%s\n\n", appName, appDescription)
	arr := strings.Split(s, "\n")
	var pr tPrinter
	var maxlen int
	for _, line := range arr {
		if strings.Contains(line, ":") {
			l := strings.Split(line, ":")
			if len(l[0]) > maxlen {
				maxlen = len(l[0])
			}
			pr = append(pr, tPrinterEl{l[0], strings.Join(l[1:], ":")})
		}
	}
	for _, el := range pr {
		fmt.Printf("%"+strconv.Itoa(maxlen)+"s\t%s\n", el.Key, el.Val)
	}
	fmt.Printf("\n")
}

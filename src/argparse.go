package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"sam/src/implant"
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

	Date struct {
		Args   []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:"" default:"now"`
		Target string   `help:"print certain date layout only, can be: [${dateLayoutsList}], default: all" enum:"${dateLayoutsList}, all" short:"t" default:"all"`
		Layout bool     `help:"additioanlly display go date layout in final table" short:"l"`
		Diff   string   `help:"date to calculate difference between the given two" short:"d"`
	} `cmd:"" help:"print different date formats"`

	Encode struct {
		Args    []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Target  string   `help:"encode target, can be: [${enum}]" enum:"url, base64" short:"t" default:"base64"`
		Reverse bool     `help:"convert the other way round" short:"r"`
	} `cmd:"" help:"encode string to"`

	Geo struct {
		Args   []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Length int      `help:"hash length, maximum precision is 32" short:"l" default:"32"`
		Target string   `help:"target case, can be: [${enum}]" enum:"enc,dec" short:"t" default:"enc"`
	} `cmd:"" help:"calculate geo hash"`

	Hash struct {
		Args   []string `help:"args passed through as string to process" arg:"" optional:"" passthrough:""`
		Length int      `help:"hash length if hash type suports it" short:"l" default:"1024"`
		Rounds int      `help:"bcrypt rounds" short:"r" default:"16"`
		Target string   `help:"target case, can be: [${enum}]" enum:"md5, sha1, sha256, sha384, sha512, blake3, rake, whirlpool, argon2, bcrypt" short:"t" default:"sha512"`
		File   string   `help:"calculate hash for a file, not with bcrypt" short:"f" type:"existingFile"`
	} `cmd:"" help:"calculate hash of a string"`

	Match struct {
		Args   []string `help:"check if hash matches input string" arg:"" optional:"" passthrough:""`
		Target string   `help:"target case, can be: [${enum}]" enum:"md5, sha1, sha256, sha384, sha512, blake3, rake, whirlpool" short:"t" default:"sha512"`
		File   string   `help:"calculate hash for a file, not with bcrypt" short:"f" type:"existingFile"`
	} `cmd:"" help:"check if hash matches string, usage: -t md5 5eb63bbbe01eeed093cb22bb8f5acdc3 hello world"`

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

func parseArgs(impl implant.Implant) {
	ctx := kong.Parse(&CLI,
		kong.Name(appName),
		kong.Description(appDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:   true,
			Summary:   true,
			FlagsLast: false,
		}),
		kong.Vars{
			"dateLayoutsList": strings.Join(impl.DateLayouts.List(), ", "),
		},
	)
	_ = ctx.Run()

	if ctx.Command() == "version" {
		printBuildTags(BUILDTAGS)
		os.Exit(0)
	}
	CLI.SubCommand = ctx.Command()
	readStdinIntoArgs()
}

func readStdinIntoArgs() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return
	}
	if stat.Mode()&os.ModeCharDevice == 0 {
		cmd := capitalize(strings.TrimSuffix(CLI.SubCommand, " "))
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return
		}

		text := strings.TrimRight(string(data), "\n\r")
		if text == "" {
			return
		}

		switch cmd {
		case "Align":
			if len(CLI.Align.Args) == 0 {
				CLI.Align.Args = []string{text}
			}
		case "Bool":
			if len(CLI.Bool.Args) == 0 {
				CLI.Bool.Args = []string{text}
			}
		case "Calc":
			if len(CLI.Calc.Args) == 0 {
				CLI.Calc.Args = []string{text}
			}
		case "Case":
			if len(CLI.Case.Args) == 0 {
				CLI.Case.Args = []string{text}
			}
		case "Color":
			if len(CLI.Color.Args) == 0 {
				CLI.Color.Args = []string{text}
			}
		case "Date":
			if len(CLI.Date.Args) == 0 {
				CLI.Date.Args = []string{text}
			}
		case "Encode":
			if len(CLI.Encode.Args) == 0 {
				CLI.Encode.Args = []string{text}
			}
		case "Geo":
			if len(CLI.Geo.Args) == 0 {
				CLI.Geo.Args = []string{text}
			}
		case "Hash":
			if len(CLI.Hash.Args) == 0 {
				CLI.Hash.Args = []string{text}
			}
		case "Match":
			if len(CLI.Match.Args) == 0 {
				CLI.Match.Args = []string{text}
			}
		case "Path":
			if len(CLI.Path.Args) == 0 {
				CLI.Path.Args = []string{text}
			}
		case "Tidy":
			if len(CLI.Tidy.Args) == 0 {
				CLI.Tidy.Args = []string{text}
			}
		case "Trim":
			if len(CLI.Trim.Args) == 0 {
				CLI.Trim.Args = []string{text}
			}
		}
	}
}

func capitalize(s string) string {
	r := []rune(s)
	if len(r) > 0 {
		r[0] = r[0] - 'a' + 'A'
	}
	return string(r)
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

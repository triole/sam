# Sam ![build](https://github.com/triole/sam/actions/workflows/build.yaml/badge.svg) ![test](https://github.com/triole/sam/actions/workflows/test.yaml/badge.svg)

<!-- toc -->

- [Synopsis](#synopsis)
- [How to Use?](#how-to-use)
- [Help](#help)

<!-- /toc -->

## Synopsis

The **S**tring **A**lteration **M**achine is a tool that can be used to manipulate and process strings. Why? Because I wanted something for simple string operations that can be used in bash scripts. I know there is `awk` and `tr` but some things (e.g. title case) are just to complicated using these two. Sam is simpler and does the job.

## How to Use?

Having simplicity in mind the syntax basically consists of two parts. The first part is the first arg and defines the command. Meaning what to do with the string afterwards. Everything from the second argument on is the string. It looks like this:

```shell
# transform to title case or uppercase
sam case -t title hello world

# or using stdin
echo hello world | sam hash -t md5

# display help
sam -h
```

## Help

```go mdox-exec="r -h"
Usage: sam <command>

a string alteration machine to ease string processing in shell scripts

Flags:
  -h, --help            Show context-sensitive help.
  -V, --version-flag    display version

Commands:
  align     align string
  bool      return bool value; returns 'true' on: 1, enable, enabled, on, true;
            returns 'false' on everything else; case insensitive
  case      convert string case
  color     display color codes, input can be hex or rgb
  encode    encode string to
  hash      calculate hash of a string
  path      get parts of a file path
  tidy      tidy string, replace multiple occurences of spaces or path
            separators by a single one
  trim      remove part of a string

Run "sam <command> --help" for more information on a command.
```

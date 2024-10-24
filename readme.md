# Sam ![build](https://github.com/triole/sam/actions/workflows/build.yaml/badge.svg) ![test](https://github.com/triole/sam/actions/workflows/test.yaml/badge.svg)

<!-- toc -->

- [Synopsis](#synopsis)
- [Usage Examples](#usage-examples)
- [Help](#help)

<!-- /toc -->

## Synopsis

The **S**tring **A**lteration **M**achine is a tool that can be used to manipulate and process strings. Why? Because I wanted something for simple string operations that can be used in bash scripts. I know there is `awk` and `tr` but some things (e.g. title case) are just to complicated using these two. Sam is simpler and does the job.

## Usage Examples

```go mdox-exec="sh/pre case -t title hello world"
$ sam case -t title hello world
Hello World
```

```go mdox-exec="sh/pre hash -t sha1 hello world"
$ sam hash -t sha1 hello world
2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
```

```go mdox-exec="sh/pre calc -p 3 '6*(3+3.111)'"
$ sam calc -p 3 6*(3+3.111)
36.666
```

```go mdox-exec="sh/pre color ff1199"
$ sam color ff1199

 TYPE   VALUE            
 Hex    #ff1199          
 RGB    [255 17 153]     
 RGBA   [255 17 153 255] 
 CMYK   [0 238 104 0]    
 YCbCr  [104 156 236]    

```

```go mdox-exec="sh/pre encode -t base64 hello world"
$ sam encode -t base64 hello world
aGVsbG8gd29ybGQ=
```

```go mdox-exec="sh/pre date 2025-01-01T15:33:11"
$ sam date 2025-01-01T15:33:11

 FORMAT         DATE                              
 UnixTimeStamp  1735741991                        
 UnixDate       Wed Jan 01 15:33:11 CET 2025      
 RFC3339        2025-01-01T15:33:11+01:00         
 RFC3339Nano    2025-01-01T15:33:11+01:00         
 RFC822Z        01 Jan 25 15:33 +0100             
 RFC850         Wednesday, 01-Jan-25 15:33:11 CET 
 RFC1123Z       Wed, 01 Jan 2025 15:33:11 +0100   
 Stamp          Jan  1 15:33:11                   
 StampNano      Jan  1 15:33:11.000000000         

```

*Note that you can also pass the input string by stdin.* Like...

```
echo hello world | sam hash -t md5
echo now | sam date
```

## Help

```go mdox-exec="r -h"
Usage: sam <command>

a string alteration machine to ease string processing in shell scripts

Flags:
  -h, --help    Show context-sensitive help.

Commands:
  align      align string
  bool       return bool value; returns 'true' on: 1, enable, enabled, on, true;
             returns 'false' on everything else; case insensitive
  calc       evaluate mathematical expressions
  case       convert string case
  color      display color code list, input can be hex or rgb
  date       print different date formats
  encode     encode string to
  hash       calculate hash of a string
  path       get parts of a file path
  tidy       tidy string, replace multiple occurences of spaces or path
             separators by a single one
  trim       remove part of a string
  version    display version

Run "sam <command> --help" for more information on a command.
```

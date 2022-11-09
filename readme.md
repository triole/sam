# Sam ![example workflow](https://github.com/triole/sam/actions/workflows/build.yaml/badge.svg)

<!--- mdtoc: toc begin -->

1. [Synopsis](#synopsis)
2. [How to Use?](#how-to-use-)
3. [Available String Transformation Operations](#available-string-transformation-operations)<!--- mdtoc: toc end -->

## Synopsis

The **S**tring **A**lteration **M**achine is a tool that can be used to manipulate and process strings. Why? Because I wanted something for simple string operations that can be used in bash scripts. I know there is `awk` and `tr` but some things (like i.e. title case) are just to complicated using these two. Sam is expandable and designed to do the job.

## How to Use?

Having simplicity in mind the syntax basically consists of two parts. The first part is the first arg and defines the command. Meaning what to do with the string afterwards. Everything from the second argument on is the string. It looks like this:

```shell
# transform to title case or uppercase
sam title hello world
sam uppercase hello world

# or using stdin
echo hello world | sam md5
# which is, apart from the dash that md5sum
# displays after the hash sum, equal to...
echo -n "hello world" | md5sum

# show a list of available operations
sam -l

# display help
sam -h
```

## Available String Transformation Operations

It is highly likely that there are more to come in the futre.

```go mdox-exec="r --list-short"

 CATEGORY | COMMAND      | ARGS        | USAGE                      
----------+--------------+-------------+----------------------------
 trim     | trimprefix   | prefix, str | sam trimprefix _ _hello    
 trim     | trimprefixag | prefix, str | sam trimprefixag _ __hello 
 trim     | trimsuffix   | suffix, str | sam trimsuffix _ hello_    
 trim     | trimsuffixag | suffix, str | sam trimsuffixag _ hello__ 
 trim     | trimspace    | str         |                            
 replace  | rms          | str         |                            
 case     | camel        | str         |                            
 case     | lower        | str         |                            
 case     | snake        | str         |                            
 bool     | str          | logical     | sam str enabled            
 case     | title        | str         |                            
 case     | upper        | str         |                            
 encoding | b64-txt      | str         |                            
 encoding | txt-b64      | str         |                            
 encoding | txt-url      | str         |                            
 encoding | url-txt      | str         |                            
 hash     | blake3       | size, str   | sam blake3 128 hello       
 hash     | md5          | str         |                            
 hash     | ripemd160    | str         |                            
 hash     | sha1         | str         |                            
 hash     | sha256       | str         |                            
 hash     | sha512       | str         |                            
 hash     | whirlpool    | str         |                            
 path     | folder       | str         |                            
 path     | tfn          | str         |                            
 path     | tps          | str         |                            

```

# Known problems

When using string processors that take multiple args (i.e. trimprefix, trimsuffix et al) the first argument can not contain spaces. This is due to the way the cli args are passed through. A possible solution would be to introduce an additional flag but currently keeping sam's usage simple outweighed increasing cli's complexity.

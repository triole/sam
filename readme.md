# Sam ![build](https://github.com/triole/sam/actions/workflows/build.yaml/badge.svg) ![test](https://github.com/triole/sam/actions/workflows/test.yaml/badge.svg)

<!-- toc -->

- [Synopsis](#synopsis)
- [How to Use?](#how-to-use)
- [Currently available string operations](#currently-available-string-operations)

<!-- /toc -->

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
# which's result is equal to...
echo -n "hello world" | md5sum

# show a list of available operations
sam -l

# display help
sam -h
```

## Currently available string operations

This is a short overview. For more info use `-l`.

```go mdox-exec="r --list-short"

 CATEGORY | COMMAND | ARGS        | USAGE               
----------+---------+-------------+---------------------
 case     | csc     | str         |                     
 case     | csl     | str         |                     
 case     | css     | str         |                     
 case     | cst     | str         |                     
 case     | csu     | str         |                     
 color    | col     | str         | sam col #999        
          |         |             | sam col 333777      
          |         |             | sam col 11 22 33    
 encoding | b64.txt | str         |                     
 encoding | txt.b64 | str         |                     
 encoding | txt.url | str         |                     
 encoding | url.txt | str         |                     
 hash     | blake3  | str         | sam blake3 64 hello 
 hash     | md5     | str         |                     
 hash     | sha1    | str         |                     
 hash     | sha256  | str         |                     
 hash     | sha384  | str         |                     
 hash     | sha512  | str         |                     
 hash     | whp     | str         |                     
 logical  | lb      | str         | sam lb enabled      
 path     | pd      | str         |                     
 path     | ptp     | str         |                     
 path     | pts     | str         |                     
 replace  | rspc    | str         |                     
 style    | sal     | str         | sam sal 8 hello     
 style    | sar     | str         | sam sar 8 hello     
 trim     | tpre    | prefix, str | sam tpre _ _hello   
 trim     | tpreg   | prefix, str | sam tpreg _ __hello 
 trim     | tspc    | str         |                     
 trim     | tsuf    | suffix, str | sam tsuf _ hello_   
 trim     | tsufg   | suffix, str | sam tsufg _ hello__ 

```

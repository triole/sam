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

```go mdox-exec="r -l"

 COMMAND      | ARGS        | DESCRIPTION                                         | CATEGORY 
--------------+-------------+-----------------------------------------------------+----------
 trimprefix   | prefix, str | remove prefix, requires two args: string, prefix to | trim     
              |             | remove                                              |          
--------------+-------------+-----------------------------------------------------+----------
 trimprefixag | prefix, str | trim prefix aggressive, remove multiple occurences  | trim     
              |             | of prefix                                           |          
--------------+-------------+-----------------------------------------------------+----------
 trimsuffix   | suffix, str | like trimprefix but removing end of a string, also  | trim     
              |             | two args                                            |          
--------------+-------------+-----------------------------------------------------+----------
 trimsuffixag | suffix, str | like trim suffix aggressive, you know...            | trim     
--------------+-------------+-----------------------------------------------------+----------
 rmmultispace | str         | remove each occurence of multiple spaces or tabs in | trim     
              |             | a string by one space                               |          
--------------+-------------+-----------------------------------------------------+----------
 trimspace    | str         | remove spaces or tabs around a string               | trim     
--------------+-------------+-----------------------------------------------------+----------
 camel        | str         | to camelcase                                        | case     
--------------+-------------+-----------------------------------------------------+----------
 lower        | str         | to lowercase                                        | case     
--------------+-------------+-----------------------------------------------------+----------
 snake        | str         | to snakecase                                        | case     
--------------+-------------+-----------------------------------------------------+----------
 title        | str         | title case                                          | case     
--------------+-------------+-----------------------------------------------------+----------
 upper        | str         | to uppercase                                        | case     
--------------+-------------+-----------------------------------------------------+----------
 bool         | str         | return boolean: 1, enable, enabled, on and true     | logical  
              |             | return true, everything else false (case doesn't    |          
              |             | matter)                                             |          
--------------+-------------+-----------------------------------------------------+----------
 fromb64      | str         | from base64 to string                               | encoding 
--------------+-------------+-----------------------------------------------------+----------
 fromurl      | str         | from url to plain string                            | encoding 
--------------+-------------+-----------------------------------------------------+----------
 tob64        | str         | to base64 from string                               | encoding 
--------------+-------------+-----------------------------------------------------+----------
 tourl        | str         | to url from plain string                            | encoding 
--------------+-------------+-----------------------------------------------------+----------
 md5          | str         | md5 hash                                            | hash     
--------------+-------------+-----------------------------------------------------+----------
 sha1         | str         | sha1 hash                                           | hash     
--------------+-------------+-----------------------------------------------------+----------
 sha256       | str         | sha256 hash                                         | hash     
--------------+-------------+-----------------------------------------------------+----------
 sha512       | str         | sha512 hash                                         | hash     
--------------+-------------+-----------------------------------------------------+----------
 folder       | str         | folder of a path string, return everything up to    | path     
              |             | last path separator, path separators trailing the   |          
              |             | input are ignored (i.e. /tmp/hello/ -> /tmp)        |          
--------------+-------------+-----------------------------------------------------+----------
 tp1          | str         | tidy path 1, remove multiple path separators        | path     
--------------+-------------+-----------------------------------------------------+----------
 tp2          | str         | as tp1, but also remove all accents, then replace   | path     
              |             | characters not being alpha numerics, dashes,        |          
              |             | underscores or path separators by underscores       |          
--------------+-------------+-----------------------------------------------------+----------
 tp3          | str         | as tp2, plus lower case conversion                  | path     
--------------+-------------+-----------------------------------------------------+----------
 tp4          | str         | as tp3, replace double underscores which may appear | path     
              |             | during conversion by a single one                   |          

```

# Known problems

When using string processors that take multiple args (i.e. trimprefix, trimsuffix et al) the first argument can not contain spaces. This is due to the way the cli args are passed through. A possible solution would be to introduce an additional flag but currently keeping sam's usage simple outweighed increasing cli's complexity.

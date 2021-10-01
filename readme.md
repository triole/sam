# Sam ![example workflow](https://github.com/triole/sam/actions/workflows/build.yaml/badge.svg)

<!--- mdtoc: toc begin -->

1.	[Synopsis](#synopsis)
2.	[How to use?](#how-to-use-)
3.	[Available transformation operations](#available-transformation-operations)<!--- mdtoc: toc end -->

## Synopsis

The **S**tring **A**lteration **M**achine is a tool that can be used to manipulate and process strings. Why? Because I wanted something for simple string operations that can be used in bash scripts. I know there is `awk` and `tr` but some things (like i.e. title case) are just to complicated using these two. Sam is expandable and designed to do the job.

## How to use?

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

## Available transformation operations

Very likely there are more operations to be added in the future.

| name   | category  | description                                          |
|--------|-----------|------------------------------------------------------|
| camel  | case      | to camelcase                                         |
| lower  | case      | to lowercase                                         |
| snake  | case      | to snakecase                                         |
| title  | case      | title case                                           |
| upper  | case      | to uppercase                                         |
|        |           |                                                      |
| bool   | logical   | return boolean: 1, enable, enabled, on and true      |
|        |           | return true, everything else false (case doesn't     |
|        |           | matter)                                              |
|        |           |                                                      |
| fr_b64 | encoding  | from base64 to string                                |
| to_b64 | encoding  | to base64 from string                                |
|        |           |                                                      |
| md5    | hash      | md5 hash                                             |
| sha1   | hash      | sha1 hash                                            |
| sha256 | hash      | sha256 hash                                          |
| sha512 | hash      | sha512 hash                                          |
|        |           |                                                      |
| dir    | file name | folder of a file name, return everything up to last  |
|        |           | path separator, path separators trailing the input   |
|        |           | are ignored (i.e. /tmp/hello/ -> /tmp)               |
| tfn1   | file name | tidy file names 1, remove multiple path separators   |
| tfn2   | file name | as tfn1, but also remove all accents, then replace   |
|        |           | characters not being alpha numerics, dashes,         |
|        |           | underscores or path separators by underscores        |
| tfn3   | file name | as tfn2, plus lower case conversion                  |
| tfn4   | file name | as tfn3, replace double underscores which may appear |
|        |           | during conversion by a single one                    |

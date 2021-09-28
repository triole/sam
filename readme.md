# Sam

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

| name      | description  |
|-----------|--------------|
| lowercase | to lowercase |
| title     | title case   |
| uppercase | to uppercase |
|           |              |
| md5       | md5 hash     |
| sha1      | sha1 hash    |
| sha256    | sha256 hash  |
| sha512    | sha512 hash  |

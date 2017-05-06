# xkpasswd.go

*Memorable password generator. A clone of my [xkpasswd-node](https://github.com/vot/xkpasswd-node) module, written in Go.*

## How it works

```
$ ./xkpasswd
offramps#schappes#skywrite
```


## Quickstart

You can [download Mac and Linux binaries](https://github.com/vot/xkpasswd-go/releases/tag/0.0.1).
Make sure to grab `xkpasswd-words.txt` too and put it in the same folder as the binary.

You can then execute it with `./xkpasswd`.

Copy the file to your bin folder with  `cp xkpasswd /usr/local/bin/xkpasswd`.



**CLI OPTIONS**

```
xkpasswd --complexity <number> --separators <string> --pattern <string> --transform <string> --number <number>
```



You can specify `--complexity [-c]` argument in accordance with [complexity levels table](#complexity-levels). Defaults to 2.

If specified `--pattern [-p]` argument overrides the [pattern](#patterns) derived from complexity level.

If `--separators [-s]` are provided they are used instead of the standard set (see complexity levels).
One separator is used per password, picked randomly from the provided set.

<!-- You can set `transform` to `alternate` or `uppercase` to trigger case transformation. -->

Finally, to generate multiple passwords at once you can specify the desired
amount with the `--number [-n]` argument. Defaults to 1.


**EXAMPLE** Default behaviour

```
$ xkpasswd
hide+threw+money+61
```

**EXAMPLE** Specify complexity

```
$ xkpasswd -c 5
join=industrial=wide=direction=lungs=16

$ xkpasswd -c 6
57!fifthly!astronauts!affectedly!nymphs!trustlessnesses!06
```

**EXAMPLE** Specify custom pattern

```
$ xkpasswd -p wdwd

adjective3solar6
```


**EXAMPLE** Specify complexity, custom separators and number of passwords

```
$ xkpasswd -c 3 -s '#!+' -n 5

dog!friend!successful!47
other#sell#close#01
hyperspatial+polyvalences+inquirendo+03
war#reassemble#inventress#93
gainsays+illumes+discontiguity+86
```


## Patterns

Patterns can consist of any combination of words, digits and separators.
The first letters (**w**, **d** and **s** respectively) are used in pattern string provided to the password generation function.

For example:

* `w` will return a single word (i.e. `demographics`). Use `w` for lowercase and `W` for uppercase.
* `wsd` will return a word and a digit, separated by one of the permitted separators (i.e. `storm#7`)
* `wswsdd` will return two words followed by a two digit number, all with separators between (i.e. `delates+dissembled+16`)



## Complexity levels

There are 6 complexity levels specified which can be used to provide
default patterns as well as expanded sets of separators.


| Complexity | Pattern         | Separators       |
|------------|-----------------|------------------|
| 1          | wsw             | #.-=+_           |
| 2          | wswsw           | #.-=+_           |
| 3          | wswswsdd        | #.-=+_           |
| 4          | wswswswsdd      | #.-=+_           |
| 5          | wswswswswsd     | #.-=+_!$*:~?     |
| 6          | ddswswswswswsdd | #.-=+_!$*:~?%^&; |



## Building project

If you'd like to build the project yourself instead of grabbing a prebuilt copy,
here are the steps:

1) Checkout a copy of this project

```
git clone https://github.com/vot/xkpasswd-go
cd xkpasswd-go
```

2) Build the project
`go build xkpasswd.go`

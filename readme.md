# xkpasswd.go

A clone of my [xkpasswd-node](https://github.com/vot/xkpasswd-node) module, written in Go.

You can build the project with
`go build`

and then execute it with
`./xkpasswd-go`

Compiled binary includes the list of words, binary I got on a Mac is 2.4MB. You can [download Mac binary here](https://github.com/vot/xkpasswd-go/releases/tag/0.0.1)

At this point this doesn't take any options but rather uses a hardcoded pattern (see `main` function if you'd like to change it).

## Example

```
$ ./xkpasswd-go
offramps#incisively#sardelles#gazetting
```

# xkpasswd.go

A clone of my [xkpasswd-node](https://github.com/vot/xkpasswd-node) module, written in Go.

You can build the project with
`go build`

apply execution permissions with
`chmod +x xkpasswd-go`

and then execute it with
`./xkpasswd-go`

You can [download Mac binary here](https://github.com/vot/xkpasswd-go/releases/tag/0.0.1).
Make sure to grab `words.txt` too.

At this point this doesn't take any options but rather uses a hardcoded pattern (see `main` function if you'd like to change it).

## Example

```
$ ./xkpasswd-go
offramps#incisively#sardelles#gazetting
```

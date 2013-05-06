serve
=====

Similair to ```python -m SimpleHTTPServer``` but written in Go

Usage
-----

```shell
$ go build
$ ./serve -h

Usage of ./serve:
  -dir=".": directory to serve
  -port="4242": port number

$ ./serve -dir=/tmp
$ ./serve -dir=/tmp -port=9999
```

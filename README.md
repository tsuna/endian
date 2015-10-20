# endian [![GoDoc](https://godoc.org/github.com/tsuna/endian?status.png)](https://godoc.org/github.com/tsuna/endian)

A small golang library to help deal with host-to-network and network-to-host
endianness conversion

## Disclaimer

Before using this library, please read and understand Rob Pike's blog post on
[the byte order fallacy](http://commandcenter.blogspot.com/2012/04/byte-order-fallacy.html).
Long story short, whenever you need to worry about whether or not your code is
running on a little endian or big endian machine, chances are you're wrong in
your approach and should instead really be thinking about writing the code in
an endianness-agnostic way, as Rob shows in his blog post.

In the limited cases where you do need to swap bytes around (perhaps because
you call into a library that is "doing it wrong"), this small library can help.

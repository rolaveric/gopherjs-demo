# gopherjs-demo
A demo project showing how to use GopherJS to share Go code between Go and JS applications

# Setup

Install `gopherjs` as a command line tool

`go get -u github.com/gopherjs/gopherjs`

# pet
A simple example of a Go struct being exported for use in JS code.

```
cd pet/js
gopherjs build main.go
node index.js
```

# user
A more complex example of a Go library that could be reused between Go and JS applications.

Note that, for brevity's sake, the `user` example does not include as much type checking as it should.
For example, `RegisterDBJS(jsdb)` should check that `jsdb` really does include a `Query()` function.
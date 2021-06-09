# Firefox bug #1714910 demo

This site provides a demonstration for [bug 1714910](https://bugzilla.mozilla.org/show_bug.cgi?id=1714910). The demonstration runs on port 8080.

The `/` route provides a simple HTML page with a form that submits a POST request to `/submit`.

The `/submit` route provides a simple text response that states the method that was used to make the request.

This demonstration has no dependencies, apart from the Go Standard Library (specifically `net/http`). It can be built with `go build main.go` or `go install`.
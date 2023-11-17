A (kinda?) simple way of splitting testing external and internal package
testing in Go. Contrived example below.

https://github.com/df-rw/go-simple-testing

## Our goal

We want to separate internal and external variables, methods and function
testing for a package.

## Directory layout

-   Each file in a package has a `_bb_test.go` and `_wb_test.go` file; `bb` is
    for black box (our external tests) and `wb` for white box (our internal
    tests).

```
.
├── Makefile
├── README.md
├── go.mod
├── main.go
└── pkg
    └── message
        ├── message.go
        ├── message_bb_test.go
        └── message_wb_test.go

3 directories, 7 files
```

## Files

### `_bb_test.go`

Black box testing files live in their own package, and import the package they
are testing, so access is only available through the exported variables, methods
and interfaces:

```go
package message_test

import (
    "test/pkg/message"
    "testing"
)
```

### `_wb_test.go`

White box testing files live in the same package as the one that is under test,
so they can access the unexported variables, methods and interfaces:

```go
package message

import (
	"testing"
)
```

## The Hate

### So I have to write twice as many tests?

No; just one set for the exported stuff and one set for internals. They are
in different files for readability.

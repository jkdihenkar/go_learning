# Setup GO Dev Env in Fedora

* Install: Download golang pkg from https://go.dev/dl/ and install using the wizard.
```
> go version
go version go1.17.6 darwin/amd64
```

* Install any code editor (I prefer VS Code).

The `go` and `gofmt` binaries will become available on the system.

Set `GO111MODULE` to `off`
```
go env -w GO111MODULE=off
```

Go code lives in a workspace which is defined by the GOPATH environment variable. A common choice among developers, and the default value of GOPATH starting from the Go 1.8 release, is to use $HOME/go:

Check GOPATH set correctly:
```
[jay@localhost]$ go env GOPATH
/Users/jaydihenkar/go
```
---

Run a simple basic program from this repo.

Source this project with GOPATH:
```
> set -x GOPATH (pwd)
> go env GOPATH
/Users/jaydihenkar/work/go_learning
```

Run the basic Code:
```
> go run src/01_basic_of_go/01_basic_of_go.go
Enter the basic go codenum to run: 011
Hello world....
36
j(0)|a(1)|y(2)|D(3)|
3_4_5_6_7_
15 123
Inferred type of := 3778 is Type :: int, Size :: 8[bytes]
Inferred type of := 2.37 + 99.42i is Type :: complex128, Size :: 16[bytes]
```

You can also run with `make`:
```
> make run filename=src/01_basic_of_go/01_basic_of_go.go
go env GOPATH
/Users/jaydihenkar/work/go_learning
go run src/01_basic_of_go/01_basic_of_go.go
Enter the basic go codenum to run: 011
Hello world....
36
j(0)|a(1)|y(2)|D(3)|
3_4_5_6_7_
15 123
Inferred type of := 3778 is Type :: int, Size :: 8[bytes]
Inferred type of := 2.37 + 99.42i is Type :: complex128, Size :: 16[bytes]
```
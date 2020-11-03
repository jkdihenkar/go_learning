# Setup GO Dev Env in Fedora

* Install: `sudo dnf install golang`
```
$ go version
go version go1.14.10 linux/amd64
```

* Install any code editor (I prefer VS Code).


The `go` and `gofmt` binaries will become available on the system.

Go code lives in a workspace which is defined by the GOPATH environment variable. A common choice among developers, and the default value of GOPATH starting from the Go 1.8 release, is to use $HOME/go:

```
$ mkdir -p $HOME/go
$ echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc
$ source $HOME/.bashrc
```

Check GOPATH set correctly:

```
[jay@localhost]$ go env GOPATH
/home/jay/go
```

---

Run a simple basic program from this repo.

Source this project with GOPATH:
```
[jay@localhost go_learning]$ export GOPATH=`pwd`
[jay@localhost go_learning]$ go env GOPATH
/home/jay/vcode_projects/go_learning
```

Run the basic Code:
```
[jay@localhost go_learning]$ go run src/01_basic_of_go/01_basic_of_go.go 
Enter the basic go codenum to run: 011
Hello world....
36
j(0)|a(1)|y(2)|D(3)|
3_4_5_6_7_
15 123
```
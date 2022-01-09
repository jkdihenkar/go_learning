GOPATH := ${HOME}/work/go_learning
export GOPATH

run:
	go env GOPATH
	go run ${filename}

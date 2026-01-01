# go install github.com/rinchsan/gosimports/cmd/gosimports@latest
gosimports:
	grep -rl --include \*.go . | xargs gosimports -w -l -local github/go-tools

tests:
	find . -name go.mod -execdir go test ./... -gcflags="all=-l -N" -coverpkg ./... -args dontInit=true \;
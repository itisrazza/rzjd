.PHONY: dist
dist :
	bash scripts/build-all.sh

.PHONY: test
test :
	go test -v ./...

.PHONY: test-cover
test-cover : tmp
	go test -v -cover -coverprofile tmp/cover.out ./...
	go tool cover -html tmp/cover.out -o tmp/cover.html

tmp :
	mkdir -p "$@"

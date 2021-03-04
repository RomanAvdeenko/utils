.PHONY: test
test:
		go test -v -race -timeout 10s ./...


.DEFAULT_GOAL:= test
.PHONY: run
run: test
	@go run ./cmd/tinyscript/main.go

.PHONY: fmt
fmt: vet
	@go fmt ./... 
.PHONY: vet
vet:
	@go vet ./...

.PHONY: test
test: fmt
	@go test -v ./...

.PHONY: start
start: clean
	@go build ./cmd/... && ./tinyscript

.PHONY: clean
clean:
	@rm -rf ./tinyscript


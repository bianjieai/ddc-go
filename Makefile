#!/usr/bin/make -f

all: tools lint


########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

proto-all: proto-tools proto-gen

proto-gen:
	@./scripts/protocgen.sh

########################################
### Testing

test: test-unit

test-unit:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock' ${PACKAGES_UNITTEST}

lint: golangci-lint
	golangci-lint run
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*"  -not -path "*.pb.go" | xargs gofmt -d -s
	go mod verify

format:
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*"  -not -path "*.pb.go" | xargs gofmt -w -s
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*"  -not -path "*.pb.go" | xargs misspell -w
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*"  -not -path "*.pb.go" | xargs goimports -w -local github.com/bianjieai/ddc-go

benchmark:
	@go test -mod=readonly -bench=. ./...

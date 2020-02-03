all: deps lint test install
deps:
	go get ./...

test: deps
	go test -count=1 ./...

vet: deps
	go vet ./...

check-fmt:
	bash -c "diff --line-format='%L' <(echo -n) <(gofmt -d -s .)"

lint: check-fmt vet

install: deps
	go install ./...



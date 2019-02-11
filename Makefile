.PHONY: test

test:
	go test -v -timeout 30s -coverprofile=/tmp/go-code-coverage
	rm /tmp/go-code-coverage

build:
	go build -o grapenut

docker:
	@test -n "$(TAG)" || { echo "ERROR: TAG must be set"; exit 1; }
	docker build -t quay.io/manywho/grapenut:$(TAG) .

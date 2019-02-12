.PHONY: test build docker

test:
	go test -v -timeout 30s -coverprofile=/tmp/go-code-coverage
	rm /tmp/go-code-coverage

build:
	go build -o grapenut

docker:
	@test -n "$(IMAGE_TAG)" || { echo "ERROR: IMAGE_TAG must be set"; exit 1; }
	docker build -t $(IMAGE_TAG) .

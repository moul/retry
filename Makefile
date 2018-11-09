NAME =		retry
SOURCE :=	$(shell find . -name "*.go")


all: install
install: $(GOPATH)/bin/retry
$(GOPATH)/bin/retry: $(SOURCE)
	GO111MODULE=auto go install -v

.PHONY: docker
docker:
	docker build -t moul/retry .

.PHONY: cross
cross:
	go get github.com/mitchellh/gox
	gox .

NAME := jcli-ishell-plugin

build:
	go build
	chmod u+x $(NAME)

copy: build
	cp $(NAME) ~/.jenkins-cli/plugins

test:
	go test ./...

fmt:
	go fmt .
	gofmt -s -w .

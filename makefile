#!make

# default command
default: help

test: ## will run all tests
	go test -v ./...

test-cover: ## will run go test --cover to show test coverage percent
	go test --cover .

test-profile: ## will run go test -coverprofile with html coverage output
	go test -coverprofile=./coverage.out && go tool cover -html=./coverage.out

clean-test-cache: ## will only clean the test cache 
	go clean -testcache

clean: ## will clean build cache
	go clean -cache

tidy: ## will run go mod tidy command
	go mod tidy -v
	@echo "done"

# help
help:
	@echo "usage: make [command]"
	@echo ""
	@echo "available commands:"
	@sed \
    		-e '/^[a-zA-Z0-9_\-]*:.*##/!d' \
    		-e 's/:.*##\s*/:/' \
    		-e 's/^\(.\+\):\(.*\)/$(shell tput setaf 6)\1$(shell tput sgr0):\2/' \
    		$(MAKEFILE_LIST) | column -c2 -t -s :
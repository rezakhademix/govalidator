#!make

# default command
default: help

test: ## run all tests
	go test -v --race ./...

test-one: ## run only tests matching the passed regex as `name`
ifdef name
	go test -run ${name}
else
	@echo "no name passed to run matching tests!"
endif

test-cover: ## run `go test --cover` to show test coverage percent
	go test --cover .

test-profile: ## run `go test -coverprofile` with html coverage output
	go test -coverprofile=./coverage.out && go tool cover -html=./coverage.out

clean-test-cache: ## only clean the test cache 
	go clean -testcache

clean: ## clean build cache
	go clean -cache

tidy: ## run `go mod tidy` command
	go mod tidy -v
	@echo "done"

# help
help:
	@echo "usage: make [command]"
	@echo ""
	@echo "commands:"
	@sed \
    		-e '/^[a-zA-Z0-9_\-]*:.*##/!d' \
    		-e 's/:.*##\s*/:/' \
    		-e 's/^\(.\+\):\(.*\)/$(shell tput setaf 6)\1$(shell tput sgr0):\2/' \
    		$(MAKEFILE_LIST) | column -c2 -t -s :
# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run/worker: run the temporal worker application
run/worker:
	go run ./app/worker

## run/client: run the temporal client application
run/client:
	go run ./app/client

## config-allow-permission: allow docker to access dynamicconfig folder
config-allow-permission:
	sudo chmod 755 dynamicconfig

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
## audit: tidy dependencies and format, vet and test all code
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## test/coverage: test all code and generate coverage.html
test/coverage:
	@echo 'Running tests...'
	go test -v -coverprofile cover.out ./...
	@echo 'Generate test result.out...'
	go tool cover -html=cover.out -o cover.html

## vendor: tidy and vendor dependencies
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

## profil: analyse heap
profil:
	@echo 'generate heap from :4000'
	curl -sK -v http://localhost:4000/debug/pprof/heap > heap.out;
	@echo 'open go tool -- use [top, png or gif]'
	go tool pprof heap.out;


.PHONY: help confirm run/worker run/client
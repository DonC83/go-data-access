### Dependencies

## Install dependencies
deps.install:
	./go install
.PHONY: deps.install

## Update dependencies
deps.update:
	./go get -u
.PHONY: deps.update

### Application - Build, clean & run

## Compile the classes of the project
build:
	./go build
.PHONY: build

## Clean project build artifacts
clean:
	./go clean
.PHONY: clean

## Run the application locally
run:
	./go run
.PHONY: run

## Create an executable jar of the application
package:
	./go build
.PHONY: package

### Application - Tests

## Execute the unit tests
test.unit:
	./go test
.PHONY: test.unit

### Verify - Code verification and static analysis

## Run static code analysis
verify:
	./go vet
.PHONY: verify

## Run static analysis and apply fixes if possible
verify.fix:
	./go fmt
.PHONY: verify.fix

### Devstack

## Start the devstack
devstack.start:
	docker-compose -f ./deployments/docker-compose.yml up -d
.PHONY: devstack.start

## Stop the devstack
devstack.stop:
	docker-compose -f ./deployments/docker-compose.yml down
.PHONY: devstack.stop

## Show the devstack's status
devstack.status:
	docker-compose -f ./deployments/docker-compose.yml ps
.PHONY: devstack.status
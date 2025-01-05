SHELL := /bin/bash
.ONESHELL:

.PHONY: all
.SILENT: all
all: tidy lint fmt vet test build

.PHONY: lint
.SILENT: lint
lint:
	golangci-lint run

.PHONY: fmt
.SILENT: fmt
fmt:
	go fmt ./...

.PHONY: tidy
.SILENT: tidy
tidy:
	go mod tidy

.PHONY: vet
.SILENT: vet
vet:
	go vet ./...

.PHONY: test 
.SILENT: test
test:
	go test -timeout 1m ./... -cover

.PHONY: cover
.SILENT: cover
cover:
	mkdir -p .tmp
	go test -timeout 1m ./... -coverprofile=.tmp/coverage.out
	go tool cover -html=.tmp/coverage.out	

.PHONY: generate-docs
.SILENT: generate-docs
generate-docs:
	rm generated-docs/*.md
	go run ./generated-docs/generate-docs.go

.PHONY: build
.SILENT: build
build:
	VERSION=$$(git describe --exact-match --tags --always --dirty)
	DATE=$$(date '+%Y-%m-%d')
	CGO_ENABLED=0 go build -ldflags "-s -w -X github.com/co-native-ab/pimctl/internal/build.Version=$${VERSION} -X github.com/co-native-ab/pimctl/internal/build.Date=$${DATE}" -o bin/pimctl ./

.PHONY: update-openapi-specifications
.SILENT: update-openapi-specifications
update-openapi-specifications:
	curl --silent --fail -o ./openapi/msgraph_openapi_v1.0.yaml -L https://aka.ms/graph/v1.0/openapi.yaml
	curl --silent --fail -o ./openapi/msgraph_openapi_beta.yaml -L https://aka.ms/graph/beta/openapi.yaml

.PHONY: generate-sdks
.SILENT: generate-sdks
generate-sdks:
	kiota generate --clean-output --openapi openapi/msgraph_openapi_v1.0.yaml --language Go --class-name msgraphsdk --namespace-name github.com/co-native-ab/pimctl/internal/generated/msgraphsdk --output ./internal/generated/msgraphsdk \
		--include-path /me \
		--include-path /roleManagement/directory/roleEligibilitySchedules/* \
		--include-path /roleManagement/directory/roleAssignmentScheduleInstances/* \
		--include-path /roleManagement/directory/roleAssignmentScheduleRequests \
		--include-path /roleManagement/directory/roleAssignmentScheduleRequests/* \
		--include-path /policies/roleManagementPolicyAssignments \
		--include-path /identityGovernance/privilegedAccess/group/eligibilitySchedules/* \
		--include-path /identityGovernance/privilegedAccess/group/assignmentScheduleInstances/* \
		--include-path /identityGovernance/privilegedAccess/group/assignmentScheduleRequests \
		--include-path /identityGovernance/privilegedAccess/group/assignmentScheduleRequests/* \
		--include-path /identityGovernance/privilegedAccess/group/assignmentApprovals/**
	kiota generate --clean-output --openapi openapi/msgraph_openapi_beta.yaml --language Go --class-name msgraphbetasdk --namespace-name github.com/co-native-ab/pimctl/internal/generated/msgraphbetasdk --output ./internal/generated/msgraphbetasdk \
		--include-path /roleManagement/directory/roleAssignmentApprovals/**

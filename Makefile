UNAME = $(shell uname -s)
JAVAOPT = '-Dio.swagger.parser.util.RemoteUrl.trustAll=true -Dio.swagger.v3.parser.util.RemoteUrl.trustAll=true'
OUTPUTLOCATION = /local/
ifeq ($(UNAME), Darwin)
	#JAVAOPT = '-Djavax.net.ssl.trustStore=/local/cacerts.jks -Djavax.net.ssl.trustStorePassword=password'
	OUTPUTLOCATION = /local/openapi/
endif
ifndef OPENAPIURL
	OPENAPIURL = https://api-dev-ext.contabo.intra/cloud-features/cloud-features-api.yaml
endif

ifndef OPENAPIVOLUME
	OPENAPIVOLUME = "${PWD}:/local"
endif
.PHONY: build
build: generate-api-clients build-only unittest

.PHONY: generate-api-clients
generate-api-clients:
	rm -rf openapi
	docker run --rm -v $(OPENAPIVOLUME) --env JAVA_OPTS=$(JAVAOPT) openapitools/openapi-generator-cli:v5.2.1 generate \
	--skip-validate-spec \
	--input-spec $(OPENAPIURL) \
	--generator-name  go \
	--output $(OUTPUTLOCATION)
	docker container create --name dummy -v openapivolume:/openapi gitlab.contabo.intra:5050/arcus/common/images/minimal-image/buster-minimal bash
	docker cp dummy:/openapi/ .
	docker rm dummy

.PHONY: build-only
build-only:
	@echo $(OPENAPIURL)
	go mod tidy
	go mod download
	export VERSION=$$(git rev-list --tags --max-count=1 | xargs -I {} git describe --tags {}); export COMMIT=$$(git rev-parse HEAD); export TIMESTAMP=$$(date -u +"%Y-%m-%dT%H:%M:%SZ"); go build -ldflags="-w -s -X \"contabo.com/cli/cntb/cmd.version=$$VERSION\" -X \"contabo.com/cli/cntb/cmd.commit=$$COMMIT\" -X \"contabo.com/cli/cntb/cmd.date=$$TIMESTAMP\""

.PHONY: unittest
unittest:
	go test ./...

.PHONY: bats
bats: build bats-only

.PHONY: bats-only
bats-only:
	rm -f ~/.cache/cntb/token
	bats -rt --timing  bats/*.bats

.PHONY: install
install: bats
	go install

.PHONY: release
release: build
	go get github.com/mitchellh/gox
	mkdir -p dist/
	rm -rf dist/*
	export VERSION=$$(git rev-list --tags --max-count=1 | xargs -I {} git describe --tags {}); export COMMIT=$$(git rev-parse HEAD); export TIMESTAMP=$$(date -u +"%Y-%m-%dT%H:%M:%SZ"); gox -osarch="darwin/amd64 linux/amd64 linux/arm64 windows/amd64 darwin/arm64" -ldflags="-w -s -X \"contabo.com/cli/cntb/cmd.version=$$VERSION\" -X \"contabo.com/cli/cntb/cmd.commit=$$COMMIT\" -X \"contabo.com/cli/cntb/cmd.date=$$TIMESTAMP\"" -output="dist/{{.OS}}_{{.Arch}}/{{.Dir}}"
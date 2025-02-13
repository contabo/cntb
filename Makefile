MAKEFLAGS := --jobs=1
JAVAOPT = '-Dio.swagger.parser.util.RemoteUrl.trustAll=true -Dio.swagger.v3.parser.util.RemoteUrl.trustAll=true'
ifndef OUTPUTLOCATION
	OUTPUTLOCATION = /local/
endif
ifndef OPENAPIURL
	OPENAPIURL = https://api.contabo.com/api-v1.yaml
endif

# ifndef OPENAPIVOLUME
# 	OPENAPIVOLUME = "$(CURDIR):/local"
# endif
ifndef OPENAPIIMAGE
	OPENAPIIMAGE = "openapitools/openapi-generator-cli:v5.2.1"
endif
.PHONY: build
build: generate-api-clients build-only

.PHONY: generate-api-clients
generate-api-clients:
	rm -rf openapi
	docker volume create openapivolume
	docker run --rm -v openapivolume:/local --env JAVA_OPTS=$(JAVAOPT) $(OPENAPIIMAGE) generate \
	--skip-validate-spec \
	--input-spec $(OPENAPIURL) \
	--additional-properties=enumClassPrefix=true \
	--generator-name  go \
	--output $(OUTPUTLOCATION)
	docker container create --name dummy -v openapivolume:/openapi alpine bash
	docker cp dummy:/openapi/ .
	docker rm dummy

.PHONY: build-only
build-only:
	go version
	go mod tidy
	go mod download
	export VERSION=$$(git rev-list --tags --max-count=1 | xargs -I {} git describe --tags {}); export COMMIT=$$(git rev-parse HEAD); export TIMESTAMP=$$(date -u +"%Y-%m-%dT%H:%M:%SZ"); go build -ldflags="-w -s -X \"contabo.com/cli/cntb/cmd.version=$$VERSION\" -X \"contabo.com/cli/cntb/cmd.commit=$$COMMIT\" -X \"contabo.com/cli/cntb/cmd.date=$$TIMESTAMP\""

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

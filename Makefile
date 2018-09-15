.PHONY: build
build:
	@go build -o bin/cli

.PHONY: component
component:
	@bin/cli generate-component --force-odd

.PHONY: combine
combine:
	@bin/cli combine
.PHONY: build
build:
	@go build -o bin/cli

.PHONY: component
component:
	@bin/cli component-key --force-odd

.PHONY: combine
combine:
	@bin/cli combined-key
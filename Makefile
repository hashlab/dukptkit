.PHONY: build
build:
	@go build -o bin/cli

.PHONY: ksi
ksi:
	@bin/cli generate-ksi

.PHONY: component
component:
	@bin/cli generate-component --force-odd

.PHONY: combine
combine:
	@bin/cli combine
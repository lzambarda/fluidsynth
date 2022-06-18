ifeq ($(shell uname),Darwin)
define install_dependencies
	brew install portmidi fluidsynth pkg-config readline
endef
endif

include config/dev.env
export $(shell sed 's/=.*//' ./config/dev.env)

.PHONY: dependencies
dependencies:
	 $(install_dependencies)

.PHONY: lint
lint:
	@go fmt ./...
	@go vet ./...

.PHONY: unit_test
unit_test:
	@go test ./...

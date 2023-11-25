APP=spot

.PHONY: build release-build

build:
	@echo "Building spot..." \
		&& go build -o build/$(APP) main.go


# TODO: release-build rule for common arch & bundle

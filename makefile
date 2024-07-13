# Generating Mocks for Type and Interface changes
mockery:
	rm -vf ./src/mocks/*
	mockery --dir src/interfaces --all --output src/mocks

# Update wire DI
wire:
	go install github.com/google/wire/cmd/wire@latest
	wire ./src/container

test:
	go test ./...

# Removing old coverage file
deleteCoverIfExists deletecoverifexists:
ifneq ("$(wildcard ./cover.out)","")
	rm ./cover.out
endif

# Generate code coverage
cover:
	mkdir -p reports
	make deleteCoverIfExists
	go test $$(go list ./... | grep -v /src/mocks ) -coverprofile ./reports/cover.out
	go tool cover -func reports/cover.out
	go tool cover -html=./reports/cover.out
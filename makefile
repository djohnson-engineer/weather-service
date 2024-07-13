# Generating Mocks for Type and Interface changes
mockery:
	rm -vf ./src/mocks/*
	mockery --dir src/interfaces --all --output src/mocks

# Update wire DI
wire:
	go install github.com/google/wire/cmd/wire@latest
	wire ./src/container
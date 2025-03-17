# Makefile for building and testing coin modules

# List of all coins
COINS = aptos bitcoin elrond ethereum flow helium near oasis polkadot stacks sui tron zil zksync avax cosmos eos filecoin kaspa nervos oracle solana starknet tezos waves

# Base directory
BASE_DIR := $(shell pwd)

# Default target to build and test all coins
all: $(COINS)

# Pattern rule for each coin
$(COINS):
	@echo "Building and testing $@..."
	@cd $(BASE_DIR)/$@ && go mod tidy && go test -v && echo "Build for $@ completed successfully.\n"

# Clean target (optional)
clean:
	@echo "Cleaning build artifacts..."
	@find $(BASE_DIR)/coins -name "*.test" -delete

# Target to build a specific coin
.PHONY: $(COINS) all clean

# Help target
help:
	@echo "Available targets:"
	@echo "  all       - Build and test all coins"
	@echo "  <coinname> - Build and test a specific coin (e.g., 'make bitcoin')"
	@echo "  clean     - Remove build artifacts"
	@echo "  help      - Show this help message"
	@echo ""
	@echo "Available coins: $(COINS)"


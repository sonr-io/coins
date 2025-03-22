# Sonr Go Wallet SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/sonr-io/coins.svg)](https://pkg.go.dev/github.com/sonr-io/coins)
[![License](https://img.shields.io/github/license/sonr/go-wallet-sdk)](https://github.com/sonr-io/coins/blob/main/LICENSE)

The Sonr Go Wallet SDK is a comprehensive solution for building wallet applications with offline transaction capabilities across multiple blockchain networks. It provides a unified interface for account management, transaction creation, and signing across various mainstream public chains.

## 🚀 Features

- **Multi-chain support:** Seamlessly interact with major blockchains.
- **Offline transaction signing:** Ensure security with local signing.
- **Account generation and management:** Derive addresses with ease.
- **Customizable transaction creation:** Flexible parameters for all supported chains.
- **BRC20/Atomical/Runes support:** Full Bitcoin token standard compatibility.
- **Extensible architecture:** Modular design for future blockchain integration.

## 📚 Documentation

For detailed documentation and API references, please refer to the README files located within each blockchain directory under the `coins` folder. Each directory, such as `aptos`, `bitcoin`, and others, contains specific usage instructions and implementation details.

Example:

- [Aptos README](https://github.com/sonr-io/coins/tree/main/aptos)
- [Bitcoin README](https://github.com/sonr-io/coins/tree/main/bitcoin)

## 🌐 Supported Chains

The Sonr Go Wallet SDK supports a wide range of blockchain networks. EVM-compatible chains (e.g., BSC, Polygon,
Arbitrum) and Solana-based chains can seamlessly reuse the same code structure for streamlined integration.

| Blockchain                       | Generate Address | Sign Transaction | Sign Message |
| -------------------------------- | ---------------- | ---------------- | ------------ |
| [Aptos](./aptos/README.md)       | ✅               | ✅               | ✅           |
| [Avax](./avax/README.md)         | ✅               | ✅               | ✅           |
| [Bitcoin](./bitcoin/README.md)   | ✅               | ✅               | ✅           |
| [Cosmos](./cosmos/README.md)     | ✅               | ✅               | ✅           |
| [Ethereum](./ethereum/README.md) | ✅               | ✅               | ✅           |
| [Filecoin](./filecoin/README.md) | ✅               | ✅               | ✅           |
| [Helium](./helium/README.md)     | ✅               | ✅               | ✅           |
| [Polkadot](./polkadot/README.md) | ✅               | ✅               | ✅           |
| [Solana](./solana/README.md)     | ✅               | ✅               | ✅           |
| [Sui](./sui/README.md)           | ✅               | ✅               | ✅           |
| [Tezos](./tezos/README.md)       | ✅               | ✅               | ✅           |

_Note: Bitcoin support includes BRC20, Atomicals, and Runes-related functions, such as deployment, minting, transfer, and trading._

## 🛠️ Architecture

The Sonr Go Wallet SDK follows a modular architecture, comprising the following core components:

1. **`github.com/sonr-io/coins`**: Implements transaction creation and signing for each blockchain.
2. **`github.com/sonr-io/crypto`**: Manages general cryptographic operations and signature algorithms.
3. **`github.com/sonr-io/crypto/util`**: Provides helper utilities for common operations.

This structure allows for easy integration and extension of new blockchains.

## 📦 Installation

To install the Sonr Go Wallet SDK, ensure you have Go 1.22+ installed, then run:

```shell
# Set up Go environment for private repositories
export GOPRIVATE=git.sonr.io
export GONOPROXY=git.sonr.io
export GONOSUMDB=git.sonr.io

# Install SDK components
go get -u github.com/sonr-io/coins/bitcoin
go get -u github.com/sonr-io/crypto
```

## ⚙️ Build and Test

To build and test all blockchain modules, use the provided Makefile:

```shell
# Build and test all coins
make

# Build and test a specific coin
make bitcoin
```

The output will display the build status for each chain. If a module fails, the error message will indicate the issue for further debugging.

## 💬 Feedback and Support

The Sonr Go Wallet SDK is based on the OKX Web3 Go Wallet SDK with enhanced features and optimizations. Each blockchain's specific usage can be found in the corresponding README within the `coins` directory. If you encounter any issues or have suggestions, please submit them through [GitLab Issues](https://github.com/sonr-io/coins/issues), and we will address them promptly.

## 📜 License

The Sonr Go Wallet SDK is open-source software licensed under the [MIT license](LICENSE).

---

This README has been updated to reflect the repository reorganization, with crypto operations moved to github.com/sonr-io/crypto and the main wallet functionality in github.com/sonr-io/coins.

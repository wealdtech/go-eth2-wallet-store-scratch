# go-eth2-wallet-store-scratch

[![Tag](https://img.shields.io/github/tag/wealdtech/go-eth2-wallet-store-scratch.svg)](https://github.com/wealdtech/go-eth2-wallet-store-scratch/releases/)
[![License](https://img.shields.io/github/license/wealdtech/go-eth2-wallet-store-scratch.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/wealdtech/go-eth2-wallet-store-scratch?status.svg)](https://godoc.org/github.com/wealdtech/go-eth2-wallet-store-scratch)
[![Travis CI](https://img.shields.io/travis/wealdtech/go-eth2-wallet-store-scratch.svg)](https://travis-ci.org/wealdtech/go-eth2-wallet-store-scratch)
[![codecov.io](https://img.shields.io/codecov/c/github/wealdtech/go-eth2-wallet-store-scratch.svg)](https://codecov.io/github/wealdtech/go-eth2-wallet-store-scratch)
[![Go Report Card](https://goreportcard.com/badge/github.com/wealdtech/go-eth2-wallet-store-scratch)](https://goreportcard.com/report/github.com/wealdtech/go-eth2-wallet-store-scratch)

Ephemeral in-memory store for the [Ethereum 2 wallet](https://github.com/wealdtech/go-eth2-wallet).

## Table of Contents

- [Install](#install)
- [Usage](#usage)
- [Maintainers](#maintainers)
- [Contribute](#contribute)
- [License](#license)

## Install

`go-eth2-wallet-store-scratch` is a standard Go module which can be installed with:

```sh
go get github.com/wealdtech/go-eth2-wallet-store-scratch
```

## Usage

In normal operation this module should not be used directly.  Instead, it should be configured to be used as part of [go-eth2-wallet](https://github.com/wealdtech/go-eth2-wallet).

The scratch store has no options.

Note that because this store is ephemeral there should be no funds left in the wallet at the end of the runtime, or else they will become inaccessible due to the lack of private key
### Example

```go
package main

import (
	e2wallet "github.com/wealdtech/go-eth2-wallet"
	scratch "github.com/wealdtech/go-eth2-wallet-store-scratch"
)

func main() {
    // Set up and use an ephemeral store
    store := scratch.New()
    e2wallet.UseStore(store)

    // Use e2wallet operations as normal.
}
```

## Maintainers

Jim McDonald: [@mcdee](https://github.com/mcdee).

## Contribute

Contributions welcome. Please check out [the issues](https://github.com/wealdtech/go-eth2-wallet-store-scratch/issues).

## License

[Apache-2.0](LICENSE) © 2019 Weald Technology Trading Ltd

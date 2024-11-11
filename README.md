# Hardline

**Hardline** is a Go library for validating blockchain addresses across multiple networks. Currently, it supports Ethereum and Solana addresses.

## Features

- Validate Ethereum addresses (including checksum validation)
- Validate Solana addresses
- Easy-to-use API
- Lightweight and efficient

## Installation

To install Hardline, run:

```bash
go get github.com/dfb-chain/hardline
```

## Usage 

Here's how to use Hardline in your Go project:


```go
package main

import (
    "fmt"

    "github.com/dfb-chain/hardline"
)

func main() {
    ethAddress := "0x27B1FdB04752BbC536007A920D24ACB045561C26"
    solAddress := "4Nd1mXwoz5Mn386YvyXg3k94KQgjwBWEknhAQrvV7z5V"

    if hardline.IsValidEthereumAddress(ethAddress) {
        fmt.Println("Valid Ethereum address")
    } else {
        fmt.Println("Invalid Ethereum address")
    }

    if hardline.IsValidSolanaAddress(solAddress) {
        fmt.Println("Valid Solana address")
    } else {
        fmt.Println("Invalid Solana address")
    }
}
```

## Documentation 

### IsValidEthereumAddress 


```go
func IsValidEthereumAddress(address string) bool
```
Checks if the provided string is a valid Ethereum address. It supports addresses with or without the `0x` prefix and verifies the checksum if the address is in mixed case.
### IsValidSolanaAddress 


```go
func IsValidSolanaAddress(address string) bool
```

Checks if the provided string is a valid Solana address by decoding it from Base58 and verifying that it is 32 bytes long.

## Contributing 

Contributions are welcome! Please open an issue or submit a pull request on GitHub.

## License 
This project is licensed under the MIT License. See the [LICENSE](LICENSE.md)  file for details.

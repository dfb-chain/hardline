package hardline

import (
	"encoding/hex"
	"strings"

	"github.com/mr-tron/base58"
	"golang.org/x/crypto/sha3"
)

// IsValidEthereumAddress checks if the provided string is a valid Ethereum address.
func IsValidEthereumAddress(address string) bool {
	// Check if the address starts with '0x' or '0X'
	if strings.HasPrefix(address, "0x") || strings.HasPrefix(address, "0X") {
		address = address[2:]
	} else {
		// Return false if '0x' prefix is missing
		return false
	}

	// Check if the address is 40 hexadecimal characters
	if len(address) != 40 {
		return false
	}

	// Check if the address contains only valid hexadecimal characters
	if _, err := hex.DecodeString(address); err != nil {
		return false
	}

	// Check if address is checksummed
	if isMixedCase(address) {
		return isChecksumValid(address)
	}

	// If address is all lowercase or all uppercase, it's valid
	return true
}

// isMixedCase checks if the address contains both uppercase and lowercase letters.
func isMixedCase(address string) bool {
	hasUpper := false
	hasLower := false
	for _, c := range address {
		if c >= 'A' && c <= 'F' {
			hasUpper = true
		} else if c >= 'a' && c <= 'f' {
			hasLower = true
		}
	}
	return hasUpper && hasLower
}

// isChecksumValid verifies the checksum of an Ethereum address.
func isChecksumValid(address string) bool {
	// Compute the Keccak-256 hash of the lowercase address
	addressLower := strings.ToLower(address)
	hasher := sha3.NewLegacyKeccak256()
	hasher.Write([]byte(addressLower))
	hash := hasher.Sum(nil)

	for i := 0; i < len(address); i++ {
		c := address[i]
		// Only letters need to be checked for case
		if (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F') {
			// Get the ith nibble from the hash
			hashNibble := (hash[i/2] >> uint(4*(1-i%2))) & 0x0F

			if (hashNibble >= 8 && c >= 'a' && c <= 'f') || (hashNibble < 8 && c >= 'A' && c <= 'F') {
				return false
			}
		}
	}
	return true
}

// IsValidSolanaAddress checks if the provided string is a valid Solana address.
func IsValidSolanaAddress(address string) bool {
	decoded, err := base58.Decode(address)
	if err != nil {
		return false
	}
	// Solana public keys are 32 bytes
	return len(decoded) == 32
}

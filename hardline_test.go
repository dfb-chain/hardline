package hardline

import "testing"

func TestIsValidEthereumAddress(t *testing.T) {
	validAddresses := []string{
		// All lowercase (no checksum)
		"0xde709f2102306220921060314715629080e2fb77",
		// All uppercase (no checksum)
		"0xDE709F2102306220921060314715629080E2FB77",
		// Checksummed addresses from EIP-55
		"0x27b1fdb04752bbc536007a920d24acb045561c26", // All lowercase, no checksum
		"0x27B1FdB04752BbC536007A920D24ACB045561C26", // Valid checksum
		"0x52908400098527886E0F7030069857D2E4169EE7", // Valid checksum
		"0x8617E340B3D01FA5F11F306F4090FD50E238070D", // Valid checksum
		"0xABCDabcdABCDabcdABCDabcdABCDabcdABCDabcd", // Valid checksum
	}

	invalidAddresses := []string{
		"0x27B1fdb04752BbC536007a920d24acb045561C26", // Invalid checksum
		"0x52908400098527886E0F7030069857D2E4169EEG", // Invalid character 'G'
		"0x12345", // Too short
		"0x27b1fdb04752bbc536007a920d24acb045561c2",  // Too short
		"0x27B1FdB04752BbC536007A920D24ACB045561C2G", // Invalid character
		"0xDE709F2102306220921060314715629080E2FB7G", // Invalid character 'G'
		"0x8617E340B3D01FA5F11F306F4090FD50E238070Z", // Invalid character 'Z'
		"0x",                     // Empty address after '0x'
		"0xGHIJKLmnopqrstuvwxyz", // Invalid characters
		"DE709F2102306220921060314715629080E2FB77", // Missing '0x' prefix
		"",     // Empty string
		"    ", // Spaces only
		"🤔",    // Emoji
	}

	for _, addr := range validAddresses {
		if !IsValidEthereumAddress(addr) {
			t.Errorf("Expected valid Ethereum address: %s", addr)
		}
	}
	for _, addr := range invalidAddresses {
		if IsValidEthereumAddress(addr) {
			t.Errorf("Expected invalid Ethereum address: %s", addr)
		}
	}
}

func TestIsValidSolanaAddress(t *testing.T) {
	validAddresses := []string{
		"4Nd1mXwoz5Mn386YvyXg3k94KQgjwBWEknhAQrvV7z5V",
		"BPFLoaderUpgradeab1e11111111111111111111111",
		"Vote111111111111111111111111111111111111111",
		"11111111111111111111111111111111", // Special case: Solana uses this in some contexts
	}

	invalidAddresses := []string{
		"ThisIsNotAValidAddress!",
		"4Nd1mXwoz5Mn386YvyXg3k94KQgjwBWEknhAQrvV7z5",   // Too short
		"4Nd1mXwoz5Mn386YvyXg3k94KQgjwBWEknhAQrvV7z5VV", // Too long
		"1111111111111111111111111111111O11111111111",   // Invalid character 'O'
		"",      // Empty string
		"     ", // Spaces only
		"🤔",     // Emoji
	}

	for _, addr := range validAddresses {
		if !IsValidSolanaAddress(addr) {
			t.Errorf("Expected valid Solana address: %s", addr)
		}
	}
	for _, addr := range invalidAddresses {
		if IsValidSolanaAddress(addr) {
			t.Errorf("Expected invalid Solana address: %s", addr)
		}
	}
}

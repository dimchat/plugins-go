/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2021 by Moky <albert.moky@gmail.com>
 *
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2021 Albert Moky
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 * ==============================================================================
 */
package mkm

import (
	"strings"

	. "github.com/dimchat/mkm-go/digest"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

// -------------------------------------------------------------------------
//  Ethereum (ETH) Address Implementation
// -------------------------------------------------------------------------

// ETHAddress implements the Address interface using Ethereum-style address format.
//
//	Format Structure:
//	    "0x{40-character hex string}" (case-insensitive with EIP-55 checksum)
//
// Generation Algorithm:
//  1. fingerprint = Public key data (PK.data)
//  2. digest      = KECCAK256(fingerprint)
//  3. address     = "0x" + EIP-55 checksummed hex of digest last 20 bytes
type ETHAddress struct {
	//Address
	ConstantString
}

// NewETHAddress creates a new ETHAddress instance with the given address string
//
// Parameters:
//   - address - EIP-55 compliant ETH address string (0x + 40 hex chars)
//
// Returns: Pointer to initialized ETHAddress instance
func NewETHAddress(address string) *ETHAddress {
	return &ETHAddress{
		ConstantString: *NewConstantString(address),
	}
}

// Override
func (address ETHAddress) Network() EntityType {
	return USER
}

// GenerateETHAddress creates a valid ETHAddress from a public key fingerprint
//
// # Follows Ethereum address generation standard (KECCAK256 hash of public key)
//
// # Implements EIP-55 checksum for case sensitivity validation
//
// Parameters:
//   - fingerprint - Public key data (PK.data, 65 bytes with 0x04 prefix or 64 bytes raw)
//
// Returns: Valid Address interface implementation (ETHAddress)
func GenerateETHAddress(fingerprint []byte) Address {
	if len(fingerprint) == 65 {
		fingerprint = fingerprint[1:]
	}
	// 1. digest = keccak256(fingerprint);
	digest := KECCAK256(fingerprint)
	// 2. address = hex_encode(digest.suffix(20));
	address := "0x" + eip55(HexEncode(digest[32-20:]))
	return NewETHAddress(address)
}

// ParseETHAddress validates and parses a string into an ETHAddress
//
// # Checks format compliance (0x prefix, 42 total characters, valid hex chars)
//
// Parameters:
//   - address - ETH address string to parse (0x + 40 hex chars)
//
// Returns: Valid Address (ETHAddress) if parsing succeeds, nil if invalid
func ParseETHAddress(address string) Address {
	if isETH(address) {
		return NewETHAddress(address)
	}
	return nil
}

// eip55 implements EIP-55 checksum for Ethereum addresses
//
// # Converts lowercase hex string to mixed-case checksum format
//
// Reference: https://eips.ethereum.org/EIPS/eip-55
//
// Parameters:
//   - hex - 40-character lowercase hex string (without 0x prefix)
//
// Returns: EIP-55 checksummed 40-character hex string
func eip55(hex string) string {
	sb := make([]byte, 40)
	utf8 := UTF8Encode(hex)
	hash := KECCAK256(utf8)
	var ch byte
	var i uint8
	for i = 0; i < 40; i++ {
		ch = utf8[i]
		if ch > '9' {
			// check for each 4 bits in the hash table
			// if the first bit is '1',
			//     change the character to uppercase
			ch -= (hash[i>>1] << (i << 2 & 4) & 0x80) >> 2
		}
		sb[i] = ch
	}
	return UTF8Decode(sb)
}

// isETH validates basic Ethereum address format
//
// # Checks: 42 characters total, 0x prefix, valid hex characters (0-9, A-F, a-f)
//
// Parameters:
//   - address - ETH address string to validate
//
// Returns: true if address has valid basic format, false otherwise
func isETH(address string) bool {
	if len(address) != 42 {
		return false
	}
	if address[0] != '0' || address[1] != 'x' {
		return false
	}
	var ch byte
	for i := 2; i < 42; i++ {
		ch = address[i]
		if ch >= '0' && ch <= '9' {
			continue
		}
		if ch >= 'A' && ch <= 'Z' {
			continue
		}
		if ch >= 'a' && ch <= 'z' {
			continue
		}
		// unexpected character
		return false
	}
	return true
}

// GetValidateETHAddressString returns EIP-55 checksummed address from a valid basic ETH address
//
// # Converts valid address to lowercase then applies EIP-55 checksum
//
// Parameters:
//   - address - Valid basic ETH address string (0x + 40 hex chars)
//
// Returns: EIP-55 checksummed address string, empty string if input is invalid
func GetValidateETHAddressString(address string) string {
	if isETH(address) {
		lower := strings.ToLower(address[2:])
		return "0x" + eip55(lower)
	}
	return ""
}

// IsValidateETHAddressString checks if an address string is EIP-55 checksum compliant
//
// # Verifies both basic format and correct case checksum
//
// Parameters:
//   - address - ETH address string to validate
//
// Returns: true if address is EIP-55 compliant, false otherwise
func IsValidateETHAddressString(address string) bool {
	validate := GetValidateETHAddressString(address)
	return validate == address
}

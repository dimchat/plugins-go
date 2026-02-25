/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2020 by Moky <albert.moky@gmail.com>
 *
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2020 Albert Moky
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
	. "github.com/dimchat/mkm-go/digest"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/types"
)

// -------------------------------------------------------------------------
//  Bitcoin (BTC) Address Implementation
// -------------------------------------------------------------------------

// BTCAddress implements the Address interface using Bitcoin-style address format.
//
//	Format Structure (base58 encoded): "network+digest+checksum"
//	    network  :  1 byte
//	    digest   : 20 bytes
//	    checksum :  4 bytes
//
// Generation Algorithm:
//  1. fingerprint = sign(seed, SK)
//  2. digest      = RIPEMD160(SHA256(fingerprint))
//  3. checksum    = SHA256(SHA256(network + digest))[:4]
//  4. address     = Base58Encode(network + digest + checksum)
type BTCAddress struct {
	//Address
	ConstantString

	// network identifies the blockchain/entity type for this address
	network EntityType
}

// NewBTCAddress creates a new BTCAddress instance with the given address string and network type
//
// Parameters:
//   - address - Base58-encoded BTC address string
//   - network - EntityType (blockchain network identifier)
//
// Returns: Pointer to initialized BTCAddress instance
func NewBTCAddress(address string, network EntityType) *BTCAddress {
	return &BTCAddress{
		ConstantString: *NewConstantString(address),
		network:        network,
	}
}

// Override
func (address BTCAddress) Network() EntityType {
	return address.network
}

// GenerateBTCAddress creates a valid BTCAddress from a fingerprint and network type
//
// # Follows standard Bitcoin address generation algorithm with double hashing and checksum
//
// Parameters:
//   - fingerprint - Meta.fingerprint or PublicKey.data
//   - network     - EntityType (blockchain network identifier)
//
// Returns: Valid Address interface implementation (BTCAddress)
func GenerateBTCAddress(fingerprint []byte, network EntityType) Address {
	// 1. digest = ripemd160(sha256(fingerprint))
	digest := RIPEMD160(SHA256(fingerprint))
	// 2. head = network + digest
	head := make([]byte, 21)
	head[0] = uint8(network)
	BytesCopy(digest, 0, head, 1, 20)
	// 3. cc = sha256(sha256(head)).prefix(4)
	cc := checkCode(head)
	// 4. data = base58_encode(head + cc)
	data := make([]byte, 25)
	BytesCopy(head, 0, data, 0, 21)
	BytesCopy(cc, 0, data, 21, 4)
	base58 := Base58Encode(data)
	return NewBTCAddress(base58, network)
}

// ParseBTCAddress validates and parses a Base58 string into a BTCAddress
//
// # Performs length validation and checksum verification before creating address
//
// Parameters:
//   - base58 - Base58-encoded BTC address string to parse
//
// Returns: Valid Address (BTCAddress) if parsing succeeds, nil if invalid
func ParseBTCAddress(base58 string) Address {
	// decode
	data := Base58Decode(base58)
	if len(data) != 25 {
		//panic("address length error")
		return nil
	}
	// CheckCode
	prefix := make([]byte, 21)
	suffix := make([]byte, 4)
	BytesCopy(data, 0, prefix, 0, 21)
	BytesCopy(data, 21, suffix, 0, 4)
	cc := checkCode(prefix)
	// verify
	if BytesEqual(cc, suffix) {
		network := EntityType(data[0])
		return NewBTCAddress(base58, network)
	}
	//panic("address check code error")
	return nil
}

// checkCode computes the 4-byte checksum for BTC address validation
//
// # Implements double SHA256 hashing (SHA256(SHA256(data))) and returns first 4 bytes
//
// Parameters:
//   - data - Byte slice to compute checksum for (network+digest for BTC addresses)
//
// Returns: 4-byte checksum slice
func checkCode(data []byte) []byte {
	sha256d := SHA256(SHA256(data))
	cc := make([]byte, 4)
	BytesCopy(sha256d, 0, cc, 0, 4)
	return cc
}

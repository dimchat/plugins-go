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
	. "github.com/dimchat/core-go/mkm"
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

// DefaultMeta is the standard Meta implementation for generating addresses for IDs
//
// # Serves as the base implementation for address generation using MKM scheme
//
// Version: 1 (MKM)
//
// Address Generation Algorithm (BTC-compatible):
//  1. fingerprint = sign(seed, SK)
//  2. digest      = RIPEMD160(SHA256(fingerprint))
//  3. checksum    = SHA256(SHA256(network + digest))[:4]
//  4. address     = Base58Encode(network + digest + checksum)
type DefaultMeta struct {
	*BaseMeta

	// addresses caches generated Address instances by EntityType (network identifier)
	addresses map[EntityType]Address
}

func NewDefaultMeta(dict StringKeyMap,
	version MetaType, key VerifyKey, seed string, fingerprint TransportableData,
) *DefaultMeta {
	return &DefaultMeta{
		BaseMeta:  NewBaseMeta(dict, version, key, seed, fingerprint),
		addresses: make(map[EntityType]Address, 1),
	}
}

// Override
func (meta *DefaultMeta) GenerateAddress(network EntityType) Address {
	// check caches
	address := meta.addresses[network]
	if address == nil {
		// generate and cache it
		fingerprint := meta.Fingerprint()
		address = GenerateBTCAddress(fingerprint.Bytes(), network)
		meta.addresses[network] = address
	}
	return address
}

// BTCMeta is a specialized Meta implementation for generating BTC-compatible addresses for IDs
//
// # Optimized for Bitcoin address format generation using public key material directly
//
// Version: 2 (BTC)
//
// Address Generation Algorithm:
//  1. CT          = Raw public key data (key.data)
//  2. digest      = RIPEMD160(SHA256(CT))
//  3. checksum    = SHA256(SHA256(network + digest))[:4]
//  4. address     = Base58Encode(network + digest + checksum)
type BTCMeta struct {
	*BaseMeta

	// addresses caches generated BTC Address instances by EntityType (network identifier)
	//
	// Supports multiple BTC network types (mainnet/testnet) with cached addresses
	addresses map[EntityType]Address
}

func NewBTCMeta(dict StringKeyMap,
	version MetaType, key VerifyKey, seed string, fingerprint TransportableData,
) *BTCMeta {
	return &BTCMeta{
		BaseMeta:  NewBaseMeta(dict, version, key, seed, fingerprint),
		addresses: make(map[EntityType]Address, 1),
	}
}

// Override
func (meta *BTCMeta) GenerateAddress(network EntityType) Address {
	// check caches
	address := meta.addresses[network]
	if address == nil {
		// TODO: compress public key?
		key := meta.PublicKey()
		ted := key.Data()
		// generate and cache it
		address = GenerateBTCAddress(ted.Bytes(), network)
		meta.addresses[network] = address
	}
	return address
}

// ETHMeta is a specialized Meta implementation for generating Ethereum-compatible addresses for IDs
//
// # Optimized for Ethereum address format generation using KECCAK256 hashing
//
// Version: 4 (ETH)
//
// Address Generation Algorithm:
//  1. CT      = Raw public key data without prefix byte (key.data[1:] for 65-byte public keys)
//  2. digest  = KECCAK256(CT)
//  3. address = "0x" + HexEncode(last 20 bytes of digest) (EIP-55 checksum compliant)
type ETHMeta struct {
	*BaseMeta

	// address caches the single generated ETH/ExETH Address instance
	//
	// ETH addresses are network-agnostic (fixed to USER EntityType) so single cache entry suffices
	address Address
}

func NewETHMeta(dict StringKeyMap,
	version MetaType, key VerifyKey, seed string, fingerprint TransportableData,
) *ETHMeta {
	return &ETHMeta{
		BaseMeta: NewBaseMeta(dict, version, key, seed, fingerprint),
		address:  nil,
	}
}

// Override
func (meta *ETHMeta) GenerateAddress(_ EntityType) Address {
	// check caches
	address := meta.address
	if address == nil {
		// generate and cache it
		key := meta.PublicKey()
		ted := key.Data()
		address = GenerateETHAddress(ted.Bytes())
		meta.address = address
	}
	return address
}

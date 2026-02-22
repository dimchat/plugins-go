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

/**
 *  Default Meta to build ID with 'name@address'
 *
 *  version:
 *      1 - MKM
 *
 *  algorithm:
 *      CT      = fingerprint = sKey.sign(seed);
 *      hash    = ripemd160(sha256(CT));
 *      code    = sha256(sha256(network + hash)).prefix(4);
 *      address = base58_encode(network + hash + code);
 */
type DefaultMeta struct {
	*BaseMeta

	// caches
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

/**
 *  Meta to build BTC address for ID
 *
 *  version:
 *      2 - BTC
 *
 *  algorithm:
 *      CT      = key.data;
 *      hash    = ripemd160(sha256(CT));
 *      code    = sha256(sha256(network + hash)).prefix(4);
 *      address = base58_encode(network + hash + code);
 */
type BTCMeta struct {
	*BaseMeta

	// caches
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

/**
 *  Meta to build ETH address for ID
 *
 *  version:
 *      0x04 - ETH
 *      0x05 - ExETH
 *
 *  algorithm:
 *      CT      = key.data;  // without prefix byte
 *      digest  = keccak256(CT);
 *      address = hex_encode(digest.suffix(20));
 */
type ETHMeta struct {
	*BaseMeta

	// cached
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

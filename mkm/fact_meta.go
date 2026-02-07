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
	. "github.com/dimchat/core-go/format"
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/ext"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/mem"
)

/**
 *  Base Meta Factory
 */
type BaseMetaFactory struct {
	//MetaFactory

	_type MetaType
}

func NewMetaFactory(version MetaType) MetaFactory {
	return &BaseMetaFactory{
		_type: version,
	}
}

func (factory BaseMetaFactory) Init(version MetaType) MetaFactory {
	factory._type = version
	return factory
}

// Override
func (factory BaseMetaFactory) GenerateMeta(sKey SignKey, seed string) Meta {
	priKey, ok := sKey.(PrivateKey)
	if !ok {
		//panic("private key error")
		return nil
	}
	pubKey := priKey.PublicKey()
	if pubKey == nil {
		//panic("private key error")
		return nil
	}
	var fingerprint TransportableData
	if seed == "" {
		fingerprint = nil
	} else {
		data := UTF8Encode(seed)
		sig := sKey.Sign(data)
		fingerprint = NewBase64DataWithBytes(sig)
	}
	return factory.CreateMeta(pubKey, seed, fingerprint)
}

// Override
func (factory BaseMetaFactory) CreateMeta(key VerifyKey, seed string, fingerprint TransportableData) Meta {
	version := factory._type
	switch version {
	case MKM:
		return NewMetaWithType(version, key, seed, fingerprint)
	case BTC:
		return NewBTCMetaWithType(version, key, seed, fingerprint)
	case ETH:
		return NewETHMetaWithType(version, key, seed, fingerprint)
	default:
		return nil
	}
}

// Override
func (factory BaseMetaFactory) ParseMeta(info StringKeyMap) Meta {
	// check 'type', 'key', 'seed', 'fingerprint'
	if !ContainsKey(info, "type") || !ContainsKey(info, "key") {
		// meta.type should not be empty
		// meta.key should not be empty
		return nil
	} else if !ContainsKey(info, "seed") {
		if ContainsKey(info, "fingerprint") {
			//panic("meta error")
			return nil
		}
	} else if !ContainsKey(info, "fingerprint") {
		//panic("meta error")
		return nil
	}
	// create meta for type
	var out Meta
	helper := GetGeneralAccountHelper()
	version := helper.GetMetaType(info, "")
	switch version {
	case MKM:
		out = NewMetaWithMap(info)
		break
	case BTC:
		out = NewBTCMetaWithMap(info)
		break
	case ETH:
		out = NewETHMetaWithMap(info)
		break
	default:
		break
	}
	if out.IsValid() {
		return out
	}
	//panic("meta error")
	return nil
}

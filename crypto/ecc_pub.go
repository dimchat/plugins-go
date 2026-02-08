/* license: https://mit-license.org
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
package crypto

import (
	. "github.com/dimchat/core-go/format"
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/digest"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
	"github.com/dimchat/plugins-go/crypto/secp256k1"
)

func NewECCPublicKeyWithMap(dict StringKeyMap) PublicKey {
	key := &ECCPublicKey{}
	return key.InitWithMap(dict)
}

/**
 *  ECC Public Key
 *
 *  <blockquote><pre>
 *  keyInfo format: {
 *      "algorithm"    : "ECC",
 *      "curve"        : "secp256k1",
 *      "data"         : "..." // base64_encode()
 *  }
 *  </pre></blockquote>
 */
type ECCPublicKey struct {
	//PublicKey
	BaseKey

	_data TransportableData
}

func (key *ECCPublicKey) InitWithMap(dict StringKeyMap) PublicKey {
	if key.BaseKey.InitWithMap(dict) != nil {
		// lazy load
		key._data = nil
	}
	return key
}

//-------- ICryptographyKey

// Override
func (key *ECCPublicKey) Data() TransportableData {
	ted := key._data
	if ted == nil {
		text := key.GetString("data", "")
		// check for raw data (33/65 bytes)
		size := len(text)
		if size == 66 || size == 130 {
			// Hex format
			bin := HexDecode(text)
			ted = NewPlainDataWithBytes(bin)
		} else {
			// TODO: PEM format?
		}
		key._data = ted
	}
	return ted
}

//-------- IPublicKey

// Override
func (key *ECCPublicKey) Verify(data []byte, signature []byte) bool {
	if len(signature) > 64 {
		signature = secp256k1.SignatureFromDER(signature)
	}
	ted := key.Data()
	pub := ted.Bytes()
	if len(pub) == 65 {
		pub = pub[1:]
	}
	return secp256k1.Verify(pub, SHA256(data), signature)
}

// Override
func (key *ECCPublicKey) MatchSignKey(sKey SignKey) bool {
	return MatchSignKey(sKey, key)
}

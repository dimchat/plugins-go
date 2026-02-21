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
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/digest"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
	"github.com/dimchat/plugins-go/crypto/secp256k1"
)

// generate key
func NewECCPrivateKey() PrivateKey {
	// generate key
	_, pri := secp256k1.Generate()
	ted := NewPlainDataWithBytes(pri)
	txt := HexEncode(pri)
	// build key info
	info := NewMap()
	info["algorithm"] = ECC
	info["data"] = txt
	info["curve"] = "SECP256k1"
	info["digest"] = "SHA256"
	return &ECCPrivateKey{
		Dictionary: NewDictionary(info),
		data:       ted,
		publicKey:  nil,
	}
}

func NewECCPrivateKeyWithMap(dict StringKeyMap) PrivateKey {
	return &ECCPrivateKey{
		Dictionary: NewDictionary(dict),
		// lazy load
		data:      nil,
		publicKey: nil,
	}
}

/**
 *  ECC Private Key
 *
 *  <blockquote><pre>
 *  keyInfo format: {
 *      "algorithm"    : "ECC",
 *      "curve"        : "secp256k1",
 *      "data"         : "..." // base64_encode()
 *  }
 *  </pre></blockquote>
 */
type ECCPrivateKey struct {
	//PrivateKey
	*Dictionary

	data TransportableData

	publicKey PublicKey
}

// Override
func (key *ECCPrivateKey) Equal(other interface{}) bool {
	return privateKeyEqual(key, other)
}

//-------- ICryptographyKey

// Override
func (key *ECCPrivateKey) Algorithm() string {
	info := key.Map()
	return GetKeyAlgorithm(info)
}

// Override
func (key *ECCPrivateKey) Data() TransportableData {
	ted := key.data
	if ted == nil {
		text := key.GetString("data", "")
		size := len(text)
		if size == 64 {
			// check for raw data (32 bytes)
			// Hex format
			bin := HexDecode(text)
			ted = NewPlainDataWithBytes(bin)
		} else {
			// TODO: PEM format?
		}
		key.data = ted
	}
	return ted
}

//-------- IPrivateKey

// Override
func (key *ECCPrivateKey) Sign(data []byte) []byte {
	ted := key.Data()
	sig := secp256k1.Sign(ted.Bytes(), SHA256(data))
	return secp256k1.SignatureToDER(sig)
}

// Override
func (key *ECCPrivateKey) PublicKey() PublicKey {
	publicKey := key.publicKey
	if publicKey == nil {
		ted := key.Data()
		pri := ted.Bytes()
		pub := secp256k1.GetPublicKey(pri)
		txt := "04" + HexEncode(pub)
		// build key info
		info := NewMap()
		info["algorithm"] = ECC
		info["data"] = txt
		info["curve"] = "SECP256k1"
		info["digest"] = "SHA256"
		publicKey = NewECCPublicKeyWithMap(info)
		key.publicKey = publicKey
	}
	return publicKey
}

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
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
)

func NewPlainKey() SymmetricKey {
	dict := NewMap()
	dict["algorithm"] = PLAIN
	return &PlainKey{
		Dictionary: NewDictionary(dict),
		data:       ZeroPlainData(),
	}
}

func NewPlainKeyWithMap(dict StringKeyMap) SymmetricKey {
	return &PlainKey{
		Dictionary: NewDictionary(dict),
		data:       ZeroPlainData(),
	}
}

// PlainKey implements the SymmetricKey interface for unencrypted/broadcast messages
//
// "Null" symmetric key that performs no actual encryption/decryption on message data
//
// Designed for broadcast messages where content should remain in plaintext
type PlainKey struct {
	//SymmetricKey
	*Dictionary

	// data contains placeholder key material (unused for plaintext operations)
	data TransportableData
}

// Override
func (key *PlainKey) Equal(other interface{}) bool {
	return symmetricKeyEqual(key, other)
}

//-------- ICryptographyKey

// Override
func (key *PlainKey) Algorithm() string {
	info := key.Map()
	return GetKeyAlgorithm(info)
}

// Override
func (key *PlainKey) Data() TransportableData {
	return key.data
}

//-------- ISymmetricKey

// Override
func (key *PlainKey) Encrypt(plaintext []byte, _ StringKeyMap) []byte {
	return plaintext
}

// Override
func (key *PlainKey) Decrypt(ciphertext []byte, _ StringKeyMap) []byte {
	return ciphertext
}

// Override
func (key *PlainKey) MatchEncryptKey(pKey EncryptKey) bool {
	return MatchEncryptKey(pKey, key)
}

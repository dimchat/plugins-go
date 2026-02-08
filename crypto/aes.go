/* license: https://mit-license.org
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
package crypto

import (
	"crypto/aes"
	"crypto/cipher"

	. "github.com/dimchat/core-go/format"
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/types"
)

// generate key
func NewAESKey() SymmetricKey {
	pwd := RandomBytes(256 / 8) // 32
	ted := NewBase64DataWithBytes(pwd)
	// build key info
	info := NewMap()
	info["algorithm"] = AES
	info["data"] = ted.Serialize()
	//info["mode"] = "CBC"
	//info["padding"] = "PKCS7"

	key := &AESKey{}
	if key.InitWithMap(info) != nil {
		key._data = ted
	}
	return key
}

func NewAESKeyWithMap(dict StringKeyMap) SymmetricKey {
	key := &AESKey{}
	return key.InitWithMap(dict)
}

/**
 *  AES Key
 *
 *  <blockquote><pre>
 *  keyInfo format: {
 *      "algorithm": "AES",
 *      "keySize"  : 32,                // optional
 *      "data"     : "{BASE64_ENCODE}}" // password data
 *  }
 *  </pre></blockquote>
 */
type AESKey struct {
	//SymmetricKey
	BaseKey

	_data TransportableData
}

func (key *AESKey) InitWithMap(dict StringKeyMap) SymmetricKey {
	if key.BaseKey.InitWithMap(dict) != nil {
		// TODO: check algorithm parameters
		// 1. check mode = 'CBC'
		// 2. check padding = 'PKCS7Padding'

		// lazy load
		key._data = nil
	}
	return key
}

// protected
func (key *AESKey) keySize() uint {
	// TODO: get from key data
	return key.GetUInt("keySize", 256/8) // 32
}

// protected
func (key *AESKey) blockSize() uint {
	// TODO: get from iv data
	return key.GetUInt("blockSize", aes.BlockSize) // 16
}

// Override
func (key *AESKey) Equal(other interface{}) bool {
	return symmetricKeyEqual(key, other)
}

//-------- ICryptographyKey

// Override
func (key *AESKey) Data() TransportableData {
	ted := key._data
	if ted == nil {
		base64 := key.Get("data")
		ted = ParseTransportableData(base64)
		key._data = ted
	}
	return ted
}

// protected
func (key *AESKey) initVector(params StringKeyMap) []byte {
	// get base64 encoded IV from params
	var base64 interface{}
	if params != nil {
		base64 = params["IV"]
		if base64 == nil {
			base64 = params["iv"]
		}
	}
	if base64 == nil {
		// compatible with old version
		base64 = key.Get("iv")
		if base64 == nil {
			base64 = key.Get("IV")
		}
	}
	// decode IV data
	iv := ParseTransportableData(base64)
	if iv == nil || iv.IsEmpty() {
		return nil
	}
	return iv.Bytes()
}

// protected
func (key *AESKey) zeroInitVector() []byte {
	// zero IV
	blockSize := key.blockSize()
	return make([]byte, blockSize)
}

// protected
func (key *AESKey) newInitVector(extra StringKeyMap) []byte {
	// random IV data
	blockSize := key.blockSize()
	iv := RandomBytes(blockSize)
	// pub encoded IV into extra
	if extra != nil {
		ted := NewBase64DataWithBytes(iv)
		extra["IV"] = ted.Serialize()
	}
	// OK
	return iv
}

//-------- ISymmetricKey

// Override
func (key *AESKey) Encrypt(plaintext []byte, extra StringKeyMap) []byte {
	// 1. if 'IV' not found in extra params, new a random 'IV'
	iv := key.initVector(extra)
	if iv == nil {
		iv = key.newInitVector(extra)
	}
	// 2. get key data
	data := key.Data()
	block, err := aes.NewCipher(data.Bytes())
	if err != nil {
		panic(err)
	}
	// 3. try to encrypt
	blockMode := cipher.NewCBCEncrypter(block, iv)
	padded := PKCS5Padding(plaintext, key.blockSize())
	ciphertext := make([]byte, len(padded))
	blockMode.CryptBlocks(ciphertext, padded)
	return ciphertext
}

// Override
func (key *AESKey) Decrypt(ciphertext []byte, params StringKeyMap) []byte {
	// 1. if 'IV' not found in extra params, use an empty 'IV'
	iv := key.initVector(params)
	if iv == nil {
		iv = key.zeroInitVector()
	}
	// 2. get key data
	data := key.Data()
	block, err := aes.NewCipher(data.Bytes())
	if err != nil {
		panic(err)
	}
	// 3. try to decrypt
	blockMode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(plaintext, ciphertext)
	return PKCS5UnPadding(plaintext)
}

// Override
func (key *AESKey) MatchEncryptKey(pKey EncryptKey) bool {
	return MatchEncryptKey(pKey, key)
}

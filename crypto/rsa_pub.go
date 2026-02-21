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
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"

	. "github.com/dimchat/core-go/format"
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/types"
)

type IRSAPublicKey interface {
	PublicKey
	EncryptKey
}

func NewRSAPublicKeyWithMap(dict StringKeyMap) *RSAPublicKey {
	return &RSAPublicKey{
		Dictionary: NewDictionary(dict),
		// lazy load
		rsaPublicKey: nil,
		data:         nil,
	}
}

/**
 *  RSA Public Key
 *
 *  <blockquote><pre>
 *  keyInfo format: {
 *      "algorithm" : "RSA",
 *      "data"      : "..." // base64_encode()
 *  }
 *  </pre></blockquote>
 */
type RSAPublicKey struct {
	//PublicKey, EncryptKey
	*Dictionary

	rsaPublicKey *rsa.PublicKey

	data TransportableData
}

func (key *RSAPublicKey) getPublicKey() *rsa.PublicKey {
	if key.rsaPublicKey == nil {
		text := key.GetString("data", "")
		block, _ := pem.Decode(UTF8Encode(text))
		pub, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			panic(err)
		}
		key.rsaPublicKey, _ = pub.(*rsa.PublicKey)
	}
	return key.rsaPublicKey
}

func (key *RSAPublicKey) getHash() crypto.Hash {
	return crypto.SHA256
}

//-------- ICryptographyKey

// Override
func (key *RSAPublicKey) Algorithm() string {
	info := key.Map()
	return GetKeyAlgorithm(info)
}

// Override
func (key *RSAPublicKey) Data() TransportableData {
	ted := key.data
	if ted == nil {
		// TODO: encode public key data to PKCS1
		pub := key.getPublicKey()
		bin := pub.N.Bytes()
		ted = NewPlainDataWithBytes(bin)
		key.data = ted
	}
	return ted
}

//-------- IPublicKey

// Override
func (key *RSAPublicKey) Verify(data []byte, signature []byte) bool {
	pub := key.getPublicKey()
	h := key.getHash().New()
	h.Write(data)
	sum := h.Sum(nil)
	err := rsa.VerifyPKCS1v15(pub, key.getHash(), sum, signature)
	return err == nil
}

// Override
func (key *RSAPublicKey) MatchSignKey(sKey SignKey) bool {
	return MatchSignKey(sKey, key)
}

//-------- IEncryptKey

// Override
func (key *RSAPublicKey) Encrypt(plaintext []byte, _ StringKeyMap) []byte {
	pub := key.getPublicKey()
	part := pub.N.BitLen()/8 - 11
	chunks := BytesSplit(plaintext, part)
	buffer := bytes.NewBufferString("")
	for _, line := range chunks {
		data, err := rsa.EncryptPKCS1v15(rand.Reader, pub, line)
		if err != nil {
			panic(err)
		}
		buffer.Write(data)
	}
	return buffer.Bytes()
}

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
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/types"
)

type IRSAPrivateKey interface {
	PrivateKey
	DecryptKey
}

// generate key
func NewRSAPrivateKey() IRSAPrivateKey {
	pri, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	der := MarshalPKCS8PrivateKey(pri)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: der,
	}
	bin := pem.EncodeToMemory(block)
	txt := UTF8Decode(bin)
	// build key info
	info := NewMap()
	info["algorithm"] = RSA
	info["data"] = txt
	info["mode"] = "ECB"
	info["padding"] = "PKCS1"
	info["digest"] = "SHA256"

	key := &RSAPrivateKey{}
	if key.InitWithMap(info) != nil {
		key._rsaPrivateKey = pri
	}
	return key
}

func NewRSAPrivateKeyWithMap(dict StringKeyMap) IRSAPrivateKey {
	key := &RSAPrivateKey{}
	return key.InitWithMap(dict)
}

/**
 *  RSA Private Key
 *
 *      keyInfo format: {
 *          algorithm    : "RSA",
 *          data         : "..." // base64_encode()
 *      }
 */
type RSAPrivateKey struct {
	BaseKey

	_rsaPrivateKey *rsa.PrivateKey

	_data TransportableData

	_publicKey PublicKey
}

func (key *RSAPrivateKey) InitWithMap(dict StringKeyMap) IRSAPrivateKey {
	if key.BaseKey.InitWithMap(dict) != nil {
		// lazy load
		key._rsaPrivateKey = nil
		key._data = nil
		key._publicKey = nil
	}
	return key
}

// Override
func (key *RSAPrivateKey) Equal(other interface{}) bool {
	return privateKeyEqual(key, other)
}

func (key *RSAPrivateKey) getPrivateKey() *rsa.PrivateKey {
	if key._rsaPrivateKey == nil {
		text := key.GetString("data", "")
		size := len(text)
		if size == 0 {
			panic("key data not found")
			return nil
		}
		block, _ := pem.Decode(UTF8Encode(text))
		pri, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			pkcs8, err := x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				panic(err)
			}
			pri, _ = pkcs8.(*rsa.PrivateKey)
		}
		key._rsaPrivateKey = pri
	}
	return key._rsaPrivateKey
}

func (key *RSAPrivateKey) getHash() crypto.Hash {
	return crypto.SHA256
}

//-------- ICryptographyKey

// Override
func (key *RSAPrivateKey) Data() TransportableData {
	ted := key._data
	if ted == nil {
		// TODO: encode private key data to PKCS1
		pri := key.getPrivateKey()
		bin := pri.D.Bytes()
		ted := NewPlainDataWithBytes(bin)
		key._data = ted
	}
	return ted
}

//-------- IPrivateKey)

// Override
func (key *RSAPrivateKey) Sign(data []byte) []byte {
	pri := key.getPrivateKey()
	h := key.getHash().New()
	h.Write(data)
	sum := h.Sum(nil)
	sig, err := rsa.SignPKCS1v15(rand.Reader, pri, key.getHash(), sum)
	if err != nil {
		panic(err)
	}
	return sig
}

// Override
func (key *RSAPrivateKey) PublicKey() PublicKey {
	if key._publicKey == nil {
		sKey := key.getPrivateKey()
		pKey := &sKey.PublicKey
		der, err := x509.MarshalPKIXPublicKey(pKey)
		if err != nil {
			panic(err)
		}
		block := &pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: der,
		}
		bin := pem.EncodeToMemory(block)
		txt := UTF8Decode(bin)
		// build key info
		info := NewMap()
		info["algorithm"] = RSA
		info["data"] = txt
		info["mode"] = "ECB"
		info["padding"] = "PKCS1"
		info["digest"] = "SHA256"

		newKey := &RSAPublicKey{}
		if newKey.InitWithMap(info) != nil {
			newKey._rsaPublicKey = pKey
		}
		key._publicKey = newKey
	}
	return key._publicKey
}

//-------- IDecryptKey

// Override
func (key *RSAPrivateKey) Decrypt(ciphertext []byte, _ StringKeyMap) []byte {
	pri := key.getPrivateKey()
	part := pri.N.BitLen() / 8
	chunks := BytesSplit(ciphertext, part)
	buffer := bytes.NewBufferString("")
	for _, line := range chunks {
		data, err := rsa.DecryptPKCS1v15(rand.Reader, pri, line)
		if err != nil {
			panic(err)
		}
		buffer.Write(data)
	}
	return buffer.Bytes()
}

// Override
func (key *RSAPrivateKey) MatchEncryptKey(pKey EncryptKey) bool {
	return MatchEncryptKey(pKey, key)
}

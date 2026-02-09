/* license: https://mit-license.org
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2026 Albert Moky
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
package ext

import (
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/ext"
	. "github.com/dimchat/mkm-go/types"
)

type ICryptoKeyGeneralFactory interface {
	GeneralCryptoHelper
	SymmetricKeyHelper
	PrivateKeyHelper
	PublicKeyHelper
}

func NewCryptoKeyGeneralFactory() ICryptoKeyGeneralFactory {
	gf := &CryptoKeyGeneralFactory{}
	return gf.Init()
}

/**
 *  CryptographyKey GeneralFactory
 */
type CryptoKeyGeneralFactory struct {
	//ICryptoKeyGeneralFactory

	symmetricKeyFactories map[string]SymmetricKeyFactory
	privateKeyFactories   map[string]PrivateKeyFactory
	publicKeyFactories    map[string]PublicKeyFactory
}

func (gf *CryptoKeyGeneralFactory) Init() ICryptoKeyGeneralFactory {
	gf.symmetricKeyFactories = make(map[string]SymmetricKeyFactory)
	gf.privateKeyFactories = make(map[string]PrivateKeyFactory)
	gf.publicKeyFactories = make(map[string]PublicKeyFactory)
	return gf
}

// Override
func (gf *CryptoKeyGeneralFactory) GetKeyAlgorithm(key StringKeyMap, defaultValue string) string {
	algorithm := key["algorithm"]
	return ConvertString(algorithm, defaultValue)
}

//
//  SymmetricKey Helper
//

// Override
func (gf *CryptoKeyGeneralFactory) SetSymmetricKeyFactory(algorithm string, factory SymmetricKeyFactory) {
	gf.symmetricKeyFactories[algorithm] = factory
}

// Override
func (gf *CryptoKeyGeneralFactory) GetSymmetricKeyFactory(algorithm string) SymmetricKeyFactory {
	return gf.symmetricKeyFactories[algorithm]
}

// Override
func (gf *CryptoKeyGeneralFactory) GenerateSymmetricKey(algorithm string) SymmetricKey {
	factory := gf.GetSymmetricKeyFactory(algorithm)
	return factory.GenerateSymmetricKey()
}

// Override
func (gf *CryptoKeyGeneralFactory) ParseSymmetricKey(key interface{}) SymmetricKey {
	if key == nil {
		return nil
	} else if symmetricKey, ok := key.(SymmetricKey); ok {
		return symmetricKey
	}
	info := FetchMap(key)
	if info == nil {
		//panic("symmetric key error")
		return nil
	}
	algorithm := gf.GetKeyAlgorithm(info, "")
	factory := gf.GetSymmetricKeyFactory(algorithm)
	if factory == nil {
		// unknown algorithm, get default key factory
		factory = gf.GetSymmetricKeyFactory("*") // unknown
		if factory == nil {
			//panic("default symmetric key factory not found")
			return nil
		}
	}
	return factory.ParseSymmetricKey(info)
}

//
//  PrivateKey Helper
//

// Override
func (gf *CryptoKeyGeneralFactory) SetPrivateKeyFactory(algorithm string, factory PrivateKeyFactory) {
	gf.privateKeyFactories[algorithm] = factory
}

// Override
func (gf *CryptoKeyGeneralFactory) GetPrivateKeyFactory(algorithm string) PrivateKeyFactory {
	return gf.privateKeyFactories[algorithm]
}

// Override
func (gf *CryptoKeyGeneralFactory) GeneratePrivateKey(algorithm string) PrivateKey {
	factory := gf.GetPrivateKeyFactory(algorithm)
	return factory.GeneratePrivateKey()
}

// Override
func (gf *CryptoKeyGeneralFactory) ParsePrivateKey(key interface{}) PrivateKey {
	if key == nil {
		return nil
	} else if privateKey, ok := key.(PrivateKey); ok {
		return privateKey
	}
	info := FetchMap(key)
	if info == nil {
		//panic("private key error")
		return nil
	}
	algorithm := gf.GetKeyAlgorithm(info, "")
	factory := gf.GetPrivateKeyFactory(algorithm)
	if factory == nil {
		// unknown algorithm, get default key factory
		factory = gf.GetPrivateKeyFactory("*") // unknown
		if factory == nil {
			//panic("default private key factory not found")
			return nil
		}
	}
	return factory.ParsePrivateKey(info)
}

//
//  PublicKey Helper
//

// Override
func (gf *CryptoKeyGeneralFactory) SetPublicKeyFactory(algorithm string, factory PublicKeyFactory) {
	gf.publicKeyFactories[algorithm] = factory
}

// Override
func (gf *CryptoKeyGeneralFactory) GetPublicKeyFactory(algorithm string) PublicKeyFactory {
	return gf.publicKeyFactories[algorithm]
}

// Override
func (gf *CryptoKeyGeneralFactory) ParsePublicKey(key interface{}) PublicKey {
	if key == nil {
		return nil
	} else if publicKey, ok := key.(PublicKey); ok {
		return publicKey
	}
	info := FetchMap(key)
	if info == nil {
		return nil
	}
	algorithm := gf.GetKeyAlgorithm(info, "")
	factory := gf.GetPublicKeyFactory(algorithm)
	if factory == nil {
		// unknown algorithm, get default key factory
		factory = gf.GetPublicKeyFactory("*") // unknown
		if factory == nil {
			//panic("default public key factory not found")
			return nil
		}
	}
	return factory.ParsePublicKey(info)
}

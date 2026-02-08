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
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/crypto"
	. "github.com/dimchat/plugins-go/mem"
)

//goland:noinspection GoSnakeCaseUsage
const (
	RSA_SHA256    = "SHA256withRSA"
	RSA_ECB_PKCS1 = "RSA/ECB/PKCS1Padding"
)

type rsaPrivateFactory struct {
	//PrivateKeyFactory
}

func (factory rsaPrivateFactory) Init() PrivateKeyFactory {
	return factory
}

// Override
func (factory rsaPrivateFactory) GeneratePrivateKey() PrivateKey {
	return NewRSAPrivateKey()
}

// Override
func (factory rsaPrivateFactory) ParsePrivateKey(key StringKeyMap) PrivateKey {
	// check 'data', 'algorithm'
	if !ContainsKey(key, "data") || !ContainsKey(key, "algorithm") {
		// key.data should not be empty
		// key.algorithm should not be empty
		return nil
	}
	return NewRSAPrivateKeyWithMap(key)
}

type rsaPublicFactory struct {
	//PublicKeyFactory
}

func (factory rsaPublicFactory) Init() PublicKeyFactory {
	return factory
}

// Override
func (factory rsaPublicFactory) ParsePublicKey(key StringKeyMap) PublicKey {
	// check 'data', 'algorithm'
	if !ContainsKey(key, "data") || !ContainsKey(key, "algorithm") {
		// key.data should not be empty
		// key.algorithm should not be empty
		return nil
	}
	return NewRSAPublicKeyWithMap(key)
}

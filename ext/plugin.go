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
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/digest"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/plugins-go/digest"
	. "github.com/dimchat/plugins-go/format"
	. "github.com/dimchat/plugins-go/mkm"
)

type IPluginLoader interface {
	Load()
}

type PluginLoader struct {
	//IPluginLoader
}

/**
 *  Register plugins
 */
func (loader PluginLoader) Load() {

	loader.RegisterCoders()
	loader.RegisterDigesters()

	loader.RegisterSymmetricKeyFactories()
	loader.RegisterAsymmetricKeyFactories()

	loader.RegisterEntityFactories()

}

/**
 *  Data coders
 */
// protected
func (loader PluginLoader) RegisterCoders() {

	// Base-58
	SetBase58Coder(NewBase58Coder())
	// Base-64
	SetBase64Coder(NewBase64Coder())
	// HEX
	SetHexCoder(NewHexCoder())

	// JSON
	SetJSONCoder(NewJSONCoder())

	// UTF-8
	SetUTF8Coder(NewUTF8Coder())

	// TED
	SetTransportableDataFactory(NewTransportableDataFactory())
	// PNF
	SetTransportableFileFactory(NewTransportableFileFactory())

}

/**
 *  Message digesters
 */
// protected
func (loader PluginLoader) RegisterDigesters() {

	// SHA-256
	SetSHA256Digester(NewSHA256Digester())

	//// RipeMD-160
	//SetRIPEMD160Digester(NewRIPEMD160Digester())
	//
	//// Keccak-256
	//SetKECCAK256Digester(NewKECCAK256Digester())

}

/**
 *  Symmetric key parsers
 */
// protected
func (loader PluginLoader) RegisterSymmetricKeyFactories() {

	// AES-256
	factory := &aesFactory{}
	SetSymmetricKeyFactory(AES, factory)
	SetSymmetricKeyFactory(AES_CBC_PKCS7, factory)
	//SetSymmetricKeyFactory("AES/CBC/PKCS7Padding", factory)

	// Plain
	SetSymmetricKeyFactory(PLAIN, &plainFactory{})

}

/**
 *  Asymmetric key parsers
 */
// protected
func (loader PluginLoader) RegisterAsymmetricKeyFactories() {

	// RSA
	rsaPri := &rsaPrivateFactory{}
	SetPrivateKeyFactory(RSA, rsaPri)
	SetPrivateKeyFactory(RSA_SHA256, rsaPri)
	SetPrivateKeyFactory(RSA_ECB_PKCS1, rsaPri)

	rsaPub := &rsaPublicFactory{}
	SetPublicKeyFactory(RSA, rsaPub)
	SetPublicKeyFactory(RSA_SHA256, rsaPub)
	SetPublicKeyFactory(RSA_ECB_PKCS1, rsaPub)

	// ECC
	eccPri := &eccPrivateFactory{}
	SetPrivateKeyFactory(ECC, eccPri)
	SetPrivateKeyFactory(ECDSA_SHA256, eccPri)

	eccPub := &eccPublicFactory{}
	SetPublicKeyFactory(ECC, eccPub)
	SetPublicKeyFactory(ECDSA_SHA256, eccPub)

}

/**
 *  ID, Address, Meta, Document parsers
 */
// protected
func (loader PluginLoader) RegisterEntityFactories() {

	// Address
	SetAddressFactory(&BaseAddressFactory{})

	// ID
	SetIDFactory(&IdentifierFactory{})

	// Meta
	registerMetaFactory(MKM)
	registerMetaFactory(BTC)
	registerMetaFactory(ETH)

	// Document
	registerDocumentFactory("*")
	registerDocumentFactory(VISA)
	registerDocumentFactory(PROFILE)
	registerDocumentFactory(BULLETIN)

}

func registerMetaFactory(version string) {
	factory := &BaseMetaFactory{
		Type: version,
	}
	SetMetaFactory(version, factory)
}

func registerDocumentFactory(docType string) {
	factory := &GeneralDocumentFactory{
		Type: docType,
	}
	SetDocumentFactory(docType, factory)
}

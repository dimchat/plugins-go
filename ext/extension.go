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
	. "github.com/dimchat/core-go/dkd"
	. "github.com/dimchat/core-go/ext"
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/dkd-go/ext"
	. "github.com/dimchat/dkd-go/protocol"
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/ext"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/plugins-go/dkd"
	. "github.com/dimchat/plugins-go/msg"
)

type ExtensionLoader struct {
}

/**
 *  Register core factories
 */
func (loader ExtensionLoader) Load() {

	loader.registerCoreHelpers()

	loader.registerMessageFactories()

	loader.registerContentFactories()
	loader.registerCommandFactories()

}

/**
 *  Core extensions
 */
func (loader ExtensionLoader) registerCoreHelpers() {

	// crypto
	cryptoHelper := NewCryptoKeyGeneralFactory()
	SetGeneralCryptoHelper(cryptoHelper)
	SetSymmetricKeyHelper(cryptoHelper)
	SetPrivateKeyHelper(cryptoHelper)
	SetPublicKeyHelper(cryptoHelper)

	// format
	formatHelper := NewFormatGeneralFactory()
	SetTransportableDataHelper(formatHelper)
	SetTransportableFileHelper(formatHelper)

	// mkm
	accountHelper := NewAccountGeneralFactory()
	SetGeneralAccountHelper(accountHelper)
	SetAddressHelper(accountHelper)
	SetIDHelper(accountHelper)
	SetMetaHelper(accountHelper)
	SetDocumentHelper(accountHelper)

	// dkd
	msgHelper := NewMessageGeneralFactory()
	SetGeneralMessageHelper(msgHelper)
	SetContentHelper(msgHelper)
	SetEnvelopeHelper(msgHelper)
	SetInstantMessageHelper(msgHelper)
	SetSecureMessageHelper(msgHelper)
	SetReliableMessageHelper(msgHelper)

	// command
	cmdHelper := NewCommandGeneralFactory()
	SetGeneralCommandHelper(cmdHelper)
	SetCommandHelper(cmdHelper)

}

/**
 *  Message factories
 */
func (loader ExtensionLoader) registerMessageFactories() {

	// envelope
	SetEnvelopeFactory(NewEnvelopeFactory())

	// message factories
	SetInstantMessageFactory(NewInstantMessageFactory())
	SetSecureMessageFactory(NewSecureMessageFactory())
	SetReliableMessageFactory(NewReliableMessageFactory())
}

/**
 *  Core content factories
 */
func (loader ExtensionLoader) registerContentFactories() {
	registerContentFactories()
	registerCustomizedFactories()
}

/**
 *  Core command factories
 */
func (loader ExtensionLoader) registerCommandFactories() {
	registerCommandFactories()
}

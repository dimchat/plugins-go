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
	. "github.com/dimchat/mkm-go/ext"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

type IAccountGeneralFactory interface {
	GeneralAccountHelper
	AddressHelper
	IDHelper
	MetaHelper
	DocumentHelper
}

/**
 *  Account GeneralFactory
 */
type AccountGeneralFactory struct {
	//IAccountGeneralFactory

	addressFactory AddressFactory
	idFactory      IDFactory
	metaFactories  map[string]MetaFactory
	docsFactories  map[string]DocumentFactory
}

func NewAccountGeneralFactory() *AccountGeneralFactory {
	return &AccountGeneralFactory{
		addressFactory: nil,
		idFactory:      nil,
		metaFactories:  make(map[string]MetaFactory, 1024),
		docsFactories:  make(map[string]DocumentFactory, 1024),
	}
}

// Override
func (gf *AccountGeneralFactory) GetMetaType(meta StringKeyMap, defaultValue MetaType) MetaType {
	version := meta["type"]
	return ConvertString(version, defaultValue)
}

// Override
func (gf *AccountGeneralFactory) GetDocumentType(doc StringKeyMap, defaultValue DocumentType) DocumentType {
	text := doc["type"]
	docType := ConvertString(text, "")
	if len(docType) > 0 && docType != "*" {
		return docType
	} else if len(defaultValue) > 0 {
		return defaultValue
	}
	// get type for did
	did := gf.GetDocumentID(doc)
	if did == nil {
		return ""
	} else if did.IsUser() {
		return VISA
	} else if did.IsGroup() {
		return BULLETIN
	}
	return PROFILE
}

// Override
func (gf *AccountGeneralFactory) GetDocumentID(doc StringKeyMap) ID {
	did := doc["did"]
	return ParseID(did)
}

//
//  Address Helper
//

// Override
func (gf *AccountGeneralFactory) SetAddressFactory(factory AddressFactory) {
	gf.addressFactory = factory
}

// Override
func (gf *AccountGeneralFactory) GetAddressFactory() AddressFactory {
	return gf.addressFactory
}

// Override
func (gf *AccountGeneralFactory) ParseAddress(address interface{}) Address {
	if address == nil {
		return nil
	} else if addr, ok := address.(Address); ok {
		return addr
	}
	str := FetchString(address)
	if str == "" {
		//panic("address error")
		return nil
	}
	factory := gf.GetAddressFactory()
	return factory.ParseAddress(str)
}

// Override
func (gf *AccountGeneralFactory) GenerateAddress(meta Meta, network EntityType) Address {
	factory := gf.GetAddressFactory()
	return factory.GenerateAddress(meta, network)
}

//
//  ID Helper
//

// Override
func (gf *AccountGeneralFactory) SetIDFactory(factory IDFactory) {
	gf.idFactory = factory
}

// Override
func (gf *AccountGeneralFactory) GetIDFactory() IDFactory {
	return gf.idFactory
}

// Override
func (gf *AccountGeneralFactory) ParseID(did interface{}) ID {
	if did == nil {
		return nil
	} else if id, ok := did.(ID); ok {
		return id
	}
	str := FetchString(did)
	if str == "" {
		//panic("ID error")
		return nil
	}
	factory := gf.GetIDFactory()
	return factory.ParseID(str)
}

// Override
func (gf *AccountGeneralFactory) CreateID(name string, address Address, terminal string) ID {
	factory := gf.GetIDFactory()
	return factory.CreateID(name, address, terminal)
}

// Override
func (gf *AccountGeneralFactory) GenerateID(meta Meta, network EntityType, terminal string) ID {
	factory := gf.GetIDFactory()
	return factory.GenerateID(meta, network, terminal)
}

//
//  Meta Helper
//

// Override
func (gf *AccountGeneralFactory) SetMetaFactory(version MetaType, factory MetaFactory) {
	gf.metaFactories[version] = factory
}

// Override
func (gf *AccountGeneralFactory) GetMetaFactory(version MetaType) MetaFactory {
	return gf.metaFactories[version]
}

// Override
func (gf *AccountGeneralFactory) CreateMeta(version MetaType, pKey VerifyKey, seed string, fingerprint TransportableData) Meta {
	factory := gf.GetMetaFactory(version)
	return factory.CreateMeta(pKey, seed, fingerprint)
}

// Override
func (gf *AccountGeneralFactory) GenerateMeta(version MetaType, sKey SignKey, seed string) Meta {
	factory := gf.GetMetaFactory(version)
	return factory.GenerateMeta(sKey, seed)
}

// Override
func (gf *AccountGeneralFactory) ParseMeta(meta interface{}) Meta {
	if meta == nil {
		return nil
	} else if m, ok := meta.(Meta); ok {
		return m
	}
	info := FetchMap(meta)
	if info == nil {
		//panic("meta error")
		return nil
	}
	version := gf.GetMetaType(info, "")
	factory := gf.GetMetaFactory(version)
	if factory == nil {
		// unknown meta type, get default meta factory
		factory = gf.GetMetaFactory("*") // unknown
		if factory == nil {
			//panic("default meta factory not found")
			return nil
		}
	}
	return factory.ParseMeta(info)
}

//
//  Document Helper
//

// Override
func (gf *AccountGeneralFactory) SetDocumentFactory(docType DocumentType, factory DocumentFactory) {
	gf.docsFactories[docType] = factory
}

// Override
func (gf *AccountGeneralFactory) GetDocumentFactory(docType DocumentType) DocumentFactory {
	return gf.docsFactories[docType]
}

// Override
func (gf *AccountGeneralFactory) CreateDocument(docType DocumentType, data string, signature TransportableData) Document {
	factory := gf.GetDocumentFactory(docType)
	return factory.CreateDocument(data, signature)
}

// Override
func (gf *AccountGeneralFactory) ParseDocument(doc interface{}) Document {
	if doc == nil {
		return nil
	} else if d, ok := doc.(Document); ok {
		return d
	}
	info := FetchMap(doc)
	if info == nil {
		//panic("document error")
		return nil
	}
	docType := gf.GetDocumentType(info, "")
	factory := gf.GetDocumentFactory(docType)
	if factory == nil {
		// unknown document type, get default document factory
		factory = gf.GetDocumentFactory("*") // unknown
		if factory == nil {
			//panic("default document factory not found")
			return nil
		}
	}
	return factory.ParseDocument(info)
}

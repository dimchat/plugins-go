/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2021 by Moky <albert.moky@gmail.com>
 *
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
package mkm

import (
	. "github.com/dimchat/core-go/mkm"
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/mkm-go/ext"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/mem"
)

/**
 *  General Document Factory
 */
type GeneralDocumentFactory struct {
	//DocumentFactory

	_type string
}

func NewDocumentFactory(docType string) DocumentFactory {
	return &GeneralDocumentFactory{
		_type: docType,
	}
}

func (factory GeneralDocumentFactory) Init(docType string) DocumentFactory {
	factory._type = docType
	return factory
}

// Override
func (factory GeneralDocumentFactory) CreateDocument(data string, signature TransportableData) Document {
	docType := factory._type
	switch docType {
	case VISA:
		return NewVisaWithData(data, signature)
	case BULLETIN:
		return NewBulletinWithData(data, signature)
	default:
		return NewDocumentWithType(docType, data, signature)
	}
}

// Override
func (factory GeneralDocumentFactory) ParseDocument(info StringKeyMap) Document {
	// check 'did', 'data', 'signature'
	if !ContainsKey(info, "data") || !ContainsKey(info, "signature") {
		// doc.data should not be empty
		// doc.signature should not be empty
		return nil
		//} else if !ContainsKey(info, "data") {
		//	// doc.did should not be empty
		//	return nil
	}

	// create document for type
	helper := GetGeneralAccountHelper()
	docType := helper.GetDocumentType(info, "")
	switch docType {
	case VISA:
		return NewVisaWithMap(info)
	case BULLETIN:
		return NewBulletinWithMap(info)
	default:
		return NewDocumentWithMap(info)
	}
}

//
//  Factory methods for Document
//

func NewDocumentWithType(docType string, data string, signature TransportableData) Document {
	doc := &BaseDocument{}
	if data != "" && signature != nil {
		return doc.InitWithType(docType, data, signature)
	} else if data != "" || signature != nil {
		panic("document info error")
	}
	return doc.Init(docType)
}

func NewDocumentWithMap(dict StringKeyMap) Document {
	doc := &BaseDocument{}
	return doc.InitWithMap(dict)
}

func NewVisaWithData(data string, signature TransportableData) Visa {
	doc := &BaseVisa{}
	if data != "" && signature != nil {
		return doc.InitWithData(data, signature)
	} else if data != "" || signature != nil {
		panic("visa info error")
	}
	return doc.Init()
}

func NewVisaWithMap(dict StringKeyMap) Visa {
	doc := &BaseVisa{}
	return doc.InitWithMap(dict)
}

func NewBulletinWithData(data string, signature TransportableData) Bulletin {
	doc := &BaseBulletin{}
	if data != "" && signature != nil {
		return doc.InitWithData(data, signature)
	} else if data != "" || signature != nil {
		panic("bulletin info error")
	}
	return doc.Init()
}

func NewBulletinWithMap(dict StringKeyMap) Bulletin {
	doc := &BaseBulletin{}
	doc.InitWithMap(dict)
	return doc
}

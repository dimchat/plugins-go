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
	"strings"

	. "github.com/dimchat/core-go/rfc"
	. "github.com/dimchat/mkm-go/crypto"
	. "github.com/dimchat/mkm-go/format"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

/**
 *  Format GeneralFactory
 */
type IFormatGeneralFactory interface {
	TransportableFileHelper
	TransportableDataHelper
}

type FormatGeneralFactory struct {
	//IFormatGeneralFactory

	tedFactory TransportableDataFactory
	pnfFactory TransportableFileFactory
}

func NewFormatGeneralFactory() *FormatGeneralFactory {
	return &FormatGeneralFactory{
		tedFactory: nil, // NewTransportableDataFactory(),
		pnfFactory: nil, // NewTransportableFileFactory(),
	}
}

///
///   TED - Transportable Encoded Data
///

// Override
func (gf *FormatGeneralFactory) SetTransportableDataFactory(factory TransportableDataFactory) {
	gf.tedFactory = factory
}

// Override
func (gf *FormatGeneralFactory) GetTransportableDataFactory() TransportableDataFactory {
	return gf.tedFactory
}

// Override
func (gf *FormatGeneralFactory) ParseTransportableData(ted interface{}) TransportableData {
	if ted == nil {
		return nil
	} else if data, ok := ted.(TransportableData); ok {
		return data
	}
	// unwrap
	str := FetchString(ted)
	if str == "" {
		//panic("TED error")
		return nil
	}
	factory := gf.GetTransportableDataFactory()
	return factory.ParseTransportableData(str)
}

///
///   PNF - Portable Network File
///

// Override
func (gf *FormatGeneralFactory) SetTransportableFileFactory(factory TransportableFileFactory) {
	gf.pnfFactory = factory
}

// Override
func (gf *FormatGeneralFactory) GetTransportableFileFactory() TransportableFileFactory {
	return gf.pnfFactory
}

// Override
func (gf *FormatGeneralFactory) CreateTransportableFile(data TransportableData, filename string,
	url URL, password DecryptKey) TransportableFile {
	factory := gf.GetTransportableFileFactory()
	return factory.CreateTransportableFile(data, filename, url, password)
}

// Override
func (gf *FormatGeneralFactory) ParseTransportableFile(pnf interface{}) TransportableFile {
	if pnf == nil {
		return nil
	} else if file, ok := pnf.(TransportableFile); ok {
		return file
	}
	// unwrap
	info := gf.getTransportableFileContent(pnf)
	if info == nil {
		return nil
	}
	factory := gf.GetTransportableFileFactory()
	return factory.ParseTransportableFile(info)
}

// protected
func (gf *FormatGeneralFactory) getTransportableFileContent(pnf interface{}) StringKeyMap {
	info := FetchMap(pnf)
	if info != nil {
		return info
	}
	text := FetchString(pnf)
	if text == "" {
		//panic("PNF error")
		return nil
	} else if text[0] == '{' {
		// decode JSON string
		return JSONDecodeMap(text)
	}
	content := NewMap()

	// 1. check for URL: "http://..."
	pos := strings.Index(text, "://")
	if 0 < pos && pos < 8 {
		content["URL"] = text
		return content
	}

	content["data"] = text
	// 2. check for data URI: "data:image/jpeg;base64,..."
	uri := ParseDataURI(text)
	if uri != nil {
		filename := uri.Filename()
		if filename != "" {
			content["filename"] = filename
		}
	} else {
		// 3. check for Base-64 encoded string?
	}

	return content
}

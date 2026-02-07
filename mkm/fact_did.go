/* license: https://mit-license.org
 *
 *  Ming-Ke-Ming : Decentralized User Identity Authentication
 *
 *                                Written in 2026 by Moky <albert.moky@gmail.com>
 *
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
package mkm

import (
	"strings"

	. "github.com/dimchat/mkm-go/mkm"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/plugins-go/mem"
)

/**
 *  General ID Factory
 */
type IdentifierFactory struct {
	//IDFactory
}

func NewIdentifierFactory() IDFactory {
	return &IdentifierFactory{}
}

func (factory IdentifierFactory) Init() IDFactory {
	return factory
}

// Override
func (factory IdentifierFactory) GenerateID(meta Meta, network EntityType, terminal string) ID {
	address := GenerateAddress(meta, network)
	return CreateID(meta.Seed(), address, terminal)
}

// Override
func (factory IdentifierFactory) CreateID(name string, address Address, terminal string) ID {
	str := IDConcat(name, address, terminal)
	cache := GetIDCache()
	did := cache.Get(str)
	if did == nil {
		did = factory.newID(str, name, address, terminal)
		cache.Put(str, did)
	}
	return did
}

// Override
func (factory IdentifierFactory) ParseID(str string) ID {
	cache := GetIDCache()
	did := cache.Get(str)
	if did == nil {
		did = factory.parse(str)
		if did != nil {
			cache.Put(str, did)
		}
	}
	return did
}

func (factory IdentifierFactory) newID(str string, name string, address Address, terminal string) ID {
	did := &Identifier{}
	return did.Init(str, name, address, terminal)
}

// protected
func (factory IdentifierFactory) parse(identifier string) ID {
	var name string
	var address Address
	var terminal string
	// split ID string
	str := identifier
	pos := strings.IndexByte(str, '/')
	// terminal
	if pos > 0 {
		terminal = str[pos+1:]
		str = str[:pos]
	} else if pos == 0 {
		panic("ID error: " + identifier)
	}
	// name & address
	pos = strings.LastIndexByte(str, '@')
	if pos > 0 {
		name = str[:pos]
		address = ParseAddress(str[pos+1:])
	} else if pos < 0 {
		address = ParseAddress(str)
	} else {
		panic("ID error: " + identifier)
	}
	if address == nil {
		//panic("ID error: " + identifier)
		return nil
	}
	return factory.newID(identifier, name, address, terminal)
}

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
	. "github.com/dimchat/dkd-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

type CreateContent func(StringKeyMap) Content
type CreateCommand func(StringKeyMap) Command

func NewContentFactory(fnCreate CreateContent) ContentFactory {
	return &generalContentFactory{
		fnCreate: fnCreate,
	}
}

func NewCommandFactory(fnCreate CreateCommand) CommandFactory {
	return &generalCommandFactory{
		fnCreate: fnCreate,
	}
}

/**
 *  General Content Factory
 *  ~~~~~~~~~~~~~~~~~~~~~~~
 */
type generalContentFactory struct {
	//ContentFactory

	fnCreate CreateContent
}

func (gf *generalContentFactory) Init(fnCreate CreateContent) ContentFactory {
	gf.fnCreate = fnCreate
	return gf
}

// Override
func (gf *generalContentFactory) ParseContent(content StringKeyMap) Content {
	return gf.fnCreate(content)
}

/**
 *  General Command Factory
 *  ~~~~~~~~~~~~~~~~~~~~~~~
 */
type generalCommandFactory struct {
	//CommandFactory

	fnCreate CreateCommand
}

func (gf *generalCommandFactory) Init(fnCreate CreateCommand) CommandFactory {
	gf.fnCreate = fnCreate
	return gf
}

// Override
func (gf *generalCommandFactory) ParseCommand(content StringKeyMap) Command {
	return gf.fnCreate(content)
}

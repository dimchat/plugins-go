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
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/dkd"
)

func registerCommandFactory(cmd string, fn CreateCommand) {
	factory := NewCommandFactory(fn)
	SetCommandFactory(cmd, factory)
}

func putCommandFactory(cmd string, factory CommandFactory) {
	SetCommandFactory(cmd, factory)
}

func registerCommandFactories() {

	// Meta
	registerCommandFactory(META, func(dict StringKeyMap) Command {
		content := &BaseMetaCommand{}
		content.InitWithMap(dict)
		return content
	})

	// Document
	registerCommandFactory(DOCUMENTS, func(dict StringKeyMap) Command {
		content := &BaseDocumentCommand{}
		content.InitWithMap(dict)
		return content
	})

	// Receipt
	registerCommandFactory(RECEIPT, func(dict StringKeyMap) Command {
		content := &BaseReceiptCommand{}
		content.InitWithMap(dict)
		return content
	})

	// Group Commands
	putCommandFactory("group", &GroupCommandFactory{})

	registerCommandFactory(INVITE, func(dict StringKeyMap) Command {
		content := &InviteGroupCommand{}
		content.InitWithMap(dict)
		return content
	})
	// 'expel' is deprecated (use 'reset' instead)
	registerCommandFactory(EXPEL, func(dict StringKeyMap) Command {
		content := &ExpelGroupCommand{}
		content.InitWithMap(dict)
		return content
	})
	registerCommandFactory(JOIN, func(dict StringKeyMap) Command {
		content := &JoinGroupCommand{}
		content.InitWithMap(dict)
		return content
	})
	registerCommandFactory(QUIT, func(dict StringKeyMap) Command {
		content := &QuitGroupCommand{}
		content.InitWithMap(dict)
		return content
	})
	// 'query' is deprecated
	//registerCommandFactory(QUERY, func(dict StringKeyMap) Command {
	//	content := &QueryGroupCommand{}
	//	content.InitWithMap(dict)
	//	return content
	//})
	registerCommandFactory(RESET, func(dict StringKeyMap) Command {
		content := &ResetGroupCommand{}
		content.InitWithMap(dict)
		return content
	})

}

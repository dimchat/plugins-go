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
	. "github.com/dimchat/plugins-go/dkd"
)

/**
 *  Core command factories
 */
func registerCommandFactories() {

	// Meta
	registerCommandCreator(META, NewMetaCommandWithMap)

	// Document
	registerCommandCreator(DOCUMENTS, NewDocumentCommandWithMap)

	// Receipt
	registerCommandCreator(RECEIPT, NewReceiptCommandWithMap)

	// Group Commands
	SetCommandFactory("group", &GroupCommandFactory{})

	registerCommandCreator(INVITE, NewInviteCommandWithMap)
	// 'expel' is deprecated (use 'reset' instead)
	registerCommandCreator(EXPEL, NewExpelCommandWithMap)
	registerCommandCreator(JOIN, NewJoinCommandWithMap)
	registerCommandCreator(QUIT, NewQuitCommandWithMap)
	// 'query' is deprecated
	//registerCommandCreator(QUERY, NewQueryCommandWithMap)
	registerCommandCreator(RESET, NewResetCommandWithMap)

}

func registerCommandCreator(cmd string, fn FuncCreateCommand) {
	factory := NewCommandFactory(fn)
	SetCommandFactory(cmd, factory)
}

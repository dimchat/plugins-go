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
	. "github.com/dimchat/dkd-go/protocol"
	. "github.com/dimchat/plugins-go/dkd"
)

/**
 *  Core content factories
 */
func registerContentFactories() {

	// Text
	registerContentCreator(ContentType.TEXT, NewTextContentWithMap)

	// File
	registerContentCreator(ContentType.FILE, NewFileContentWithMap)
	// Image
	registerContentCreator(ContentType.IMAGE, NewImageContentWithMap)
	// Audio
	registerContentCreator(ContentType.AUDIO, NewAudioContentWithMap)
	// Video
	registerContentCreator(ContentType.VIDEO, NewVideoContentWithMap)

	// Web Page
	registerContentCreator(ContentType.PAGE, NewPageContentWithMap)

	// Name Card
	registerContentCreator(ContentType.NAME_CARD, NewNameCardWithMap)

	// Quote
	registerContentCreator(ContentType.QUOTE, NewQuoteContentWithMap)

	// Money
	registerContentCreator(ContentType.MONEY, NewMoneyContentWithMap)
	registerContentCreator(ContentType.TRANSFER, NewTransferContentWithMap)
	// ...

	// Command
	SetContentFactory(ContentType.COMMAND, &GeneralCommandFactory{})

	// History Command
	SetContentFactory(ContentType.HISTORY, &HistoryCommandFactory{})

	// Content Array
	registerContentCreator(ContentType.ARRAY, NewArrayContentWithMap)

	// Combine and Forward
	registerContentCreator(ContentType.COMBINE_FORWARD, NewCombineContentWithMap)

	// Top-Secret
	registerContentCreator(ContentType.FORWARD, NewForwardContentWithMap)

	// unknown content type
	registerContentCreator(ContentType.ANY, NewContentWithMap)

}

/**
 *  Customized content factories
 */
func registerCustomizedFactories() {

	// Application Customized Content
	registerContentCreator(ContentType.CUSTOMIZED, NewCustomizedContentWithMap)
	//registerContentCreator(ContentType.APPLICATION, NewCustomizedContentWithMap)

}

func registerContentCreator(msgType string, fn FuncCreateContent) {
	factory := NewContentFactory(fn)
	SetContentFactory(msgType, factory)
}

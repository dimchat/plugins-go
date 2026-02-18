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
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/dkd"
)

/**
 *  Core content factories
 */
func registerContentFactories() {

	// Text
	registerContentCreator(ContentType.TEXT, func(dict StringKeyMap) Content {
		return NewTextContentWithMap(dict)
	})

	// File
	registerContentCreator(ContentType.FILE, func(dict StringKeyMap) Content {
		return NewFileContentWithMap(dict)
	})
	// Image
	registerContentCreator(ContentType.IMAGE, func(dict StringKeyMap) Content {
		return NewImageContentWithMap(dict)
	})
	// Audio
	registerContentCreator(ContentType.AUDIO, func(dict StringKeyMap) Content {
		return NewAudioContentWithMap(dict)
	})
	// Video
	registerContentCreator(ContentType.VIDEO, func(dict StringKeyMap) Content {
		return NewVideoContentWithMap(dict)
	})

	// Web Page
	registerContentCreator(ContentType.PAGE, func(dict StringKeyMap) Content {
		return NewPageContentWithMap(dict)
	})

	// Name Card
	registerContentCreator(ContentType.NAME_CARD, func(dict StringKeyMap) Content {
		return NewNameCardWithMap(dict)
	})

	// Quote
	registerContentCreator(ContentType.QUOTE, func(dict StringKeyMap) Content {
		return NewQuoteContentWithMap(dict)
	})

	// Money
	registerContentCreator(ContentType.MONEY, func(dict StringKeyMap) Content {
		return NewMoneyContentWithMap(dict)
	})
	registerContentCreator(ContentType.TRANSFER, func(dict StringKeyMap) Content {
		return NewTransferContentWithMap(dict)
	})
	// ...

	// Command
	SetContentFactory(ContentType.COMMAND, &GeneralCommandFactory{})

	// History Command
	SetContentFactory(ContentType.HISTORY, &HistoryCommandFactory{})

	// Content Array
	registerContentCreator(ContentType.ARRAY, func(dict StringKeyMap) Content {
		return NewArrayContentWithMap(dict)
	})

	// Combine and Forward
	registerContentCreator(ContentType.COMBINE_FORWARD, func(dict StringKeyMap) Content {
		return NewCombineContentWithMap(dict)
	})

	// Top-Secret
	registerContentCreator(ContentType.FORWARD, func(dict StringKeyMap) Content {
		return NewForwardContentWithMap(dict)
	})

	// unknown content type
	registerContentCreator(ContentType.ANY, func(dict StringKeyMap) Content {
		return NewContentWithMap(dict)
	})

}

/**
 *  Customized content factories
 */
func registerCustomizedFactories() {

	// Application Customized Content
	registerContentCreator(ContentType.CUSTOMIZED, func(dict StringKeyMap) Content {
		return NewCustomizedContentWithMap(dict)
	})
	//registerContentCreator(ContentType.APPLICATION, func(dict StringKeyMap) Content {
	//	return NewCustomizedContentWithMap(dict)
	//})

}

func registerContentCreator(msgType string, fn FuncCreateContent) {
	factory := NewContentFactory(fn)
	SetContentFactory(msgType, factory)
}

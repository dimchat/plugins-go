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
	registerContentCreator(ContentType.TEXT, "text", func(dict StringKeyMap) Content {
		content := &BaseTextContent{}
		content.InitWithMap(dict)
		return content
	})

	// File
	registerContentCreator(ContentType.FILE, "file", func(dict StringKeyMap) Content {
		content := &BaseFileContent{}
		content.InitWithMap(dict)
		return content
	})
	// Image
	registerContentCreator(ContentType.IMAGE, "image", func(dict StringKeyMap) Content {
		content := &ImageFileContent{}
		content.InitWithMap(dict)
		return content
	})
	// Audio
	registerContentCreator(ContentType.AUDIO, "audio", func(dict StringKeyMap) Content {
		content := &AudioFileContent{}
		content.InitWithMap(dict)
		return content
	})
	// Video
	registerContentCreator(ContentType.VIDEO, "video", func(dict StringKeyMap) Content {
		content := &VideoFileContent{}
		content.InitWithMap(dict)
		return content
	})

	// Web Page
	registerContentCreator(ContentType.PAGE, "page", func(dict StringKeyMap) Content {
		content := &WebPageContent{}
		content.InitWithMap(dict)
		return content
	})

	// Name Card
	registerContentCreator(ContentType.NAME_CARD, "card", func(dict StringKeyMap) Content {
		content := &NameCardContent{}
		content.InitWithMap(dict)
		return content
	})

	// Quote
	registerContentCreator(ContentType.QUOTE, "quote", func(dict StringKeyMap) Content {
		content := &BaseQuoteContent{}
		content.InitWithMap(dict)
		return content
	})

	// Money
	registerContentCreator(ContentType.MONEY, "money", func(dict StringKeyMap) Content {
		content := &BaseMoneyContent{}
		content.InitWithMap(dict)
		return content
	})
	registerContentCreator(ContentType.TRANSFER, "transfer", func(dict StringKeyMap) Content {
		content := &TransferMoneyContent{}
		content.InitWithMap(dict)
		return content
	})
	// ...

	// Command
	registerContentFactory(ContentType.COMMAND, "command", &GeneralCommandFactory{})

	// History Command
	registerContentFactory(ContentType.HISTORY, "history", &HistoryCommandFactory{})

	// Content Array
	registerContentCreator(ContentType.ARRAY, "array", func(dict StringKeyMap) Content {
		content := &ListContent{}
		content.InitWithMap(dict)
		return content
	})

	// Combine and Forward
	registerContentCreator(ContentType.COMBINE_FORWARD, "combine", func(dict StringKeyMap) Content {
		content := &CombineForwardContent{}
		content.InitWithMap(dict)
		return content
	})

	// Top-Secret
	registerContentCreator(ContentType.FORWARD, "forward", func(dict StringKeyMap) Content {
		content := &SecretContent{}
		content.InitWithMap(dict)
		return content
	})

	// unknown content type
	registerContentCreator(ContentType.ANY, "*", func(dict StringKeyMap) Content {
		content := &BaseContent{}
		content.InitWithMap(dict)
		return content
	})

}

/**
 *  Customized content factories
 */
func registerCustomizedFactories() {

	// Application Customized Content
	registerContentCreator(ContentType.CUSTOMIZED, "customized", func(dict StringKeyMap) Content {
		content := &AppCustomizedContent{}
		content.InitWithMap(dict)
		return content
	})
	//registerContentCreator(ContentType.APPLICATION, "application", func(dict StringKeyMap) Content {
	//	content := &AppCustomizedContent{}
	//	content.InitWithMap(dict)
	//	return content
	//})

}

func registerContentFactory(msgType, alias string, factory ContentFactory) {
	SetContentFactory(msgType, factory)
	SetContentFactory(alias, factory)
}

func registerContentCreator(msgType, alias string, fn FuncCreateContent) {
	factory := NewContentFactory(fn)
	SetContentFactory(msgType, factory)
	SetContentFactory(alias, factory)
}

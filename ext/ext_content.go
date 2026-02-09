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

func putContentFactory(msgType, alias string, factory ContentFactory) {
	SetContentFactory(msgType, factory)
	SetContentFactory(alias, factory)
}

func registerContentFactory(msgType, alias string, fn CreateContent) {
	factory := NewContentFactory(fn)
	SetContentFactory(msgType, factory)
	SetContentFactory(alias, factory)
}

func registerContentFactories() {

	// Text
	registerContentFactory(ContentType.TEXT, "text", func(dict StringKeyMap) Content {
		content := &BaseTextContent{}
		content.InitWithMap(dict)
		return content
	})

	// File
	registerContentFactory(ContentType.FILE, "file", func(dict StringKeyMap) Content {
		content := &BaseFileContent{}
		content.InitWithMap(dict)
		return content
	})
	// Image
	registerContentFactory(ContentType.IMAGE, "image", func(dict StringKeyMap) Content {
		content := &ImageFileContent{}
		content.InitWithMap(dict)
		return content
	})
	// Audio
	registerContentFactory(ContentType.AUDIO, "audio", func(dict StringKeyMap) Content {
		content := &AudioFileContent{}
		content.InitWithMap(dict)
		return content
	})
	// Video
	registerContentFactory(ContentType.VIDEO, "video", func(dict StringKeyMap) Content {
		content := &VideoFileContent{}
		content.InitWithMap(dict)
		return content
	})

	// Web Page
	registerContentFactory(ContentType.PAGE, "page", func(dict StringKeyMap) Content {
		content := &WebPageContent{}
		content.InitWithMap(dict)
		return content
	})

	// Name Card
	registerContentFactory(ContentType.NAME_CARD, "card", func(dict StringKeyMap) Content {
		content := &NameCardContent{}
		content.InitWithMap(dict)
		return content
	})

	// Quote
	registerContentFactory(ContentType.QUOTE, "quote", func(dict StringKeyMap) Content {
		content := &BaseQuoteContent{}
		content.InitWithMap(dict)
		return content
	})

	// Money
	registerContentFactory(ContentType.MONEY, "money", func(dict StringKeyMap) Content {
		content := &BaseMoneyContent{}
		content.InitWithMap(dict)
		return content
	})
	registerContentFactory(ContentType.TRANSFER, "transfer", func(dict StringKeyMap) Content {
		content := &TransferMoneyContent{}
		content.InitWithMap(dict)
		return content
	})
	// ...

	// Command
	putContentFactory(ContentType.COMMAND, "command", &GeneralCommandFactory{})

	// History Command
	putContentFactory(ContentType.HISTORY, "history", &HistoryCommandFactory{})

	// Content Array
	registerContentFactory(ContentType.ARRAY, "array", func(dict StringKeyMap) Content {
		content := &ListContent{}
		content.InitWithMap(dict)
		return content
	})

	// Combine and Forward
	registerContentFactory(ContentType.COMBINE_FORWARD, "combine", func(dict StringKeyMap) Content {
		content := &CombineForwardContent{}
		content.InitWithMap(dict)
		return content
	})

	// Top-Secret
	registerContentFactory(ContentType.FORWARD, "forward", func(dict StringKeyMap) Content {
		content := &SecretContent{}
		content.InitWithMap(dict)
		return content
	})

	// unknown content type
	registerContentFactory(ContentType.ANY, "*", func(dict StringKeyMap) Content {
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
	registerContentFactory(ContentType.CUSTOMIZED, "customized", func(dict StringKeyMap) Content {
		content := &AppCustomizedContent{}
		content.InitWithMap(dict)
		return content
	})
	//registerContentFactory(ContentType.APPLICATION, "application", func(dict StringKeyMap) Content {
	//	content := &AppCustomizedContent{}
	//	content.InitWithMap(dict)
	//	return content
	//})

}

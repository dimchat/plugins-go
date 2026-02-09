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
	. "github.com/dimchat/dkd-go/ext"
	. "github.com/dimchat/dkd-go/protocol"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

type IMessageGeneralFactory interface {
	GeneralMessageHelper
	ContentHelper
	EnvelopeHelper
	InstantMessageHelper
	SecureMessageHelper
	ReliableMessageHelper
}

func NewMessageGeneralFactory() IMessageGeneralFactory {
	gf := &MessageGeneralFactory{}
	return gf.Init()
}

/**
 *  Message GeneralFactory
 */
type MessageGeneralFactory struct {
	//MessageGeneralFactory

	contentFactories       map[string]ContentFactory
	envelopeFactory        EnvelopeFactory
	instantMessageFactory  InstantMessageFactory
	secureMessageFactory   SecureMessageFactory
	reliableMessageFactory ReliableMessageFactory
}

func (gf *MessageGeneralFactory) Init() IMessageGeneralFactory {
	gf.contentFactories = make(map[string]ContentFactory)
	gf.envelopeFactory = nil
	gf.instantMessageFactory = nil
	gf.secureMessageFactory = nil
	gf.reliableMessageFactory = nil
	return gf
}

// Override
func (gf *MessageGeneralFactory) GetContentType(content StringKeyMap, defaultValue MessageType) MessageType {
	msgType := content["type"]
	return ConvertString(msgType, defaultValue)
}

//
//  Content Helper
//

// Override
func (gf *MessageGeneralFactory) SetContentFactory(msgType MessageType, factory ContentFactory) {
	gf.contentFactories[msgType] = factory
}

// Override
func (gf *MessageGeneralFactory) GetContentFactory(msgType MessageType) ContentFactory {
	return gf.contentFactories[msgType]
}

// Override
func (gf *MessageGeneralFactory) ParseContent(content interface{}) Content {
	if content == nil {
		return nil
	} else if c, ok := content.(Content); ok {
		return c
	}
	info := FetchMap(content)
	if info == nil {
		//panic("content error")
		return nil
	}
	// get factory by content type
	msgType := gf.GetContentType(info, "")
	factory := gf.GetContentFactory(msgType)
	if factory == nil {
		// unknown content type, get default content factory
		factory = gf.GetContentFactory("*") // unknown
		if factory == nil {
			//panic("default content factory not found")
			return nil
		}
	}
	return factory.ParseContent(info)
}

//
//  Envelope Helper
//

// Override
func (gf *MessageGeneralFactory) SetEnvelopeFactory(factory EnvelopeFactory) {
	gf.envelopeFactory = factory
}

// Override
func (gf *MessageGeneralFactory) GetEnvelopeFactory() EnvelopeFactory {
	return gf.envelopeFactory
}

// Override
func (gf *MessageGeneralFactory) CreateEnvelope(from, to ID, when Time) Envelope {
	factory := gf.GetEnvelopeFactory()
	return factory.CreateEnvelope(from, to, when)
}

// Override
func (gf *MessageGeneralFactory) ParseEnvelope(env interface{}) Envelope {
	if env == nil {
		return nil
	} else if envelope, ok := env.(Envelope); ok {
		return envelope
	}
	info := FetchMap(env)
	if info == nil {
		//panic("envelope error")
		return nil
	}
	factory := gf.GetEnvelopeFactory()
	return factory.ParseEnvelope(info)
}

//
//  InstantMessage Helper
//

// Override
func (gf *MessageGeneralFactory) SetInstantMessageFactory(factory InstantMessageFactory) {
	gf.instantMessageFactory = factory
}

// Override
func (gf *MessageGeneralFactory) GetInstantMessageFactory() InstantMessageFactory {
	return gf.instantMessageFactory
}

// Override
func (gf *MessageGeneralFactory) CreateInstantMessage(head Envelope, body Content) InstantMessage {
	factory := gf.GetInstantMessageFactory()
	return factory.CreateInstantMessage(head, body)
}

// Override
func (gf *MessageGeneralFactory) ParseInstantMessage(msg interface{}) InstantMessage {
	if msg == nil {
		return nil
	} else if iMsg, ok := msg.(InstantMessage); ok {
		return iMsg
	}
	info := FetchMap(msg)
	if info == nil {
		//panic("instant message error")
		return nil
	}
	factory := gf.GetInstantMessageFactory()
	return factory.ParseInstantMessage(info)
}

// Override
func (gf *MessageGeneralFactory) GenerateSerialNumber(msgType MessageType, now Time) SerialNumberType {
	factory := gf.GetInstantMessageFactory()
	return factory.GenerateSerialNumber(msgType, now)
}

//
//  SecureMessage Helper
//

// Override
func (gf *MessageGeneralFactory) SetSecureMessageFactory(factory SecureMessageFactory) {
	gf.secureMessageFactory = factory
}

// Override
func (gf *MessageGeneralFactory) GetSecureMessageFactory() SecureMessageFactory {
	return gf.secureMessageFactory
}

// Override
func (gf *MessageGeneralFactory) ParseSecureMessage(msg interface{}) SecureMessage {
	if msg == nil {
		return nil
	} else if sMsg, ok := msg.(SecureMessage); ok {
		return sMsg
	}
	info := FetchMap(msg)
	if info == nil {
		//panic("secure message error")
		return nil
	}
	factory := gf.GetSecureMessageFactory()
	return factory.ParseSecureMessage(info)
}

//
//  ReliableMessage Helper
//

// Override
func (gf *MessageGeneralFactory) SetReliableMessageFactory(factory ReliableMessageFactory) {
	gf.reliableMessageFactory = factory
}

// Override
func (gf *MessageGeneralFactory) GetReliableMessageFactory() ReliableMessageFactory {
	return gf.reliableMessageFactory
}

// Override
func (gf *MessageGeneralFactory) ParseReliableMessage(msg interface{}) ReliableMessage {
	if msg == nil {
		return nil
	} else if rMsg, ok := msg.(ReliableMessage); ok {
		return rMsg
	}
	info := FetchMap(msg)
	if info == nil {
		//panic("reliable message error")
		return nil
	}
	factory := gf.GetReliableMessageFactory()
	return factory.ParseReliableMessage(info)
}

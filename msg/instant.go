/* license: https://mit-license.org
 *
 *  Dao-Ke-Dao: Universal Message Module
 *
 *                                Written in 2022 by Moky <albert.moky@gmail.com>
 *
 * ==============================================================================
 * The MIT License (MIT)
 *
 * Copyright (c) 2022 Albert Moky
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
package dkd

import (
	"math/rand"
	"sync/atomic"

	. "github.com/dimchat/core-go/msg"
	. "github.com/dimchat/dkd-go/protocol"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/mem"
)

type PlainMessageFactory struct {
	//InstantMessageFactory
}

func NewInstantMessageFactory() InstantMessageFactory {
	return &PlainMessageFactory{}
}

// Override
func (factory PlainMessageFactory) CreateInstantMessage(head Envelope, body Content) InstantMessage {
	return NewInstantMessage(head, body)
}

// Override
func (factory PlainMessageFactory) ParseInstantMessage(msg StringKeyMap) InstantMessage {
	// check 'sender', 'content'
	if !ContainsKey(msg, "sender") || !ContainsKey(msg, "content") {
		// msg.sender should not be empty
		// msg.content should not be empty
		return nil
	}
	return NewInstantMessageWithMap(msg)
}

// Override
func (factory PlainMessageFactory) GenerateSerialNumber(_ MessageType, _ Time) SerialNumberType {
	// because we must make sure all messages in a same chat box won't have
	// same serial numbers, so we can't use time-related numbers, therefore
	// the best choice is a totally random number, maybe.
	sn := sharedGenerator.nextSN()
	return SerialNumberType(sn)
}

//
//  Serial Number Generator
//

var sharedGenerator = &snGenerator{
	_sn: rand.Uint32(), // 0 ~ 0x7fffffff
}

type snGenerator struct {
	_sn uint32
}

/**
 *  next sn
 *
 * @return 1 ~ 2^31-1
 */
func (gen *snGenerator) nextSN() uint32 {
	var current uint32
	var next uint32
	var swapped bool
	for {
		current = atomic.LoadUint32(&gen._sn)
		if current < 0x7fffffff {
			next = current + 1
		} else {
			next = 1
		}
		swapped = atomic.CompareAndSwapUint32(&gen._sn, current, next)
		if swapped {
			return next
		}
	}
}

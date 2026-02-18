/* license: https://mit-license.org
 *
 *  Dao-Ke-Dao: Universal Message Module
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
package dkd

import (
	. "github.com/dimchat/core-go/dkd"
	. "github.com/dimchat/core-go/ext"
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/dkd-go/protocol"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/mem"
)

type IGroupCommandFactory interface {
	ContentFactory
	CommandFactory
}

type GroupCommandFactory struct {
	//IGroupCommandFactory
}

func (gf GroupCommandFactory) Init() IGeneralCommandFactory {
	return gf
}

// Override
func (gf GroupCommandFactory) ParseContent(content StringKeyMap) Content {
	helper := GetGeneralCommandHelper()
	cmdHelper := GetCommandHelper()
	// get factory by command name
	cmd := helper.GetCMD(content, "")
	factory := cmdHelper.GetCommandFactory(cmd)
	if factory == nil {
		factory = gf
	}
	return factory.ParseCommand(content)
}

// Override
func (gf GroupCommandFactory) ParseCommand(content StringKeyMap) Command {
	// check 'sn', 'command', 'group'
	if !ContainsKey(content, "sn") || !ContainsKey(content, "command") {
		// content.sn should not be empty
		// content.command should not be empty
		return nil
	} else if !ContainsKey(content, "group") {
		// content.group should not be empty
		return nil
	}
	return NewGroupCommandWithMap(content)
}

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
	. "github.com/dimchat/core-go/ext"
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/dkd-go/ext"
	. "github.com/dimchat/dkd-go/protocol"
	. "github.com/dimchat/mkm-go/types"
)

type ICommandGeneralFactory interface {
	GeneralCommandHelper
	CommandHelper
}

func NewCommandGeneralFactory() ICommandGeneralFactory {
	gf := &CommandGeneralFactory{}
	return gf.Init()
}

type CommandGeneralFactory struct {
	//ICommandGeneralFactory

	commandFactories map[string]CommandFactory
}

func (gf *CommandGeneralFactory) Init() ICommandGeneralFactory {
	gf.commandFactories = make(map[string]CommandFactory)
	return gf
}

// Override
func (gf *CommandGeneralFactory) GetCMD(content StringKeyMap, defaultValue string) string {
	cmd := content["command"]
	return ConvertString(cmd, defaultValue)
}

//
//  Command Helper
//

// Override
func (gf *CommandGeneralFactory) SetCommandFactory(cmd string, factory CommandFactory) {
	gf.commandFactories[cmd] = factory
}

// Override
func (gf *CommandGeneralFactory) GetCommandFactory(cmd string) CommandFactory {
	return gf.commandFactories[cmd]
}

// Override
func (gf *CommandGeneralFactory) ParseCommand(content interface{}) Command {
	if content == nil {
		return nil
	} else if command, ok := content.(Command); ok {
		return command
	}
	info := FetchMap(content)
	if info == nil {
		//panic("command error")
		return nil
	}
	// get factory by command name
	cmd := gf.GetCMD(info, "")
	factory := gf.GetCommandFactory(cmd)
	if factory == nil {
		// unknown command name, get base command factory
		factory = getDefaultFactory(info)
		if factory == nil {
			//panic("cannot parse command")
			return nil
		}
	}
	return factory.ParseCommand(info)
}

func getDefaultFactory(info StringKeyMap) CommandFactory {
	helper := GetGeneralMessageHelper()
	contentHelper := GetContentHelper()
	// get factory by content type
	msgType := helper.GetContentType(info, "")
	if msgType == "" {
		//panic("cannot parse command")
		return nil
	}
	factory := contentHelper.GetContentFactory(msgType)
	if fact, ok := factory.(CommandFactory); ok {
		return fact
	}
	//panic("command factory not implement CommandFactory")
	return nil
}

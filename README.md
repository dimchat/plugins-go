# DIM Plugins (Go)

[![License](https://img.shields.io/github/license/dimchat/plugins-go)](https://github.com/dimchat/plugins-go/blob/main/LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/dimchat/plugins-go/pulls)
[![Platform](https://img.shields.io/github/go-mod/go-version/dimchat/plugins-go)](https://github.com/dimchat/plugins-go/wiki)
[![Issues](https://img.shields.io/github/issues/dimchat/plugins-go)](https://github.com/dimchat/plugins-go/issues)
[![Repo Size](https://img.shields.io/github/repo-size/dimchat/plugins-go)](https://github.com/dimchat/plugins-go/archive/refs/heads/main.zip)
[![Tags](https://img.shields.io/github/tag/dimchat/plugins-go)](https://github.com/dimchat/plugins-go/tags)

[![Watchers](https://img.shields.io/github/watchers/dimchat/plugins-go)](https://github.com/dimchat/plugins-go/watchers)
[![Forks](https://img.shields.io/github/forks/dimchat/plugins-go)](https://github.com/dimchat/plugins-go/forks)
[![Stars](https://img.shields.io/github/stars/dimchat/plugins-go)](https://github.com/dimchat/plugins-go/stargazers)
[![Followers](https://img.shields.io/github/followers/dimchat)](https://github.com/orgs/dimchat/followers)

## Plugins

1. Data Coding
   * Base-58
   * Base-64
   * Hex
   * UTF-8
   * JsON
   * PNF _(Portable Network File)_
   * TED _(Transportable Encoded Data)_
2. Digest Digest
   * MD-5
   * SHA-1
   * SHA-256
   * Keccak-256
   * RipeMD-160
3. Cryptography
   * AES-256 _(AES/CBC/PKCS7Padding)_
   * RSA-1024 _(RSA/ECB/PKCS1Padding)_, _(SHA256withRSA)_
   * ECC _(Secp256k1)_
4. Address
   * BTC
   * ETH
5. Meta
   * MKM _(Default)_
   * BTC
   * ETH
6. Document
   * Visa _(User)_
   * Profile
   * Bulletin _(Group)_

## Extends

### Address

```go
import (
	"strings"

	. "github.com/dimchat/mkm-go/mkm"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/mem"
	. "github.com/dimchat/plugins-go/mkm"
)

func NewCompatibleAddressFactory() AddressFactory {
	return &compatibleAddressFactory{}
}

/**
 *  Compatible Address Factory
 */

type compatibleAddressFactory struct {
	//AddressFactory
}

// Override
func (compatibleAddressFactory) GenerateAddress(meta Meta, network EntityType) Address {
	address := meta.GenerateAddress(network)
	if address != nil {
		cache := GetAddressCache()
		cache.Put(address.String(), address)
	}
	return address
}

// Override
func (compatibleAddressFactory) ParseAddress(str string) Address {
	cache := GetAddressCache()
	address := cache.Get(str)
	if address == nil {
		address = parseAddress(str)
		if address != nil {
			cache.Put(str, address)
		}
	}
	return address
}

// protected
func parseAddress(str string) Address {
	if str == "" {
		//panic("address is empty")
		return nil
	}
	size := len(str)
	if size == 0 {
		//panic("address is empty")
		return nil
	} else if size == 8 {
		// "anywhere"
		lower := strings.ToLower(str)
		if ANYWHERE.Equal(lower) {
			return ANYWHERE
		}
	} else if size == 10 {
		// "everywhere"
		lower := strings.ToLower(str)
		if EVERYWHERE.Equal(lower) {
			return EVERYWHERE
		}
	}
	var result Address
	if 26 <= size && size <= 35 {
		// BTC
		result = ParseBTCAddress(str)
	} else if size == 42 {
		// ETH
		result = ParseETHAddress(str)
	} else {
		//panic("invalid address")
	}
	//
	//  TODO: parse for other types of address
	//
	if result == nil {
		if 4 <= size && size <= 64 {
			return NewUnknownAddress(str)
		}
		panic("invalid address: " + str)
	}
	return result
}

/**
 *  Unsupported Address
 */

type UnknownAddress struct {
	//Address
	ConstantString
}

func NewUnknownAddress(address string) Address {
	return &UnknownAddress{
		ConstantString: *NewConstantString(address),
	}
}

// Override
func (UnknownAddress) Network() EntityType {
	return USER
}
```

### Meta

```go
import (
	. "github.com/dimchat/mkm-go/ext"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/mkm-go/types"
	. "github.com/dimchat/plugins-go/mem"
	. "github.com/dimchat/plugins-go/mkm"
)

func NewCompatibleMetaFactory(version MetaType) MetaFactory {
	return &compatibleMetaFactory{
		BaseMetaFactory{
			Type: version,
		},
	}
}

/**
 *  Compatible Meta Factory
 */

type compatibleMetaFactory struct {
	BaseMetaFactory
}

// Override
func (factory compatibleMetaFactory) ParseMeta(info StringKeyMap) Meta {
	// check 'type', 'key', 'seed', 'fingerprint'
	if !ContainsKey(info, "type") || !ContainsKey(info, "key") {
		// meta.type should not be empty
		// meta.key should not be empty
		return nil
	} else if !ContainsKey(info, "seed") {
		if ContainsKey(info, "fingerprint") {
			//panic("meta error")
			return nil
		}
	} else if !ContainsKey(info, "fingerprint") {
		//panic("meta error")
		return nil
	}
	// create meta for type
	var out Meta
	helper := GetGeneralAccountHelper()
	version := helper.GetMetaType(info, "")
	switch version {
	case "1", "mkm", "MKM":
		out = NewDefaultMeta(info, "", nil, "", nil)
		break
	case "2", "btc", "BTC":
		out = NewBTCMeta(info, "", nil, "", nil)
		break
	case "4", "eth", "ETH":
		out = NewETHMeta(info, "", nil, "", nil)
		break
	default:
		break
	}
	if out.IsValid() {
		return out
	}
	//panic("meta error")
	return nil
}
```

### Plugin Loader

```go
import (
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/mkm-go/digest"
	. "github.com/dimchat/mkm-go/protocol"
	. "github.com/dimchat/plugins-go/ext"
	. "github.com/dimpart/demo-go/sdk/common/digest"
	. "github.com/dimpart/demo-go/sdk/common/mkm"
)

/**
 *  Plugin Loader
 */

type CompatiblePluginLoader struct {
	PluginLoader
}

// Override
func (loader CompatiblePluginLoader) Load() {
	loader.PluginLoader.Load()

	registerAddressFactory()

	registerMetaFactories()

}

/**
 *  Address factory
 */

func registerAddressFactory() {

	// Address
	SetAddressFactory(NewCompatibleAddressFactory())

}

/**
 *  Meta factories
 */

func registerMetaFactories() {
	
	mkm := NewCompatibleMetaFactory(MKM)
	btc := NewCompatibleMetaFactory(BTC)
	eth := NewCompatibleMetaFactory(ETH)

	SetMetaFactory("1", mkm)
	SetMetaFactory("2", btc)
	SetMetaFactory("4", eth)

	SetMetaFactory("mkm", mkm)
	SetMetaFactory("btc", btc)
	SetMetaFactory("eth", eth)

	SetMetaFactory("MKM", mkm)
	SetMetaFactory("BTC", btc)
	SetMetaFactory("ETH", eth)
}
```

### ExtensionLoader

```go
import (
	. "github.com/dimchat/core-go/protocol"
	. "github.com/dimchat/dkd-go/protocol"
	. "github.com/dimchat/plugins-go/ext"
	. "github.com/dimpart/demo-go/sdk/common/dkd"
	. "github.com/dimpart/demo-go/sdk/common/protocol"
)

/**
 *  Extensions Loader
 */

type CommonExtensionLoader struct {
	ExtensionLoader
}

// Override
func (loader CommonExtensionLoader) Load() {
	loader.ExtensionLoader.Load()

	registerContentFactories()

	registerCommandFactories()

}

/**
 *  Extended content factories
 */

func registerContentFactories() {

	// Application Customized Content
	registerContentCreator(ContentType.APPLICATION, NewAppContentWithMap)
	registerContentCreator(ContentType.CUSTOMIZED, NewCustomizedContentWithMap)

}

func registerContentCreator(msgType string, fn FuncCreateContent) {
	factory := NewContentFactory(fn)
	SetContentFactory(msgType, factory)
}

/**
 *  Extended command factories
 */

func registerCommandFactories() {

	// Handshake Command
	registerCommandCreator(HANDSHAKE, NewHandshakeCommandWithMap)

}

func registerCommandCreator(cmd string, fn FuncCreateCommand) {
	factory := NewCommandFactory(fn)
	SetCommandFactory(cmd, factory)
}
```

## Usage

You must load all plugins before your business run:

```go
import (
	. "github.com/dimchat/plugins-go/ext"
	. "github.com/dimpart/demo-go/sdk/common/ext"
)

type LibraryLoader struct {
	extensionLoader IExtensionLoader
	pluginLoader    IPluginLoader

	// flag
	loaded bool
}

func NewLibraryLoader() ILibraryLoader {
	return &LibraryLoader{
		extensionLoader: &ClientExtensionLoader{},
		pluginLoader:    &CommonPluginLoader{},
		loaded:          false,
	}
}

func (loader *LibraryLoader) Run() {
	if loader.loaded {
		// no need to load it again
		return
	}
	// mark it to loaded
	loader.loaded = true
	// try to load all extensions
	loader.Load()
}

// protected
func (loader *LibraryLoader) Load() {
	loader.extensionLoader.Load()
	loader.pluginLoader.Load()
}


func init() {

  loader := NewLibraryLoader()
  loader.Run()
  
  // do your jobs after all extensions & plugins loaded
  
}
```

You must ensure that every ```Address``` you extend has a ```Meta``` type that can correspond to it one by one.

----

Copyright &copy; 2020-2026 Albert Moky
[![Followers](https://img.shields.io/github/followers/moky)](https://github.com/moky?tab=followers)

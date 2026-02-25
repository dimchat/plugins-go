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
package mem

import . "github.com/dimchat/mkm-go/types"

// MemoryCache defines a generic in-memory cache interface for key-value storage
//
// Generic Constraints:
//
//	K: comparable - Cache key type (must support ==/!= comparison for lookups)
//	V: any        - Cache value type (can be any Go type)
//
// Core capabilities: Key-value access, storage, and memory optimization
type MemoryCache[K comparable, V any] interface {

	// Get retrieves the cached value for the specified key
	//
	// If the key does not exist in the cache, returns the zero value for type V
	//
	// Parameters:
	//   - key - Cache key to look up (must be comparable type)
	// Returns: Cached value for the key (zero value if key not found)
	Get(key K) V

	// Put stores a value in the cache associated with the specified key
	//
	// Overwrites any existing value if the key already exists
	//
	// Parameters:
	//   - key   - Cache key to associate with the value (must be comparable type)
	//   - value - Value to store in the cache (can be any type)
	Put(key K, value V)

	// Size returns the current number of entries in the cache
	//
	// Provides real-time count of key-value pairs stored in memory
	//
	// Returns: Total number of cached entries (0 if cache is empty)
	Size() int

	// ReduceMemory performs garbage collection/optimization on the cache
	//
	// Implements cache eviction strategies (e.g., LRU, TTL, size limit) to free memory
	//
	// Returns: Number of cache entries remaining after memory reduction (post-eviction count)
	ReduceMemory() int
}

// ContainsKey checks if a specific key exists in a StringKeyMap
//
// # Provides a convenience wrapper for map existence checking
//
// Parameters:
//   - info - StringKeyMap (map[string]any) to check for the key
//   - key  - String key to look up in the map
//
// Returns: true if key exists in the map, false otherwise
func ContainsKey(info StringKeyMap, key string) bool {
	_, exist := info[key]
	return exist
}

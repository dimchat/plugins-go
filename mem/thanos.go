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

/**
 *  Random Cache
 */
type ThanosCache[K comparable, V any] struct {
	//MemoryCache

	table map[K]V
}

func NewThanosCache[K comparable, V any]() *ThanosCache[K, V] {
	return &ThanosCache[K, V]{
		table: make(map[K]V, 512),
	}
}

// Override
func (cache *ThanosCache[K, V]) Get(key K) V {
	value, _ := cache.table[key]
	return value
}

// Override
func (cache *ThanosCache[K, V]) Put(key K, value V) {
	cache.table[key] = value
}

// Override
func (cache *ThanosCache[K, V]) ReduceMemory() int {
	finger := 0
	finger = thanos(cache.table, finger)
	return finger >> 1
}

/**
 *  Thanos can kill half lives of a world with a snap of the finger
 */
func thanos[K comparable, V any](planet map[K]V, finger int) int {
	for key := range planet {
		finger++
		if (finger & 1) == 1 {
			// kill it
			delete(planet, key)
		}
		// let it go
	}
	return finger
}

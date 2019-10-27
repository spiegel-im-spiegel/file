package file

import (
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type pathList struct {
	mutex *sync.Mutex
	paths []string
}

func newPathList() *pathList {
	return &pathList{mutex: &sync.Mutex{}, paths: make([]string, 0)}
}

func (p *pathList) Init() {
	p.mutex.Lock()
	p.paths = p.paths[:0]
	p.mutex.Unlock()
}

func (p *pathList) Append(paths ...string) {
	if len(paths) == 0 {
		return
	}
	p.mutex.Lock()
	for _, path := range paths {
		p.paths = append(p.paths, normalizePath(path))
	}
	p.mutex.Unlock()
}

func (p *pathList) Add(list *pathList) {
	if list.Len() == 0 {
		return
	}
	p.mutex.Lock()
	p.paths = append(p.paths, list.paths...)
	p.mutex.Unlock()
}

func (p *pathList) Strings(slash bool) []string {
	p.mutex.Lock()
	paths := make([]string, len(p.paths))
	if slash {
		for i, path := range p.paths {
			paths[i] = filepath.ToSlash(path)
		}
	} else {
		copy(paths, p.paths)
	}
	sort.Strings(paths)
	p.mutex.Unlock()
	return paths
}

func (p *pathList) Len() int {
	return len(p.paths)
}

func normalizePath(path string) string {
	info, err := os.Stat(path)
	if err == nil && info.IsDir() {
		return filepath.Clean(path) + string(filepath.Separator)
	}
	return filepath.Clean(path)
}

/* MIT License
 *
 * Copyright 2019 Spiegel
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
 */

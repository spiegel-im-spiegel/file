package file

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/saracen/walker"
	"github.com/spiegel-im-spiegel/errs"
)

type glob struct {
	flags Flags
	ctx   context.Context
}

//Glob returns the names of all files matching pattern.
//This function expands filepath.Glob
func Glob(pattern string, opts ...GlobOptFunc) ([]string, error) {
	return GlobWithContext(context.Background(), pattern, opts...)
}

//GlobWithContext returns the names of all files matching pattern with context.Context.
func GlobWithContext(ctx context.Context, pattern string, opts ...GlobOptFunc) ([]string, error) {
	g := newGlob(ctx, opts...)
	paths, err := g.getPaths("", pattern)
	return g.normalize(paths), err
}

//GlobOptFunc is self-referential function for functional options pattern.
type GlobOptFunc func(*glob)

func newGlob(ctx context.Context, opts ...GlobOptFunc) *glob {
	g := &glob{flags: StdFlags, ctx: ctx}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

//WithFlags returns function for setting glob instance.
func WithFlags(f Flags) GlobOptFunc {
	return func(g *glob) {
		g.flags = f
	}
}

func (g *glob) getPaths(rootDir, path string) (*pathList, error) {
	paths := newPathList()
	roots, err := g.getRoots(rootDir)
	if err != nil {
		return paths, errs.Wrap(err, "", errs.WithContext("rootDir", rootDir))
	}
	for _, root := range roots.Strings(true) {
		ps, err := g.getPathsNext(root, path)
		if err != nil {
			return paths, errs.Wrap(err, "", errs.WithContext("root", root), errs.WithContext("path", path))
		}
		paths.Add(ps)
	}
	return paths, nil
}

func (g *glob) getPathsNext(root, path string) (*pathList, error) {
	paths := newPathList()
	path = filepath.ToSlash(path)
	if strings.HasPrefix(path, "**/") {
		rt := root
		if len(rt) == 0 {
			rt = "./"
		}
		roots, err := g.walkDir(rt)
		if err != nil {
			return paths, errs.Wrap(err, "", errs.WithContext("root", root), errs.WithContext("path", path))
		}
		subPath := path[3:] //exclude "**/"
		if roots.Len() > 0 {
			var wg sync.WaitGroup
			var lastErr atomic.Value
			for _, rt := range roots.Strings(true) {
				wg.Add(1)
				go func(r, p string) {
					defer wg.Done()
					ps, err := g.getPaths(r, p)
					if err != nil {
						lastErr.Store(errs.Wrap(err, "", errs.WithContext("root", r), errs.WithContext("path", p)))
						return
					}
					paths.Add(ps)
				}(rt, subPath)
			}
			wg.Wait()
			err := lastErr.Load()
			if err != nil {
				return paths, err.(error)
			}
			return paths, nil
		}
		path = subPath
	} else if i := strings.Index(path, "/**/"); i >= 0 {
		return g.getPaths(filepath.ToSlash(filepath.Join(root, path[:i+1])), path[i+1:])
	}

	if len(path) > 0 {
		paths.Init()
		dirFlag := false
		if strings.HasSuffix(path, "/") {
			path = path[:len(path)-1]
			dirFlag = true
		}
		pattern := filepath.ToSlash(filepath.Join(root, path))
		ps, err := filepath.Glob(pattern)
		if err != nil {
			return paths, errs.Wrap(err, "", errs.WithContext("pattern", pattern))
		}
		for _, p := range ps {
			info, err := os.Stat(p)
			if err != nil {
				return paths, errs.Wrap(err, "", errs.WithContext("path", p))
			}
			if dirFlag && info.IsDir() || !dirFlag {
				paths.Append(p)
			}
		}
		return paths, nil
	}

	paths.Init()
	paths.Append(root)
	return paths, nil
}

func (g *glob) getRoots(rootDir string) (*pathList, error) {
	roots := newPathList()
	if len(rootDir) == 0 {
		roots.Append("")
		return roots, nil
	}
	rootDir = filepath.ToSlash(rootDir)
	if strings.HasSuffix(rootDir, "/") {
		rootDir = rootDir[:len(rootDir)-1]
	}
	paths, err := filepath.Glob(rootDir)
	if err != nil {
		return roots, errs.Wrap(err, "", errs.WithContext("rootDir", rootDir))
	}
	for _, path := range paths {
		info, err := os.Stat(path)
		if err != nil {
			return roots, errs.Wrap(err, "", errs.WithContext("path", path))
		}
		if info.IsDir() {
			roots.Append(path)
		}
	}
	return roots, nil
}

func (g *glob) walkDir(root string) (*pathList, error) {
	paths := newPathList()
	err := walker.WalkWithContext(g.ctx, root, func(path string, info os.FileInfo) error {
		if info.IsDir() {
			paths.Append(path)
		}
		return nil
	})
	return paths, err
}

func (g *glob) normalize(paths *pathList) []string {
	if paths.Len() == 0 {
		return []string{}
	}
	pathMap := map[string]string{}
	for _, path := range paths.Strings(false) {
		absPath, err := filepath.Abs(path)
		if err != nil {
			absPath = ""
		} else {
			absPath = normalizePath(absPath)
		}
		if g.flags.AbsolutePath() {
			path = absPath
		}
		if len(absPath) > 0 {
			if strings.HasSuffix(absPath, string(filepath.Separator)) {
				if g.flags.ContainsDir() {
					pathMap[absPath] = path
				}
			} else if g.flags.ContainsFile() {
				pathMap[absPath] = path
			}
		}
	}
	unqPaths := make([]string, 0, len(pathMap))
	for _, path := range pathMap {
		if g.flags.SeparatorSlash() {
			path = filepath.ToSlash(path)
		}
		unqPaths = append(unqPaths, path)
	}
	sort.Strings(unqPaths)
	return unqPaths
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

package file_test

import (
	"reflect"
	"testing"

	"github.com/spiegel-im-spiegel/file"
)

func TestGlob(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "XXX/**/", result: []string{}},
		{path: "testdata/**/", result: []string{"testdata/", "testdata/include/"}},
		{path: "testdata/*/", result: []string{"testdata/include/"}},
		{path: "testdata/**/*", result: []string{"testdata/include/", "testdata/include/source.h", "testdata/source.c"}},
		{path: "**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "**/**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "./**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "test*/**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "test*/**/../**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
	}
	for _, tc := range testCases {
		strs, err := file.Glob(tc.path, file.WithFlags(file.StdFlags|file.SeparatorSlash))
		if err != nil {
			t.Errorf("file.Glob(\"%s\") is \"%+v\", want <nil>.", tc.path, err)
		} else if !reflect.DeepEqual(strs, tc.result) {
			t.Errorf("file.Glob(\"%s\") = %v, want %v.", tc.path, strs, tc.result)
		}
	}
}

func TestGlobFileOnly(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "XXX/**/", result: []string{}},
		{path: "testdata/**/", result: []string{}},
		{path: "testdata/*/", result: []string{}},
		{path: "testdata/**/*", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "**/**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "./**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "test*/**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
		{path: "test*/**/../**/*.[ch]", result: []string{"testdata/include/source.h", "testdata/source.c"}},
	}
	for _, tc := range testCases {
		strs, err := file.Glob(tc.path, file.WithFlags(file.ContainsFile|file.SeparatorSlash))
		if err != nil {
			t.Errorf("file.Glob(\"%s\") is \"%+v\", want <nil>.", tc.path, err)
		} else if !reflect.DeepEqual(strs, tc.result) {
			t.Errorf("file.Glob(\"%s\") = %v, want %v.", tc.path, strs, tc.result)
		}
	}
}

func TestGlobDirOnly(t *testing.T) {
	testCases := []struct {
		path   string
		result []string
	}{
		{path: "XXX/**/", result: []string{}},
		{path: "testdata/**/", result: []string{"testdata/", "testdata/include/"}},
		{path: "testdata/*/", result: []string{"testdata/include/"}},
		{path: "testdata/**/*", result: []string{"testdata/include/"}},
		{path: "**/*.[ch]", result: []string{}},
	}
	for _, tc := range testCases {
		strs, err := file.Glob(tc.path, file.WithFlags(file.ContainsDir|file.SeparatorSlash))
		if err != nil {
			t.Errorf("file.Glob(\"%s\") is \"%+v\", want <nil>.", tc.path, err)
		} else if !reflect.DeepEqual(strs, tc.result) {
			t.Errorf("file.Glob(\"%s\") = %v, want %v.", tc.path, strs, tc.result)
		}
	}
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

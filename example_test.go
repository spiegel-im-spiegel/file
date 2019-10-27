package file_test

import (
	"context"
	"fmt"
	"os"

	"github.com/spiegel-im-spiegel/file"
)

func ExampleGlob() {
	matches, err := file.Glob("**/*.[ch]", file.WithFlags(file.StdFlags|file.SeparatorSlash))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	for _, path := range matches {
		fmt.Println(path)
	}
	//Output:
	//testdata/include/source.h
	//testdata/source.c
}

func ExampleGlobWithContext() {
	matches, err := file.GlobWithContext(context.Background(), "**/*.[ch]", file.WithFlags(file.StdFlags|file.SeparatorSlash))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	for _, path := range matches {
		fmt.Println(path)
	}
	//Output:
	//testdata/include/source.h
	//testdata/source.c
}

func ExampleWithFlags() {
	matches, err := file.GlobWithContext(context.Background(), "**/*.[ch]", file.WithFlags(file.StdFlags|file.AbsolutePath|file.SeparatorSlash))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		return
	}
	for _, path := range matches {
		fmt.Println(path)
	}
	// Output:
	// testdata/include/source.h
	// testdata/source.c
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

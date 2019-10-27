package file

//Flags is type of operation flag in Glob() function.
type Flags uint

//Operation flag in Glob() function.
const (
	ContainsFile   Flags = 1 << iota //contains file
	ContainsDir                      //contains directory
	SeparatorSlash                   //use slash ('/') character for separator character in path
	AbsolutePath                     //output absolute representation of path
	StdFlags       = ContainsFile | ContainsDir
)

//ContainsFile returns status of ContainsFile flag.
func (f Flags) ContainsFile() bool {
	return (f & ContainsFile) != 0
}

//ContainsDir returns status of ContainsDir flag.
func (f Flags) ContainsDir() bool {
	return (f & ContainsDir) != 0
}

//SeparatorSlash returns status of SeparatorSlash flag.
func (f Flags) SeparatorSlash() bool {
	return (f & SeparatorSlash) != 0
}

//AbsolutePath returns status of AbsolutePath flag.
func (f Flags) AbsolutePath() bool {
	return (f & AbsolutePath) != 0
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

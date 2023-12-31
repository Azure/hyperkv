/*
Copyright © 2023 (c) Microsoft Corporation

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package hyperkv

import (
	"strings"
	"unicode/utf8"
)

type hyperkvLexerImpl struct {
	line   []byte
	len    int
	index  int
	peek   rune
	Result []*HyperkvItem
}

func newHyperkvLexer(line []byte) *hyperkvLexerImpl {
	lexerr := &hyperkvLexerImpl{line: line, len: len(line), index: 0}
	lexerr.peek = lexerr.next()
	lexerr.Result = make([]*HyperkvItem, 0)
	return lexerr
}

func (l *hyperkvLexerImpl) Lex(lval *hyperkvSymType) int {
	if l.index == l.len || l.peek == utf8.RuneError {
		return eof
	}
	if l.peek == '\u0000' {
		l.peek = l.next()
		for ; l.peek == '\u0000'; l.peek = l.next() {
		}
		return space
	} else {
		var result strings.Builder
		for ; l.peek != '\u0000' && l.peek != utf8.RuneError; l.peek = l.next() {
			_, _ = result.WriteRune(l.peek)
		}
		lval.content = result.String()
		return str
	}
}

func (l *hyperkvLexerImpl) Error(s string) {
}

func (l *hyperkvLexerImpl) next() rune {
	rune, width := utf8.DecodeRune(l.line[l.index:])
	l.index += width
	if rune == utf8.RuneError && width == 1 {
		return l.next()
	}
	return rune
}

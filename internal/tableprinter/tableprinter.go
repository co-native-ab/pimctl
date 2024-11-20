// This code is heavily borrowed from the GitHub CLI project. The original code can be found here:
//
// https://github.com/cli/cli/blob/trunk/internal/tableprinter/table_printer.go
// https://github.com/cli/cli/blob/trunk/internal/text/text.go
//
// MIT License
//
// Copyright (c) 2019 GitHub Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package tableprinter

import (
	"fmt"
	"io"
	"math"
	"net/url"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/cli/cli/v2/pkg/iostreams"
	"github.com/cli/go-gh/v2/pkg/tableprinter"
	"github.com/cli/go-gh/v2/pkg/text"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type TablePrinter struct {
	tableprinter.TablePrinter
	isTTY bool
	cs    *iostreams.ColorScheme
}

// IsTTY gets whether the TablePrinter will render to a terminal.
func (t *TablePrinter) IsTTY() bool {
	return t.isTTY
}

// AddTimeField in TTY mode displays the fuzzy time difference between now and t.
// In non-TTY mode it just displays t with the time.RFC3339 format.
func (tp *TablePrinter) AddTimeField(now, t time.Time, c func(string) string) {
	var tf string
	if tp.isTTY {
		tf = FuzzyAgo(now, t)
	} else {
		tf = t.Format(time.RFC3339)
	}
	tp.AddField(tf, WithColor(c))
}

var (
	WithColor    = tableprinter.WithColor
	WithPadding  = tableprinter.WithPadding
	WithTruncate = tableprinter.WithTruncate
)

type headerOption struct {
	columns []string
}

// New creates a TablePrinter from an IOStreams.
func New(ios *iostreams.IOStreams, headers headerOption) *TablePrinter {
	maxWidth := 80
	isTTY := ios.IsStdoutTTY()
	if isTTY {
		maxWidth = ios.TerminalWidth()
	}

	return NewWithWriter(ios.Out, isTTY, maxWidth, ios.ColorScheme(), headers)
}

// NewWithWriter creates a TablePrinter from a Writer, whether the output is a terminal, the terminal width, and more.
func NewWithWriter(w io.Writer, isTTY bool, maxWidth int, cs *iostreams.ColorScheme, headers headerOption) *TablePrinter {
	tp := &TablePrinter{
		TablePrinter: tableprinter.New(w, isTTY, maxWidth),
		isTTY:        isTTY,
		cs:           cs,
	}

	if isTTY && len(headers.columns) > 0 {
		// Make sure all headers are uppercase, taking a copy of the headers to avoid modifying the original slice.
		upperCasedHeaders := make([]string, len(headers.columns))
		for i := range headers.columns {
			upperCasedHeaders[i] = strings.ToUpper(headers.columns[i])
		}

		// Make sure all header columns are padded - even the last one. Previously, the last header column
		// was not padded. In tests cs.Enabled() is false which allows us to avoid having to fix up
		// numerous tests that verify header padding.
		var paddingFunc func(int, string) string
		if cs.Enabled() {
			paddingFunc = text.PadRight
		}

		tp.AddHeader(
			upperCasedHeaders,
			WithPadding(paddingFunc),
			WithColor(cs.LightGrayUnderline),
		)
	}

	return tp
}

// WithHeader defines the column names for a table.
// Panics if columns is nil or empty.
func WithHeader(columns ...string) headerOption {
	if len(columns) == 0 {
		panic("must define header columns")
	}
	return headerOption{columns}
}

// NoHeader disable printing or checking for a table header.
//
// Deprecated: use WithHeader unless required otherwise.
var NoHeader = headerOption{}

var whitespaceRE = regexp.MustCompile(`\s+`)

func Indent(s, indent string) string {
	return text.Indent(s, indent)
}

// Title returns a copy of the string s with all Unicode letters that begin words mapped to their Unicode title case.
func Title(s string) string {
	c := cases.Title(language.English)
	return c.String(s)
}

// RemoveExcessiveWhitespace returns a copy of the string s with excessive whitespace removed.
func RemoveExcessiveWhitespace(s string) string {
	return whitespaceRE.ReplaceAllString(strings.TrimSpace(s), " ")
}

func DisplayWidth(s string) int {
	return text.DisplayWidth(s)
}

func Truncate(maxWidth int, s string) string {
	return text.Truncate(maxWidth, s)
}

func Pluralize(num int, thing string) string {
	return text.Pluralize(num, thing)
}

func FuzzyAgo(a, b time.Time) string {
	return text.RelativeTimeAgo(a, b)
}

// FuzzyAgoAbbr is an abbreviated version of FuzzyAgo. It returns a human readable string of the
// time duration between a and b that is estimated to the nearest unit of time.
func FuzzyAgoAbbr(a, b time.Time) string {
	ago := a.Sub(b)

	if ago < time.Hour {
		return fmt.Sprintf("%d%s", int(ago.Minutes()), "m")
	}
	if ago < 24*time.Hour {
		return fmt.Sprintf("%d%s", int(ago.Hours()), "h")
	}
	if ago < 30*24*time.Hour {
		return fmt.Sprintf("%d%s", int(ago.Hours())/24, "d")
	}

	return b.Format("Jan _2, 2006")
}

// DisplayURL returns a copy of the string urlStr removing everything except the scheme, hostname, and path.
// If the scheme is not specified, "https" is assumed.
// If there is an error parsing urlStr then urlStr is returned without modification.
func DisplayURL(urlStr string) string {
	u, err := url.Parse(urlStr)
	if err != nil {
		return urlStr
	}
	scheme := u.Scheme
	if scheme == "" {
		scheme = "https"
	}
	return scheme + "://" + u.Hostname() + u.Path
}

// RemoveDiacritics returns the input value without "diacritics", or accent marks
func RemoveDiacritics(value string) string {
	return text.RemoveDiacritics(value)
}

func PadRight(maxWidth int, s string) string {
	return text.PadRight(maxWidth, s)
}

// FormatSlice concatenates elements of the given string slice into a
// well-formatted, possibly multiline, string with specific line length limit.
// Elements can be optionally surrounded by custom strings (e.g., quotes or
// brackets). If the lineLength argument is non-positive, no line length limit
// will be applied.
func FormatSlice(values []string, lineLength uint, indent uint, prependWith string, appendWith string, sort bool) string {
	if lineLength <= 0 {
		lineLength = math.MaxInt
	}

	sortedValues := values
	if sort {
		sortedValues = slices.Clone(values)
		slices.Sort(sortedValues)
	}

	pre := strings.Repeat(" ", int(indent))
	if len(sortedValues) == 0 {
		return pre
	} else if len(sortedValues) == 1 {
		return pre + sortedValues[0]
	}

	builder := strings.Builder{}
	currentLineLength := 0
	sep := ","
	ws := " "

	for i := 0; i < len(sortedValues); i++ {
		v := prependWith + sortedValues[i] + appendWith
		isLast := i == -1+len(sortedValues)

		if currentLineLength == 0 {
			builder.WriteString(pre)
			builder.WriteString(v)
			currentLineLength += len(v)
			if !isLast {
				builder.WriteString(sep)
				currentLineLength += len(sep)
			}
		} else {
			if !isLast && currentLineLength+len(ws)+len(v)+len(sep) > int(lineLength) ||
				isLast && currentLineLength+len(ws)+len(v) > int(lineLength) {
				currentLineLength = 0
				builder.WriteString("\n")
				i--
				continue
			}

			builder.WriteString(ws)
			builder.WriteString(v)
			currentLineLength += len(ws) + len(v)
			if !isLast {
				builder.WriteString(sep)
				currentLineLength += len(sep)
			}
		}
	}
	return builder.String()
}

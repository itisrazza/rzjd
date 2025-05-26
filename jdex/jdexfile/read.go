// rzjd - Razza's Johnny.Decimal Management System
// Copyright (C) 2025 Raresh Nistor
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package jdexfile

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/itisrazza/rzjd/jdex"
)

type readContext struct {
	index *jdex.Index

	lineNumber       int
	multilineComment bool

	lastID    jdex.ACID
	lastEntry jdex.Entry
}

var ErrParse = errors.New("failed to parse jdex")

var commentSingleRegex = regexp.MustCompile(`//.*$`)
var commentMultilineStartRegex = regexp.MustCompile(`/\*.*$`)
var commentMultilineEndRegex = regexp.MustCompile(`^.*\*/`)

var areaRegex = regexp.MustCompile(`^([A-Z0-9]0-[A-Z0-9]9)\s+(.+)$`)
var categoryRegex = regexp.MustCompile(`^([A-Z0-9]+)\s+(.+)$`)
var entryRegex = regexp.MustCompile(`^([A-Z0-9\.\+]+)?\s+(.+)$`)
var metadataRegex = regexp.MustCompile(`^-\s*(.+?)\s*:\s*(.+?)\s*$`)

func Read(r io.Reader) (index jdex.Index, err error) {
	index, err = jdex.NewIndex()
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	ctx := readContext{
		index:            &index,
		multilineComment: false,
		lineNumber:       0,
	}

	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return
		}
		ctx.lineNumber++

		line := scanner.Text()
		line, err = readPreprocessLine(line, &ctx)
		if err != nil {
			err = fmt.Errorf("line %d: %w", ctx.lineNumber, err)
			return
		}

		if line == "" {
			continue
		}

		err = readProcessLine(line, &ctx)
		if err != nil {
			err = fmt.Errorf("line %d: %w", ctx.lineNumber, err)
			return
		}
	}

	return
}

func readPreprocessLine(l string, ctx *readContext) (line string, err error) {
	line = strings.TrimSpace(l)

	// remove single comments
	line = commentSingleRegex.ReplaceAllString(line, "")

	// remove multiline comments
	if !ctx.multilineComment {
		line = commentMultilineEndRegex.ReplaceAllStringFunc(line, func(comment string) string {
			ctx.multilineComment = true
			return ""
		})
	} else {
		line = commentMultilineEndRegex.ReplaceAllStringFunc(line, func(comment string) string {
			ctx.multilineComment = false
			return ""
		})
	}

	// blank out the line entirely if we're still in a multi-line comment
	if ctx.multilineComment {
		line = ""
	}

	return
}

func readProcessLine(line string, ctx *readContext) (err error) {
	if areaRegex.MatchString(line) {
		return readProcessAreaLine(line, ctx)
	}

	if categoryRegex.MatchString(line) {
		return readProcessCategoryLine(line, ctx)
	}

	if entryRegex.MatchString(line) {
		return readProcessEntryLine(line, ctx)
	}

	if metadataRegex.MatchString(line) {
		return readProcessMetadataLine(line, ctx)
	}

	panic("unimplemented line kind")
}

func readProcessAreaLine(line string, ctx *readContext) error {
	matches := areaRegex.FindStringSubmatch(line)

	// FIXME: add some more guardrails around ID indexing like that
	id := jdex.ACID{Area: matches[1][0]}
	name := matches[2]

	err := ctx.index.PutArea(id, name)
	if err != nil {
		return err
	}

	ctx.lastID = id
	return nil
}

func readProcessCategoryLine(line string, ctx *readContext) error {
	matches := categoryRegex.FindStringSubmatch(line)

	// FIXME: add some more guardrails around ID indexing like that
	id := jdex.ACID{
		Area:     matches[1][0],
		Category: matches[1][1:],
	}
	name := matches[2]

	if id.Area != ctx.lastID.Area {
		return fmt.Errorf("category %q is orphaned in %q",
			matches[1],
			ctx.lastID.AreaString(),
		)
	}

	err := ctx.index.PutCategory(id, name)
	if err != nil {
		return err
	}

	ctx.lastID = id
	return nil
}

func readProcessEntryLine(line string, ctx *readContext) error {
	matches := entryRegex.FindStringSubmatch(line)

	id, err := jdex.ParseACID(matches[1])
	if err != nil {
		return err
	}

	if id.Area != ctx.lastID.Area || id.Category != ctx.lastID.Category {
		return fmt.Errorf("entry %q is orphaned in %q",
			matches[1],
			ctx.lastID.CategoryString(),
		)
	}

	entry := jdex.Entry{
		ID:       id,
		Name:     matches[2],
		Metadata: make(map[string]string),
	}

	err = ctx.index.PutEntry(entry)
	if err != nil {
		return err
	}

	ctx.lastID = id
	ctx.lastEntry = entry
	return nil
}

func readProcessMetadataLine(line string, ctx *readContext) error {
	if ctx.lastID != ctx.lastEntry.ID {
		return errors.New("metadata can only be added to entries")
	}

	matches := metadataRegex.FindStringSubmatch(line)

	key := matches[1]
	value := matches[2]

	ctx.lastEntry.Metadata[key] = value

	err := ctx.index.PutEntry(ctx.lastEntry)
	if err != nil {
		return err
	}

	return nil
}

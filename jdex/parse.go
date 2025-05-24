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

package jdex

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var ErrParse = errors.New("failed to parse jdex")

var commentSingleRegex = regexp.MustCompile(`//.*$`)
var commentMultilineStartRegex = regexp.MustCompile(`/\*.*$`)
var commentMultilineEndRegex = regexp.MustCompile(`^.*\*/`)

var areaRegex = regexp.MustCompile(`^([A-Z0-9]{2}-[A-Z0-9]{2})\s+(.+)$`)
var categoryRegex = regexp.MustCompile(`^([A-Z0-9]{2})\s+(.+)$`)
var entryRegex = regexp.MustCompile(`^([A-Z0-9]{2}\.[A-Z0-9]{2})(\+[A-Z0-9]+)?\s+(.+)$`)
var metadataRegex = regexp.MustCompile(`^-\s*(.+?)\s*:\s*(.+?)\s*$`)

type parseContext struct {
	jdex *Jdex

	multilineComment bool

	currentAreaID     AreaID
	currentCategoryID CategoryID
	currentEntryID    EntryID
}

func ReadJdex(r io.Reader) (jdex Jdex, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	jdex.Areas = make(map[string]Area)

	ctx := parseContext{
		jdex:             &jdex,
		multilineComment: false,
	}

	lineNumber := 0
	for scanner.Scan() {
		if err = scanner.Err(); err != nil {
			return
		}

		lineNumber++

		line := scanner.Text()
		line, err = readJdexPreprocessLine(line, &ctx)
		if err != nil {
			err = fmt.Errorf("line %d: %w", lineNumber, err)
			return
		}

		if line == "" {
			continue
		}

		err = readJdexParseLine(line, &ctx)
		if err != nil {
			err = fmt.Errorf("line %d: %w", lineNumber, err)
			return
		}
	}

	return
}

func readJdexPreprocessLine(input string, ctx *parseContext) (line string, err error) {
	line = input

	// trim whitespace
	line = strings.TrimSpace(line)

	// strip out line comments
	line = commentSingleRegex.ReplaceAllString(line, "")

	// multiline comments
	line = commentMultilineStartRegex.ReplaceAllStringFunc(line, func(comment string) string {
		ctx.multilineComment = true
		return ""
	})
	line = commentMultilineEndRegex.ReplaceAllStringFunc(line, func(comment string) string {
		ctx.multilineComment = false
		return ""
	})

	return
}

func readJdexParseLine(line string, ctx *parseContext) error {
	if areaRegex.MatchString(line) {
		return readJdexParseAreaLine(line, ctx)
	}

	if categoryRegex.MatchString(line) {
		return readJdexParseCategoryLine(line, ctx)
	}

	if entryRegex.MatchString(line) {
		return readJdexParseEntryLine(line, ctx)
	}

	if metadataRegex.MatchString(line) {
		return readJdexParseMetadataLine(line, ctx)
	}

	panic("unimplemented line kind")
}

func readJdexParseAreaLine(line string, ctx *parseContext) (err error) {
	matches := areaRegex.FindStringSubmatch(line)
	area := Area{
		ID:         matches[1],
		Name:       matches[2],
		Categories: make(map[string]Category),
	}

	id, err := NewAreaID(area.ID)
	if err != nil {
		return
	}

	ctx.jdex.Areas[id.String()] = area
	ctx.currentAreaID = id

	return
}

func readJdexParseCategoryLine(line string, ctx *parseContext) (err error) {
	matches := categoryRegex.FindStringSubmatch(line)
	category := Category{
		ID:      matches[1],
		Name:    matches[2],
		Entries: make(map[string]Entry),
	}

	id, err := NewCategoryID(category.ID)
	if err != nil {
		return
	}

	if id.Area != ctx.currentAreaID.Area {
		return fmt.Errorf("category %q is an orphan in area %q",
			id.String(),
			ctx.currentAreaID.String(),
		)
	}

	area := ctx.jdex.Areas[ctx.currentAreaID.String()]
	area.Categories[id.String()] = category
	ctx.currentCategoryID = id

	return
}

func readJdexParseEntryLine(line string, ctx *parseContext) (err error) {
	matches := entryRegex.FindStringSubmatch(line)
	entry := Entry{
		ID:       matches[1],
		Name:     matches[2],
		Metadata: make(map[string]string),
	}

	if len(matches) > 3 {
		entry.ID = matches[1] + matches[2]
		entry.Name = matches[3]
	}

	id, err := NewEntryID(entry.ID)
	if err != nil {
		return
	}

	if id.Area != ctx.currentAreaID.Area || id.Category != ctx.currentCategoryID.Category {
		return fmt.Errorf("entry %q is an orphan in category %q",
			id.String(),
			ctx.currentCategoryID.String(),
		)
	}

	category := ctx.jdex.
		Areas[ctx.currentAreaID.String()].
		Categories[ctx.currentCategoryID.String()]

	category.Entries[id.String()] = entry
	ctx.currentEntryID = id

	return
}

func readJdexParseMetadataLine(line string, ctx *parseContext) (err error) {
	matches := metadataRegex.FindStringSubmatch(line)

	key := matches[1]
	value := matches[2]

	entry := ctx.jdex.
		Areas[ctx.currentAreaID.String()].
		Categories[ctx.currentCategoryID.String()].
		Entries[ctx.currentEntryID.String()]

	entry.Metadata[key] = value

	return
}

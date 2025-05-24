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
	"fmt"
)

var AllowedIDRunes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type IDParseError struct {
	kind string
	id   string
}

func (err *IDParseError) Error() string {
	return fmt.Sprintf("%q is not a valid %s ID", err.id, err.kind)
}

type AreaID struct {
	Area          byte
	CategoryStart byte
	CategoryEnd   byte
}

type CategoryID struct {
	Area     byte
	Category byte
}

type EntryID struct {
	Area     byte
	Category byte
	Entry    string
	Plus     string
}

func NewAreaID(input string) (id AreaID, err error) {
	if len(input) != 5 {
		err = &IDParseError{
			kind: "area",
			id:   input,
		}
		return
	}

	if input[2] != '-' {
		err = &IDParseError{
			kind: "area",
			id:   input,
		}
		return
	}

	if input[0] != input[3] {
		err = &IDParseError{
			kind: "area",
			id:   input,
		}
		return
	}

	id.Area = input[0]
	id.CategoryStart = input[1]
	id.CategoryEnd = input[4]

	return
}

func NewCategoryID(input string) (id CategoryID, err error) {
	if len(input) != 2 {
		err = &IDParseError{
			kind: "category",
			id:   input,
		}
		return
	}

	id.Area = input[0]
	id.Category = input[1]

	return
}

func NewEntryID(input string) (id EntryID, err error) {
	if len(input) < 5 {
		err = &IDParseError{
			kind: "entry",
			id:   input,
		}
	}

	if input[2] != '.' {
		err = &IDParseError{
			kind: "entry",
			id:   input,
		}
		return
	}

	if len(input) > 5 {
		if len(input) == 6 || input[5] != '+' {
			err = &IDParseError{
				kind: "entry",
				id:   input,
			}
		}

		id.Plus = input[6:]
	}

	id.Area = input[0]
	id.Category = input[1]
	id.Entry = input[3:5]

	return
}

func (id *AreaID) String() string {
	return fmt.Sprintf("%c%c-%c%c", id.Area, id.CategoryStart, id.Area, id.CategoryEnd)
}

func (id *CategoryID) String() string {
	return fmt.Sprintf("%c%c", id.Area, id.Category)
}

func (id *EntryID) String() string {
	suffix := ""
	if id.Plus != "" {
		suffix = fmt.Sprintf("+%s", id.Plus)
	}

	return fmt.Sprintf("%c%c.%s%s", id.Area, id.Category, id.Entry, suffix)
}

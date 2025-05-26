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
	"errors"
	"fmt"
	"strings"
)

// Valid characters for AC.IDs.
const ACIDCharset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*
Represents an ID in the system. Can decode `SYS.AC.ID+SUB`.

See https://johnnydecimal.com/10-19-concepts/12-advanced/12.05-acid-notation/
*/
type ACID struct {
	System   string
	Area     byte
	Category string
	Entry    string
	Sub      string
}

var ErrParseACIDBadSeparatorCount = errors.New("ID is expected to have 2 or 3 dot separators")
var ErrACIDInvalidChars = errors.New("ID contains invalid characters")
var ErrACIDRemote = errors.New("ID contains a remote when a local one is needed")

func (id *ACID) String() (str string) {
	str = fmt.Sprintf("%c%s.%s", id.Area, id.Category, id.Entry)

	if id.System != "" {
		str = id.System + "." + str
	}

	if id.Sub != "" {
		str = str + "+" + id.Sub
	}

	return
}

// Returns the area string in the form of `A0-A9`.
func (id *ACID) AreaString() (str string) {
	return fmt.Sprintf("%c0-%c9", id.Area, id.Area)
}

// Returns the category string in the form of `AC`.
func (id *ACID) CategoryString() (str string) {
	return fmt.Sprintf("%c%s", id.Area, id.Category)
}

func ParseACID(input string) (acid ACID, err error) {
	splits := strings.Split(input, ".")

	var ac string
	var id string

	if len(splits) == 2 {
		ac = splits[0]
		id = splits[1]
	} else if len(splits) == 3 {
		acid.System = splits[0]
		ac = splits[1]
		id = splits[2]
	} else {
		err = ErrParseACIDBadSeparatorCount
		return
	}

	acid.Area = ac[0]
	acid.Category = ac[1:]

	plusIndex := strings.IndexByte(id, '+')
	if plusIndex < 0 {
		acid.Entry = id
	} else {
		acid.Entry = id[:plusIndex]
		acid.Sub = id[plusIndex+1:]
	}

	err = acid.Valid()
	return
}

func MustParseACID(input string) (id ACID) {
	id, err := ParseACID(input)
	if err != nil {
		panic(err)
	}

	return
}

func (id *ACID) Valid() (err error) {
	if !strings.ContainsRune(ACIDCharset, rune(id.Area)) {
		err = ErrACIDInvalidChars
	}

	if err == nil {
		err = checkACIDCharset(id.System, id.Category, id.Entry, id.Sub)
	}

	return
}

func (id *ACID) ValidLocal() (err error) {
	if id.System != "" {
		err = ErrACIDRemote
		return
	}

	err = id.Valid()
	return
}

func checkACIDCharset(v ...string) error {
	for _, s := range v {
		for _, c := range s {
			if !strings.ContainsRune(ACIDCharset, c) {
				return ErrACIDInvalidChars
			}
		}
	}

	return nil
}

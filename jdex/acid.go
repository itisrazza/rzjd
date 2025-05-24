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
var ErrParseACIDInvalidChars = errors.New("ID contains invalid characters")

func (acid *ACID) String() (str string) {
	str = fmt.Sprintf("%c%s.%s", acid.Area, acid.Category, acid.Entry)

	if acid.System != "" {
		str = acid.System + "." + str
	}

	if acid.Sub != "" {
		str = str + "+" + acid.Sub
	}

	return
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

	if !strings.ContainsRune(ACIDCharset, rune(acid.Area)) {
		err = ErrParseACIDInvalidChars
	}

	if err == nil {
		err = checkACIDCharset(acid.System, acid.Category, acid.Entry, acid.Sub)
	}

	return
}

func checkACIDCharset(v ...string) error {
	for _, s := range v {
		for _, c := range s {
			if !strings.ContainsRune(ACIDCharset, c) {
				return ErrParseACIDInvalidChars
			}
		}
	}

	return nil
}

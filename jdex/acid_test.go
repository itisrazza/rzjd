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

package jdex_test

import (
	"testing"

	"github.com/itisrazza/rzjd/jdex"
	"github.com/stretchr/testify/assert"
)

var ACIDStringTestCases []struct {
	Name   string
	String string
	ACID   jdex.ACID
} = []struct {
	Name   string
	String string
	ACID   jdex.ACID
}{
	{
		Name:   "Simple",
		String: "12.34",
		ACID: jdex.ACID{
			Area:     '1',
			Category: "2",
			Entry:    "34",
		},
	},
	{
		Name:   "System",
		String: "W01.15.14",
		ACID: jdex.ACID{
			System:   "W01",
			Area:     '1',
			Category: "5",
			Entry:    "14",
		},
	},
	{
		Name:   "Sub",
		String: "11.11+VLD",
		ACID: jdex.ACID{
			Area:     '1',
			Category: "1",
			Entry:    "11",
			Sub:      "VLD",
		},
	},
	{
		Name:   "Fill",
		String: "P01.12.34+VLD",
		ACID: jdex.ACID{
			System:   "P01",
			Area:     '1',
			Category: "2",
			Entry:    "34",
			Sub:      "VLD",
		},
	},
	{
		Name:   "Expansion: Alpha",
		String: "A2.BC",
		ACID: jdex.ACID{
			Area:     'A',
			Category: "2",
			Entry:    "BC",
		},
	},
	{
		Name:   "Expansion: Wider Area",
		String: "WFLX.10",
		ACID: jdex.ACID{
			Area:     'W',
			Category: "FLX",
			Entry:    "10",
		},
	},
}

//
// ACID.String
//

func TestACIDString_TestCases(t *testing.T) {
	for _, testCase := range ACIDStringTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual := testCase.ACID.String()
			assert.Equal(t, testCase.String, actual)
		})
	}
}

//
// ParseACID
//

func testParseACIDFailure(t *testing.T, toParse string, expectedErr error) {
	_, actualErr := jdex.ParseACID(toParse)
	if !assert.Error(t, actualErr) {
		t.FailNow()
	}

	assert.Equal(t, expectedErr, actualErr)
}

func TestParseACID_TestCases(t *testing.T) {
	for _, testCase := range ACIDStringTestCases {
		t.Run(testCase.Name, func(t *testing.T) {
			actual, err := jdex.ParseACID(testCase.String)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			assert.Equal(t, testCase.ACID, actual)
		})
	}
}

func TestParseACID_TooFewDots(t *testing.T) {
	testParseACIDFailure(t, "11", jdex.ErrParseACIDBadSeparatorCount)
}

func TestParseACID_BadAreaChar(t *testing.T) {
	testParseACIDFailure(t, "Ă1.23", jdex.ErrParseACIDInvalidChars)
}

func TestParseACID_BadChar(t *testing.T) {
	testParseACIDFailure(t, "1Ă.23", jdex.ErrParseACIDInvalidChars)
}

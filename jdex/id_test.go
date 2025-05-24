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

func TestNewAreaID(t *testing.T) {
	id, err := jdex.NewAreaID("10-19")
	assert.NoError(t, err)

	assert.Equal(t, jdex.AreaID{
		Area:          '1',
		CategoryStart: '0',
		CategoryEnd:   '9',
	}, id)
}

func TestAreaIDString(t *testing.T) {
	id := jdex.AreaID{
		Area:          '1',
		CategoryStart: '0',
		CategoryEnd:   '9',
	}

	assert.Equal(t, "10-19", id.String())
}

//

func TestNewCategoryID(t *testing.T) {
	id, err := jdex.NewCategoryID("12")
	assert.NoError(t, err)

	assert.Equal(t, jdex.CategoryID{
		Area:     '1',
		Category: '2',
	}, id)
}

func TestCategoryIDString(t *testing.T) {
	id := jdex.CategoryID{
		Area:     '1',
		Category: '2',
	}

	assert.Equal(t, "12", id.String())
}

//

func TestNewEntryID(t *testing.T) {
	id, err := jdex.NewEntryID("12.13")
	assert.NoError(t, err)

	assert.Equal(t, jdex.EntryID{
		Area:     '1',
		Category: '2',
		Entry:    "13",
	}, id)
}

func TestNewEntryIDPlus(t *testing.T) {
	id, err := jdex.NewEntryID("12.13+ABC")
	assert.NoError(t, err)

	assert.Equal(t, jdex.EntryID{
		Area:     '1',
		Category: '2',
		Entry:    "13",
		Plus:     "ABC",
	}, id)
}

func TestEntryIDString(t *testing.T) {
	id := jdex.EntryID{
		Area:     '1',
		Category: '2',
		Entry:    "13",
	}

	assert.Equal(t, "12.13", id.String())
}

func TestEntryIDWithPlusString(t *testing.T) {
	id := jdex.EntryID{
		Area:     '1',
		Category: '2',
		Entry:    "13",
		Plus:     "ABC",
	}

	assert.Equal(t, "12.13+ABC", id.String())
}

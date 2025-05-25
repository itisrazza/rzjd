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

func Test_NewIndex_HasSystemIndex(t *testing.T) {
	index, err := jdex.NewIndex()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	indexID := jdex.MustParseACID("00.00")

	area, err := index.AreaName(indexID)
	assert.NoError(t, err)

	category, err := index.CategoryName(indexID)
	assert.NoError(t, err)

	entry, err := index.Entry(indexID)
	assert.NoError(t, err)

	assert.Equal(t, "System", area)
	assert.Equal(t, "Index", category)
	assert.Equal(t, jdex.Entry{
		ID: jdex.ACID{
			Area:     '0',
			Category: "0",
			Entry:    "00",
		},
		Name: "System Index",
		Metadata: map[string]string{
			"Format": "jdex",
		},
	}, entry)

	if !assert.NoError(t, err) {
		t.FailNow()
	}
}

func Test_Index_PutEntry_FailRemoteID(t *testing.T) {
	id := jdex.MustParseACID("W01.11.11")

	index, _ := jdex.NewIndex()
	index.PutArea(id, "")
	index.PutCategory(id, "")

	err := index.PutEntry(jdex.Entry{ID: id})
	assert.ErrorIs(t, err, jdex.ErrInvalidID)
	assert.ErrorIs(t, err, jdex.ErrACIDRemote)
}

func Test_Index_PutEntry_FailNoArea(t *testing.T) {
	index, _ := jdex.NewIndex()
	err := index.PutEntry(jdex.Entry{
		ID: jdex.MustParseACID("11.11"),
	})

	assert.ErrorIs(t, err, jdex.ErrAreaNotFound)
}

func Test_Index_PutEntry_FailNoCategory(t *testing.T) {
	index, _ := jdex.NewIndex()
	id := jdex.MustParseACID("11.11")

	index.PutArea(id, "")
	err := index.PutEntry(jdex.Entry{ID: id})

	assert.ErrorIs(t, err, jdex.ErrCategoryNotFound)
}

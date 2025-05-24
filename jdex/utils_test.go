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

import "github.com/itisrazza/rzjd/jdex"

func buildJdex(areas []jdex.Area) (index jdex.Jdex) {
	index.Areas = make(map[string]jdex.Area)

	for _, area := range areas {
		index.Areas[area.ID] = area
	}

	return
}

func buildArea(id, name string, categories []jdex.Category) (area jdex.Area) {
	area.ID = id
	area.Name = name
	area.Categories = make(map[string]jdex.Category)

	for _, category := range categories {
		area.Categories[category.ID] = category
	}

	return
}

func buildCategory(id, name string, entries []jdex.Entry) (category jdex.Category) {
	category.ID = id
	category.Name = name
	category.Entries = make(map[string]jdex.Entry)

	for _, entry := range entries {
		category.Entries[entry.ID] = entry
	}

	return
}

func buildEntry(id, name string, metadata map[string]string) (entry jdex.Entry) {
	entry.ID = id
	entry.Name = name
	entry.Metadata = metadata

	return
}

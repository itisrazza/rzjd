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
	"slices"
)

// Index is the entry database. It stores the entries, their names and
// immediate metadata.
//
// This data is stored in "00.00 System Index" within the system.
type Index struct {
	entries map[string]Entry
	areas   map[byte]indexArea
}

// Represents a single entry in the system.
type Entry struct {
	ID       ACID              // Entry's AC.ID.
	Name     string            // Entry's name.
	Metadata map[string]string // Entry's immediate metadata.
}

type indexArea struct {
	name       string
	categories map[string]indexCategory
}

type indexCategory struct {
	name    string
	entries map[string]bool
}

// These AC.IDs are reserved by the system. Users should not edit these
// manually.
var ProtectedACIDs = []string{
	"00.00", // system index
}

var ErrInvalidID = errors.New("entry ID is invalid")
var ErrProtectedID = errors.New("entry ID is used by the system")

var ErrEntryNotFound = errors.New("entry does not exist")
var ErrCategoryNotFound = errors.New("category does not exist")
var ErrAreaNotFound = errors.New("area does not exist")

// Creates a new index
func NewIndex() (Index, error) {
	index := Index{
		entries: make(map[string]Entry),
		areas:   make(map[byte]indexArea),
	}

	indexID := MustParseACID("00.00")

	index.PutArea(indexID, "System")
	index.PutCategory(indexID, "Index")
	index.putEntryUnsafe(Entry{
		ID:   indexID,
		Name: "System Index",
		Metadata: map[string]string{
			"Format": "jdex",
		},
	})

	return index, nil
}

func (index *Index) Entry(id ACID) (entry Entry, err error) {
	if err = id.ValidLocal(); err != nil {
		err = errors.Join(ErrInvalidID, err)
		return
	}

	entry, ok := index.entries[id.String()]
	if !ok {
		err = ErrEntryNotFound
		return
	}

	return
}

func (index *Index) AreaName(id ACID) (name string, err error) {
	if err = id.ValidLocal(); err != nil {
		err = errors.Join(ErrInvalidID, err)
		return
	}

	area, ok := index.areas[id.Area]
	if !ok {
		err = ErrAreaNotFound
		return
	}

	return area.name, nil
}

func (index *Index) CategoryName(id ACID) (name string, err error) {
	if err = id.ValidLocal(); err != nil {
		err = errors.Join(ErrInvalidID, err)
		return
	}

	area, ok := index.areas[id.Area]
	if !ok {
		err = ErrAreaNotFound
		return
	}

	category, ok := area.categories[id.Category]
	if !ok {
		err = ErrCategoryNotFound
		return
	}

	return category.name, nil
}

func (index *Index) PutArea(id ACID, name string) error {
	if err := id.ValidLocal(); err != nil {
		return errors.Join(ErrInvalidID, err)
	}

	area, ok := index.areas[id.Area]
	if !ok {
		area = indexArea{
			name:       name,
			categories: make(map[string]indexCategory),
		}
	} else {
		area.name = name
	}

	index.areas[id.Area] = area
	return nil
}

func (index *Index) PutCategory(id ACID, name string) error {
	if err := id.ValidLocal(); err != nil {
		return errors.Join(ErrInvalidID, err)
	}

	area, ok := index.areas[id.Area]
	if !ok {
		return ErrAreaNotFound
	}

	category, ok := area.categories[id.Category]
	if !ok {
		category = indexCategory{
			name:    name,
			entries: map[string]bool{},
		}
	} else {
		category.name = name
	}

	area.categories[id.Category] = category
	return nil
}

func (index *Index) putEntryUnsafe(entry Entry) (err error) {
	id := entry.ID

	if err = id.ValidLocal(); err != nil {
		return errors.Join(ErrInvalidID, err)
	}

	area, ok := index.areas[id.Area]
	if !ok {
		err = ErrAreaNotFound
		return
	}

	category, ok := area.categories[id.Category]
	if !ok {
		err = ErrCategoryNotFound
		return
	}

	category.entries[id.String()] = true
	index.entries[id.String()] = entry

	return
}

func (index *Index) PutEntry(entry Entry) error {
	if IsProtectedACID(entry.ID) {
		return ErrProtectedID
	}

	return index.putEntryUnsafe(entry)
}

func IsProtectedACID(id ACID) bool {
	return slices.Contains(ProtectedACIDs, id.String())
}

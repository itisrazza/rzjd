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

package jdfs

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/itisrazza/rzjd/jdex"
)

/*
 */
type Store struct {
	Root  string     // Path to where the store is located.
	Index jdex.Index // Pointer to index to use for name lookup.
}

var ErrPathNotDir = errors.New("path is not a directory")
var ErrEntryNotFound = errors.New("entry was not found in index")

func newStoreImpl(path string) (store *Store, err error) {
	pathInfo, err := os.Stat(path)
	if err != nil || !pathInfo.IsDir() {
		err = errors.Join(ErrPathNotDir, err)
		return
	}

	index, err := jdex.NewIndex()
	if err != nil {
		return
	}

	return &Store{
		Root:  path,
		Index: index,
	}, nil
}

// Create a new store at the given path.
func NewStore(path string) (store *Store, err error) {
	store, err = newStoreImpl(path)
	if err != nil {
		return
	}

	panic("not implemented")
}

// Open the store at the given path.
func OpenStore(path string) (store *Store, err error) {
	panic("not implemented")

	// TODO: load empty index with only the index reference
	// TODO: get path to "00.00"
	// TODO: load file into the index
	// TODO: ???
	// TODO: profit
}

func (store *Store) AreaPath(id jdex.ACID) (path_ string, err error) {
	panic("not implemented")
}

func (store *Store) CategoryPath(id jdex.ACID) (path_ string, err error) {
	categoryName, err := store.Index.CategoryName(id)
	if err != nil {
		err = errors.Join(ErrEntryNotFound, err)
		return
	}

	areaPath, err := store.AreaPath(id)
	if err != nil {
		return
	}

	path_ = path.Join(areaPath,
		fmt.Sprintf("%s %s",
			id.CategoryString(),
			TransformFilename(categoryName),
		),
	)

	return
}

func (store *Store) EntryPath(id jdex.ACID) (path_ string, err error) {
	entry, err := store.Index.Entry(id)
	if err != nil {
		err = errors.Join(ErrEntryNotFound, err)
		return
	}

	categoryPath, err := store.CategoryPath(id)
	if err != nil {
		return
	}

	path_ = path.Join(categoryPath, EntryFilename(entry))

	pathInfo, err := os.Stat(path_)
	if err != nil || !pathInfo.IsDir() {
		err = errors.Join(ErrPathNotDir, err)
		return
	}

	return
}

func EntryFilename(entry jdex.Entry) string {
	return fmt.Sprintf("%s %s",
		entry.ID.String(),
		TransformFilename(entry.Name),
	)
}

func TransformFilename(text string) string {
	return strings.Map(func(r rune) rune {
		const badChars = `<>:"/\|?$*`
		if strings.ContainsRune(badChars, r) {
			return '_'
		}

		return r
	}, text)
}

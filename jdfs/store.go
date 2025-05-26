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
	"github.com/itisrazza/rzjd/jdex/jdexfile"
)

/*
 */
type Store struct {
	Root  string     // Path to where the store is located.
	Index jdex.Index // Pointer to index to use for name lookup.
}

var ErrPathNotDir = errors.New("path is not a directory")
var ErrEntryNotFound = errors.New("entry was not found in index")

const EntryIndexFilename = "Index.txt"

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

	indexPath, err := store.IndexPath()
	if err != nil {
		err = fmt.Errorf("failed to create index file: %w", err)
		return
	}

	indexFile, err := CreateWithParents(indexPath)
	if err != nil {
		err = fmt.Errorf("failed to create index file: %w", err)
		return
	}
	defer indexFile.Close()

	err = jdexfile.Write(&store.Index, indexFile)
	if err != nil {
		err = fmt.Errorf("failed to create index file: %w", err)
		return
	}

	return
}

// Open the store at the given path.
func OpenStore(path string) (store *Store, err error) {
	store, err = newStoreImpl(path)
	if err != nil {
		return
	}

	indexPath, err := store.IndexPath()
	if err != nil {
		return
	}

	indexFile, err := os.Open(indexPath)
	if err != nil {
		return
	}
	defer indexFile.Close()

	store.Index, err = jdexfile.Read(indexFile)
	if err != nil {
		return
	}

	return
}

// Get the path to the system index file.
func (store *Store) IndexPath() (entryPath string, err error) {
	return store.EntryIndexPath(jdex.MustParseACID("00.00"))
}

// Get the path to the area directory.
func (store *Store) AreaPath(id jdex.ACID) (areaPath string, err error) {
	areaName, err := store.Index.AreaName(id)
	if err != nil {
		err = errors.Join(ErrEntryNotFound, err)
		return
	}

	return path.Join(store.Root, fmt.Sprintf("%s %s", id.AreaString(), areaName)), nil
}

// Get the path to the category directory.
func (store *Store) CategoryPath(id jdex.ACID) (path_ string, err error) {
	areaPath, err := store.AreaPath(id)
	if err != nil {
		return
	}

	categoryName, err := store.Index.CategoryName(id)
	if err != nil {
		err = errors.Join(ErrEntryNotFound, err)
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

// Get the path to the entry directory.
func (store *Store) EntryPath(id jdex.ACID) (entryPath string, err error) {
	categoryPath, err := store.CategoryPath(id)
	if err != nil {
		return
	}

	entry, err := store.Index.Entry(id)
	if err != nil {
		err = errors.Join(ErrEntryNotFound, err)
		return
	}

	entryPath = path.Join(categoryPath, EntryFilename(entry))
	return
}

func (store *Store) EntryIndexPath(id jdex.ACID) (entryIndexPath string, err error) {
	entryPath, err := store.EntryPath(id)
	if err != nil {
		return
	}

	entryIndexPath = path.Join(entryPath, EntryIndexFilename)
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

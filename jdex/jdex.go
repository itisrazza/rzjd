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

// Implements reading and writing to the index with the following format:
//   https://github.com/johnnydecimal/index-spec

type Jdex struct {
	Areas map[string]Area
}

type Area struct {
	ID         string
	Name       string
	Categories map[string]Category
}

type Category struct {
	ID      string
	Name    string
	Entries map[string]Entry
}

type Entry struct {
	ID       string
	Name     string
	Metadata map[string]string
}

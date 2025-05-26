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

package jdexfile

import (
	"fmt"
	"io"

	"github.com/itisrazza/rzjd/jdex"
)

func Write(index *jdex.Index, w io.Writer) (err error) {
	for _, areaID := range index.AreaIndexes() {
		areaName, _ := index.AreaName(areaID)

		_, err = fmt.Fprintf(w, "%s %s\n", areaID.AreaString(), areaName)
		if err != nil {
			return
		}

		categoryIndexes, _ := index.Categories(areaID)
		for _, categoryID := range categoryIndexes {
			categoryName, _ := index.CategoryName(categoryID)

			_, err = fmt.Fprintf(w, "  %s %s\n", categoryID.CategoryString(), categoryName)
			if err != nil {
				return
			}

			entries, _ := index.Entries(categoryID)
			for _, entryID := range entries {
				entry, _ := index.Entry(entryID)

				_, err = fmt.Fprintf(w, "    %s %s\n", entry.ID.String(), entry.Name)
				if err != nil {
					return
				}

				if entry.Metadata != nil {
					for key, value := range entry.Metadata {
						_, err = fmt.Fprintf(w, "      - %s: %s", key, value)
						if err != nil {
							return
						}
					}
				}
			}
		}
	}

	return nil
}

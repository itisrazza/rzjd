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

package main

import (
	"fmt"

	"github.com/itisrazza/rzjd/jdex"
)

type EditCmd struct {
	ID string `arg:"" help:"ID of the entry to edit"`
}

func (cmd *EditCmd) Run() error {
	id, err := jdex.ParseACID(cmd.ID)
	if err != nil {
		return err
	}

	fmt.Printf("opening editor to %#v\n", id)

	return nil
}

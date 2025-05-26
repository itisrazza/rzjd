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

package rzinteractive

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/huh"
)

var ErrCancel = errors.New("cancelled by user")

func NewStorePrompt(storePath string) error {
	confirm := false

	form := newForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Create a new store?").
				Description(fmt.Sprintf("A new Johnny.Decimal system will be created in \"%s\".", storePath)).
				Value(&confirm),
		),
	)

	if err := form.Run(); err != nil {
		return err
	}

	if !confirm {
		return ErrCancel
	}

	return nil
}

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
	"errors"
	"os"
	"path"

	"github.com/adrg/xdg"
	"github.com/itisrazza/rzjd/jdfs"
	"github.com/itisrazza/rzjd/rzinteractive"
)

func fullStorePath() (string, error) {
	if cli.Store != nil {
		return *cli.Store, nil
	}

	storePath, ok := os.LookupEnv("RZJD_STORE")
	if ok {
		return storePath, nil
	}

	return path.Join(xdg.UserDirs.Documents, "rzjd"), nil
}

func OpenOrCreateStore() (*jdfs.Store, error) {
	storePath, err := fullStorePath()
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(storePath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		if cli.NonInteractive {
			return nil, err
		}

		err = rzinteractive.NewStorePrompt(storePath)
		if err != nil {
			return nil, err
		}

		err = os.MkdirAll(storePath, 0755)
		if err != nil {
			return nil, err
		}

		store, err := jdfs.NewStore(storePath)
		return store, err
	} else if err != nil {
		return nil, err
	}

	return jdfs.OpenStore(storePath)
}

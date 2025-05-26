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
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/itisrazza/rzjd/jdex"
	"github.com/itisrazza/rzjd/jdfs"
)

type EditCmd struct {
	ID   string  `arg:"" help:"ID of the entry to edit"`
	Name *string `arg:"" optional:"" help:"Providing a name will rename the entry."`

	Editor *string `short:"e" type:"path" help:"Path to text editor."`
}

func (cmd *EditCmd) Run() error {
	id, err := jdex.ParseACID(cmd.ID)
	if err != nil {
		return err
	}

	// if jdex.IsProtectedACID(id) {
	// 	return fmt.Errorf("%q is a protected ID", id.String())
	// }

	store, err := OpenOrCreateStore()
	if err != nil {
		return err
	}

	entryPath, err := store.EntryPath(id)
	if err != nil {
		return err
	}

	err = os.MkdirAll(entryPath, 0755)
	if err != nil {
		return err
	}

	return cmd.openEditor(path.Join(entryPath, jdfs.EntryIndexFilename))
}

func (cmd *EditCmd) openEditor(path string) error {
	editorName, err := cmd.editorName()
	if err != nil {
		return fmt.Errorf("couldn't find a suitable text editor: %w", err)
	}

	editorCmd := exec.Command(editorName, path)
	editorCmd.Stdin = os.Stdin
	editorCmd.Stdout = os.Stdout
	editorCmd.Stderr = os.Stderr

	err = editorCmd.Run()
	if err != nil {
		return fmt.Errorf("editor failed to run: %w", err)
	}

	return nil
}

func (cmd *EditCmd) editorName() (string, error) {
	if cmd.Editor != nil {
		return *cmd.Editor, nil
	}

	editorName, ok := os.LookupEnv("RZJD_EDITOR")
	if ok {
		return editorName, nil
	}

	editorName, ok = os.LookupEnv("EDITOR")
	if ok {
		return editorName, nil
	}

	if runtime.GOOS == "windows" {
		return "notepad.exe", nil
	}

	return exec.LookPath("nano")
}

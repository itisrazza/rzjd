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
	"github.com/alecthomas/kong"
)

var cli struct {
	Store          *string `short:"s" type:"path" default:"$RZJD_STORE" help:"Path to where your system is stored."`
	NonInteractive bool    `default:"false" help:"Fail instead of interactively solving issues."`

	New     NewCmd     `cmd:"" help:"Create a new store."`
	Explore ExploreCmd `cmd:"" default:"true" help:"Explore your store interactively."`
	View    ViewCmd    `cmd:"" help:"View an entry in the store."`
	Edit    EditCmd    `cmd:"" help:"Edit an entry in the store."`
	Archive ArchiveCmd `cmd:"" help:"Archive an entry."`
	Setup   SetupCmd   `cmd:"" help:"Set up rzjd in your environment."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}

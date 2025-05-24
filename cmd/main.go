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

var CLI struct {
	Store          *string `short:"C" type:"path" default:"$RZJD_STORE" help:"Path to store."`
	NonInteractive bool    `default:"false" help:"Fail instead of interactively solving issues."`

	New     struct{} `cmd:"" help:"Create a new store."`
	Explore struct{} `cmd:"" help:"Explore your store interactively."`
	View    struct{} `cmd:"" help:"View an entry in the store."`
	Edit    struct{} `cmd:"" help:"Edit an entry in the store."`
	Archive struct{} `cmd:"" help:"Archive an entry."`

	Setup struct {
		Shell struct{} `cmd:"" help:"Output shell initialisation script."`
	} `cmd:"" help:"Set up rzjd in your environment."`
}

func main() {
	ctx := kong.Parse(&CLI)

	switch ctx.Command() {
	default:
		panic(ctx.Command())
	}
}

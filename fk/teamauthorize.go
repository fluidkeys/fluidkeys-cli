// Copyright 2019 Paul Furley and Ian Drysdale
//
// This file is part of Fluidkeys Client which makes it simple to use OpenPGP.
//
// Fluidkeys Client is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Fluidkeys Client is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with Fluidkeys Client.  If not, see <https://www.gnu.org/licenses/>.

package fk

import (
	"github.com/fluidkeys/fluidkeys/out"
	"github.com/fluidkeys/fluidkeys/ui"
)

func teamAuthorize() exitCode {
	out.Print(ui.FormatInfo(
		"Authorizing a key adds it to the team roster",
		[]string{
			"By authorizing a key, everyone in your team will fetch and trust that key.",
		},
	))

	printHeader("Authorize keys")

	out.Print(ui.FormatFailure("Not implemented", nil, nil))
	return 1
}

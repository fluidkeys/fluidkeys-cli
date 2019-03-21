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
	"log"
	"time"

	docopt "github.com/docopt/docopt-go"
	"github.com/fluidkeys/fluidkeys/out"
	"github.com/fluidkeys/fluidkeys/status"
	"github.com/fluidkeys/fluidkeys/table"
)

func statusSubcommand(args docopt.Opts) exitCode {
	out.Print("\n")

	groupedMemberships, err := user.GroupedMemberships()
	if err != nil {
		log.Panic(err)
	}

	for _, groupedMembership := range groupedMemberships {
		printHeader(groupedMembership.Team.Name)
		teamKeysWithWarnings := []table.KeyWithWarnings{}

		for _, membership := range groupedMembership.Memberships {
			key, err := loadPgpKey(membership.Me.Fingerprint)
			if err != nil {
				return 1
			}

			keyWithWarnings := table.KeyWithWarnings{
				Key:      key,
				Warnings: status.GetKeyWarnings(*key, &Config),
			}

			teamKeysWithWarnings = append(teamKeysWithWarnings, keyWithWarnings)

		}
		out.Print(table.FormatKeyTable(teamKeysWithWarnings))

		peopleRows := []table.PersonRow{}
		for _, person := range groupedMembership.Team.People {
			lastFetched, err := db.GetLast("fetch", person.Fingerprint)
			if err != nil {
				continue
			}
			peopleRows = append(
				peopleRows, table.PersonRow{
					Email:              person.Email,
					IsAdmin:            person.IsAdmin,
					TimeSinceLastFetch: time.Since(lastFetched),
				},
			)
		}

		out.Print(table.FormatPeopleTable(peopleRows))
	}

	return 0
}
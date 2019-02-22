package team

import (
	"strings"
	"testing"

	"github.com/fluidkeys/fluidkeys/assert"
	"github.com/fluidkeys/fluidkeys/fingerprint"
	"github.com/gofrs/uuid"
)

func TestParse(t *testing.T) {
	reader := strings.NewReader(validRoster)
	team, err := parse(reader)

	assert.ErrorIsNil(t, err)
	expectedPeople := []Person{
		{
			Email:       "paul@fluidkeys.com",
			Fingerprint: fingerprint.MustParse("B79F0840DEF12EBBA72FF72D7327A44C2157A758"),
		},
		{
			Email:       "ian@fluidkeys.com",
			Fingerprint: fingerprint.MustParse("E63AF0E74EB5DE3FB72DC981C991709318ECBDE7"),
		},
	}
	assert.Equal(t, expectedPeople, team.People)

	assert.Equal(t, uuid.Must(uuid.FromString("38be2a70-23d8-11e9-bafd-7f97f2e239a3")), team.UUID)
	assert.Equal(t, "Fluidkeys CIC", team.Name)
}

const validRoster = `# Fluidkeys team roster

uuid = "38be2a70-23d8-11e9-bafd-7f97f2e239a3"
name = "Fluidkeys CIC"

[[person]]
email = "paul@fluidkeys.com"
fingerprint = "B79F 0840 DEF1 2EBB A72F  F72D 7327 A44C 2157 A758"

[[person]]
email = "ian@fluidkeys.com"
fingerprint = "E63A F0E7 4EB5 DE3F B72D  C981 C991 7093 18EC BDE7"
`

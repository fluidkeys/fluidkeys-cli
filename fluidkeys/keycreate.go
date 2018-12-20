// Copyright 2018 Paul Furley and Ian Drysdale
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

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fluidkeys/fluidkeys/backupzip"
	"github.com/fluidkeys/fluidkeys/colour"
	"github.com/fluidkeys/fluidkeys/emailutils"
	"github.com/fluidkeys/fluidkeys/fingerprint"
	"github.com/fluidkeys/fluidkeys/humanize"
	"github.com/fluidkeys/fluidkeys/out"
	"github.com/fluidkeys/fluidkeys/pgpkey"
	spin "github.com/tj/go-spin"

	"github.com/sethvargo/go-diceware/diceware"
)

const DicewareNumberOfWords int = 6
const DicewareSeparator string = "."

type DicewarePassword struct {
	words     []string
	separator string
}

func (d DicewarePassword) AsString() string {
	return strings.Join(d.words, d.separator)
}

type generatePgpKeyResult struct {
	pgpKey *pgpkey.PgpKey
	err    error
}

func keyCreate() exitCode {

	if !gpg.IsWorking() {
		out.Print(colour.Warning("\nGPG isn't working on your system 🤒\n\n"))
		out.Print("You can still use FluidKeys to make a key and then " +
			"later import it from your backup.\n\n" +
			"Alternatively, quit now [ctrl-c], install GPG then " +
			"run FluidKeys again.\n\n")
		promptForInput("Press enter to continue. ")
	}
	out.Print("\n")

	printHeader("What's your email address?")

	out.Print("This is how other people using Fluidkeys will find you.\n\n")
	out.Print("We'll send you an email to verify your address.\n\n")

	email := promptForInput("[email] : ")
	for !emailutils.RoughlyValidateEmail(email) {
		printWarning("Not a valid email address")
		out.Print("\n")
		email = promptForInput("[email] : ")
	}

	channel := make(chan generatePgpKeyResult)
	go generatePgpKey(email, channel)

	password := generatePassword(DicewareNumberOfWords, DicewareSeparator)

	out.Print("Your key will be protected with this password:\n\n")
	displayPassword(password)
	if !userConfirmedRandomWord(password) {
		out.Print("Those words did not match. Here it is again:\n\n")
		displayPassword(password)
		if !userConfirmedRandomWord(password) {
			out.Print("Those words didn't match again. Quitting...\n")
			os.Exit(1)
		}
	}

	out.Print("Creating key for " + colour.Info(email) + ":\n\n")

	generateJob := <-channel

	if generateJob.err != nil {
		log.Panicf("Failed to generate key: %v", generateJob.err)
	}
	printSuccessfulAction("Generate key for " + email)

	pushPrivateKeyBackToGpg(generateJob.pgpKey, password.AsString(), &gpg)
	printSuccessfulAction("Store key in " + colour.Info("gpg"))

	fingerprint := generateJob.pgpKey.Fingerprint()
	db.RecordFingerprintImportedIntoGnuPG(fingerprint)
	if err := tryEnableMaintainAutomatically(generateJob.pgpKey, password.AsString()); err == nil {
		printSuccessfulAction("Store password in " + Keyring.Name())
		printSuccessfulAction("Setup automatic maintenance using " + colour.Info("cron"))
	} else {
		printFailedAction("Setup automatic maintenance")
	}

	filename, err := backupzip.OutputZipBackupFile(fluidkeysDirectory, generateJob.pgpKey, password.AsString())
	if err != nil {
		printFailedAction("Make a backup ZIP file")
	}
	directory, _ := filepath.Split(filename)
	printSuccessfulAction("Make a backup ZIP file in")
	out.Print("        " + colour.Info(directory) + "\n\n")

	printSuccess("Successfully created key for " + email)
	out.Print("\n")

	promptToEnableConfigPublishToAPI(generateJob.pgpKey)

	if Config.ShouldPublishToAPI(generateJob.pgpKey.Fingerprint()) {
		err := publishKeyToAPI(generateJob.pgpKey)
		if err != nil {
			printFailed("Failed to publish key")
			out.Print(err.Error())
		} else {
			printSuccess("Successfully published key")
		}
	}
	out.Print("\n")

	return 0
}

func generatePgpKey(email string, channel chan generatePgpKeyResult) {
	key, err := pgpkey.Generate(email, time.Now(), nil)

	channel <- generatePgpKeyResult{key, err}
}

func generatePassword(numberOfWords int, separator string) DicewarePassword {
	return DicewarePassword{
		words:     diceware.MustGenerate(numberOfWords),
		separator: separator,
	}
}

func displayPassword(password DicewarePassword) {
	out.Print(out.NoLogCharacter + "   " + colour.Info(password.AsString()) + "\n\n")
	out.Print("If you use a password manager, save it there now.\n\n")
	out.Print(colour.Warning("Store this safely, otherwise you won’t be able to use your key\n\n"))

	promptForInput("Press enter when you've stored it safely. ")
}

func userConfirmedRandomWord(password DicewarePassword) bool {
	clearScreen()
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(password.words))
	correctWord := password.words[randomIndex]
	wordOrdinal := humanize.Ordinal(randomIndex + 1)

	out.Print(fmt.Sprintf("Enter the %s word from your password\n\n", wordOrdinal))
	givenWord := promptForInput("[" + wordOrdinal + " word] : ")
	return givenWord == correctWord
}

func clearScreen() {
	out.Print("\033[H\033[2J")
}

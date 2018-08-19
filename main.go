package main

import (
	"fmt"
	"gopkg.in/satori/go.uuid.v1"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 3 {
		usage()
	}
	text := os.Args[1]
	if text == "-h" || text == "--help" || text == "help" {
		usage()
	}
	var reverse bool
	if text == "-r" || text == "--reverse" {
		reverse = true
	}
	if reverse {
		text = os.Args[2]
	}
	fmt.Println(formatUUID(text, reverse))
}

func formatUUID(text string, reverse bool) string {
	u, err := uuid.FromString(strings.ToLower(normalize(text)))
	if err != nil {
		fmt.Printf("Provided %s text does not look like UUID %v\n", text, err)
		os.Exit(1)
	}
	result := u.String()
	if reverse {
		return strings.ToUpper(normalize(result))
	}
	return result
}

func normalize(text string) string {
	return strings.Replace(text, "-", "", -1)
}

func usage() {
	fmt.Println(
		"Formats UUID into the canonical presentation, " +
			"ex.: 3A2DD5E0D2C04F13A3E2F600C9530793 -> 3a2dd5e0-d2c0-4f13-a3e2-f600c9530793.\n" +
			"Usage: uuidfmt [-r/--reverse] <uuid>\n" +
			"    -r reverse, i.e. do the opposite of the formatting to canonical form, " +
			"i.e. 3a2dd5e0-d2c0-4f13-a3e2-f600c9530793 -> 3A2DD5E0D2C04F13A3E2F600C9530793")
	os.Exit(0)
}

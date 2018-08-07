package main

import (
	"os"
	"strings"
	"gopkg.in/satori/go.uuid.v1"
	"fmt"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) != 2 {
		usage()
	}
	text := os.Args[1]
	if text == "-h" || text == "--help" || text == "help" {
		usage()
	}
	u := strings.ToLower(text)
	fmt.Println(formatUUID(u))
}

func formatUUID(text string) string {
	u, err := uuid.FromString(strings.ToLower(strings.Replace(text, "-", "", -1)))
	if err != nil {
		fmt.Printf("Provided %s text does not look like UUID %v\n", text, err)
		os.Exit(1)
	}
	return u.String()
}

func usage() {
	fmt.Println(
		"Formats UUID into the canonical presentation, " +
			"ex.: 3A2DD5E0D2C04F13A3E2F600C9530793 -> 3a2dd5e0-d2c0-4f13-a3e2-f600c9530793.\n" +
			"Usage: uuidfmt <uuid>")
	os.Exit(0)
}

package main

import "testing"

func TestUUIDFmt(t *testing.T) {
	actual := "3A2DD5E0D2C04F13A3E2F600C9530793"
	uuid := formatUUID(actual)
	if uuid != "3a2dd5e0-d2c0-4f13-a3e2-f600c9530793" {
		t.Error("UUID format does not match the expected form", actual, uuid)
	}
}

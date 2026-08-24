//go:build integration

package main

import (
	"regexp"
	"testing"
)

var versionFormatRe = regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`)

func TestGetLatestVersion(t *testing.T) {
	got, err := GetLatestVersion(t.Context())
	if err != nil {
		t.Fatalf("GetLatestVersion returned an error: %v", err)
	}

	if !versionFormatRe.MatchString(got) {
		t.Fatalf("GetLatestVersion() output = %q, want an x.y.z version", got)
	}
}

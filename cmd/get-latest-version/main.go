package main

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"os"
	"os/signal"
	"regexp"
	"slices"
	"syscall"

	"cloud.google.com/go/storage"
	"golang.org/x/mod/semver"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var versionRe = regexp.MustCompile(`.*-(\d+\.\d+\.\d+).*`)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	latestVersion, err := GetLatestVersion(ctx)
	if err != nil {
		cancel()
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(2)
	}

	fmt.Fprintln(os.Stdout, latestVersion)
}

func GetLatestVersion(ctx context.Context) (string, error) {
	client, err := storage.NewClient(ctx, option.WithoutAuthentication())
	if err != nil {
		return "", fmt.Errorf("failed to create storage client: %w", err)
	}

	latestVersion, err := getLatestVersion(ctx, client)
	if err != nil {
		return "", err
	}

	return latestVersion, nil
}

func getLatestVersion(ctx context.Context, client *storage.Client) (string, error) {
	// https://console.cloud.google.com/storage/browser/cloud-sdk-release
	var q storage.Query
	q.MatchGlob = "*-linux-x86_64.tar.gz"
	if err := q.SetAttrSelection([]string{"Name"}); err != nil {
		return "", fmt.Errorf("failed to set attribute selection: %w", err)
	}

	var names []string

	it := client.Bucket("cloud-sdk-release").Objects(ctx, &q)
	for {
		attrs, err := it.Next()
		if errors.Is(err, iterator.Done) {
			break
		}

		if err != nil {
			return "", fmt.Errorf("failed to get object: %w", err)
		}

		names = append(names, attrs.Name)
	}

	versionsMap := make(map[string]struct{}, len(names))
	for _, name := range names {
		sub := versionRe.FindStringSubmatch(name)
		if len(sub) < 2 {
			continue
		}

		versionsMap[sub[1]] = struct{}{}
	}
	if len(versionsMap) == 0 {
		return "", errors.New("no gcloud-sdk versions found")
	}

	versions := slices.SortedFunc(maps.Keys(versionsMap), func(a, b string) int {
		// semver.Compare requires a "v" prefix, so we add it here.
		return semver.Compare("v"+a, "v"+b)
	})

	return versions[len(versions)-1], nil
}

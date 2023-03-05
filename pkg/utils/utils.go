package utils

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

func RunHelmGetValues(app string) (int, []byte) {
	out, err := getValues(app)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return http.StatusNotFound, []byte("Not found")
	}

	outString, err := yaml.Marshal(out)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return http.StatusNotFound, []byte("Not found")
	}

	return http.StatusOK, outString
}

func RunHelmGetReleases(release string) (int, []byte) {
	releases, err := listReleases(release)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return http.StatusNotFound, []byte("Not found")
	}

	outString, err := yaml.Marshal(releases)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return http.StatusNotFound, []byte("Not found")
	}

	return http.StatusOK, outString
}

func RunHelmGetRelease(release string) (int, []byte) {
	releases, err := getRelease(release)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return http.StatusNotFound, []byte("Not found")
	}

	outString, err := yaml.Marshal(releases)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return http.StatusNotFound, []byte("Not found")
	}

	return http.StatusOK, outString
}

package utils

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/release"
)

func newConfig() {

}

func getValues(app string) (map[string]interface{}, error) {
	settings := cli.New()
	actionConfig := new(action.Configuration)

	if err := actionConfig.Init(settings.RESTClientGetter(), app, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	client := action.NewGetValues(actionConfig)
	client.AllValues = true

	results, err := client.Run(app)

	if err != nil {
		return nil, err
	}

	return results, nil
}

func listReleases(release string) ([]string, error) {
	settings := cli.New()
	actionConfig := new(action.Configuration)

	if err := actionConfig.Init(settings.RESTClientGetter(), release, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	client := action.NewList(actionConfig)
	client.Short = true
	client.Deployed = true
	if release == "" {
		client.AllNamespaces = true
	}

	results, err := client.Run()

	names := make([]string, 0, len(results))
	for _, res := range results {
		names = append(names, res.Name)
	}

	if err != nil {
		return nil, err
	}

	return names, nil
}

func getRelease(release string) (*release.Release, error) {
	settings := cli.New()
	actionConfig := new(action.Configuration)

	if err := actionConfig.Init(settings.RESTClientGetter(), release, os.Getenv("HELM_DRIVER"), log.Printf); err != nil {
		log.Printf("%+v", err)
		return nil, err
	}

	client := action.NewGet(actionConfig)

	results, err := client.Run(release)

	if err != nil {
		return nil, err
	}

	return results, nil
}

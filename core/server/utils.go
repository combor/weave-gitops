package server

import (
	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func getMatchingLabels(appName string) client.MatchingLabels {
	return matchLabel(withPartOfLabel(appName))
}

func GetTenant(namespace, clusterName string, clusterUserNamespaces map[string][]v1.Namespace) string {
	for _, ns := range clusterUserNamespaces[clusterName] {
		if ns.GetName() == namespace {
			return ns.Labels["toolkit.fluxcd.io/tenant"]
		}
	}

	return ""
}

// ExtractValueFromMap gets string value from map or empty string if the value is empty
func ExtractStringValueFromMap(mapName map[string]string, key string) string {
	value, ok := mapName[key]
	if !ok {
		return ""
	}
	return value
}

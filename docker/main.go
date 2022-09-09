package main

import (
	"gnocchi.irideos.it/m/v2/pkg"
)

func main() {
	token := pkg.NewAuthentication()
	client := pkg.NewK8sClient()
	configMap := pkg.GetConfigMap(client, "grafana", "gnocchi-datasource")
	configMap = pkg.UpdateConfigMap(token, configMap, "datasource.yaml")
	pkg.ApplyConfigMap(client, "grafana", configMap)
}

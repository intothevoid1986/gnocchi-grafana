package pkg

import (
	"context"
	"errors"
	"log"
	"strings"

	"gnocchi.irideos.it/m/v2/utils"
	apiV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// NewK8sClient function create a new Kubernetes Client
func NewK8sClient() *kubernetes.Clientset {
	config, err := rest.InClusterConfig()
	utils.HandleError(err)
	k8sClient, err := kubernetes.NewForConfig(config)
	utils.HandleError(err)
	return k8sClient
}

// GetConfigMap function retrieve a configmap in certain namespace, with certain name
func GetConfigMap(client *kubernetes.Clientset, namespace string, name string) *apiV1.ConfigMap {
	configMap, err := client.CoreV1().ConfigMaps(namespace).Get(context.TODO(), name, v1.GetOptions{})
	utils.HandleError(err)
	log.Printf("fetched configmap %v", name)
	return configMap
}

// UpdateConfigMap function update the token field in configmap and push it to the cluster
func UpdateConfigMap(token string, configMap *apiV1.ConfigMap, dataName string) *apiV1.ConfigMap {
	oldToken := findOldToken(configMap.Data[dataName])
	configMap.Data[dataName] = strings.Replace(configMap.Data[dataName], oldToken, token, 1)
	log.Printf("configMap %v updated", configMap.Name)
	return configMap
}

// ApplyConfigMap function apply the new generated configMap to the target Kubernetes Cluster
func ApplyConfigMap(client *kubernetes.Clientset, namespace string, configMap *apiV1.ConfigMap) {
	client.CoreV1().ConfigMaps(namespace).Update(context.TODO(), configMap, v1.UpdateOptions{})
	log.Printf("configuration applied for configmap %v", configMap.Name)
}

func extractFields(data string) []string {
	return strings.Fields(data)
}

func findOldToken(data string) (oldToken string) {
	fields := extractFields(data)
	tokenIdx := -1
	for i, v := range fields {
		if v == "token:" {
			tokenIdx = i + 1
			break
		}
	}
	if tokenIdx < 0 {
		utils.HandleError(errors.New("could not find 'token' key in configmap"))
	}
	return fields[tokenIdx]
}

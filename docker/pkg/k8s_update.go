package pkg

import (
	"context"
	"errors"
	"log"
	"path"
	"strings"

	"gnocchi.irideos.it/m/v2/openstack-kubectl/utils"
	apiV1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// NewK8sClient function create a new Kubernetes Client
func NewK8sClient() *kubernetes.Clientset {
	config, err := clientcmd.BuildConfigFromFlags("", path.Join(homedir.HomeDir(), ".kube/config"))
	utils.HandleError(err)
	k8sClient, err := kubernetes.NewForConfig(config)
	utils.HandleError(err)
	return k8sClient
}

// GetConfigMap function retrieve a configmap in certain namespace, with certain name
func GetConfigMap(client *kubernetes.Clientset, namespace string, name string) *apiV1.ConfigMap {
	configMap, err := client.CoreV1().ConfigMaps(namespace).Get(context.Background(), name, v1.GetOptions{})
	utils.HandleError(err)
	return configMap
}

// UpdateConfigMap function update the token field in configmap and push it to the cluster
func UpdateConfigMap(token string, configMap *apiV1.ConfigMap, dataName string) *apiV1.ConfigMap {
	log.Print("new token issued")
	oldToken := findOldToken(configMap.Data[dataName])
	configMap.Data[dataName] = strings.Replace(configMap.Data[dataName], oldToken, token, 1)
	log.Print("configMap updated")
	return configMap
}

// ApplyConfigMap function apply the new generated configMap to the target Kubernetes Cluster
func ApplyConfigMap(client *kubernetes.Clientset, namespace string, configMap *apiV1.ConfigMap) {
	client.CoreV1().ConfigMaps(namespace).Update(context.Background(), configMap, v1.UpdateOptions{})
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
		utils.HandleError(errors.New("could not find token"))
	}
	return fields[tokenIdx]
}

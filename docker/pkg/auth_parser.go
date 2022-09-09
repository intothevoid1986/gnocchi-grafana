package pkg

import (
	"os"

	"gnocchi.irideos.it/m/v2/openstack-kubectl/utils"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

type authentication struct {
}

// NewAuthentication build a new authentication and return a token
func NewAuthentication() (token string) {
	auth := authentication{}
	return auth.authenticate()
}

func (auth *authentication) authenticate() string {
	authOpts, err := openstack.AuthOptionsFromEnv()
	utils.HandleError(err)
	authOpts.DomainName = "Default"
	provider, err := openstack.AuthenticatedClient(authOpts)
	utils.HandleError(err)
	client, err := openstack.NewIdentityV2(provider, gophercloud.EndpointOpts{
		Region: os.Getenv("OS_REGION_NAME"),
	})
	utils.HandleError(err)
	return client.TokenID
}

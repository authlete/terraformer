package authlete

import (
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	authlete "github.com/authlete/openapi-for-go"
)

type AuthleteService struct { // nolint
	terraformutils.Service
}

func (this *AuthleteService) getClient() *authlete.APIClient {

	cnf := authlete.NewConfiguration()
	cnf.UserAgent = "terraformer-authlete"
	cnf.Servers[0].URL = this.GetArgs()["api_server"].(string)

	return authlete.NewAPIClient(cnf)
}

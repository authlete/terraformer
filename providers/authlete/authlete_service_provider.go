package authlete

import (
	"errors"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/zclconf/go-cty/cty"
)

type AuthleteProvider struct { // nolint
	terraformutils.Provider
	apiServer string
	soKey     string
	soSecret  string
	apiKey    string
	apiSecret string
}

func (this *AuthleteProvider) GetResourceConnections() map[string]map[string][]string {
	return map[string]map[string][]string{}
}

func (this *AuthleteProvider) GetProviderData(arg ...string) map[string]interface{} {
	authleteConfig := map[string]interface{}{}

	authleteConfig["api_server"] = this.apiServer
	authleteConfig["service_owner_key"] = this.soKey
	authleteConfig["service_owner_secret"] = this.soSecret
	authleteConfig["api_key"] = this.apiKey
	authleteConfig["api_secret"] = this.apiSecret

	return map[string]interface{}{
		"provider": map[string]interface{}{
			"authlete": authleteConfig,
		},
	}
}

func (this *AuthleteProvider) Init(args []string) error {
	this.apiServer = args[0]
	this.soKey = args[1]
	this.soSecret = args[2]
	this.apiKey = args[3]
	this.apiSecret = args[4]
	return nil
}

func (this *AuthleteProvider) GetName() string {
	return "authlete"
}

func (this *AuthleteProvider) GetConfig() cty.Value {
	return cty.ObjectVal(map[string]cty.Value{
		"api_server":           cty.StringVal(this.apiServer),
		"service_owner_key":    cty.StringVal(this.soKey),
		"service_owner_secret": cty.StringVal(this.soSecret),
		"api_key":              cty.StringVal(this.apiKey),
		"api_secret":           cty.StringVal(this.apiSecret),
	})
}

func (this *AuthleteProvider) InitService(serviceName string, verbose bool) error {
	var isSupported bool
	if _, isSupported = this.GetSupportedService()[serviceName]; !isSupported {
		return errors.New(this.GetName() + ": " + serviceName + " not supported service")
	}
	this.Service = this.GetSupportedService()[serviceName]
	this.Service.SetName(serviceName)
	this.Service.SetVerbose(verbose)
	this.Service.SetProviderName(this.GetName())
	this.Service.SetArgs(map[string]interface{}{
		"api_server":           this.apiServer,
		"service_owner_key":    this.soKey,
		"service_owner_secret": this.soSecret,
		"api_key":              this.apiKey,
		"api_secret":           this.apiSecret,
	})
	return nil
}

func (this *AuthleteProvider) GetSupportedService() map[string]terraformutils.ServiceGenerator {
	return map[string]terraformutils.ServiceGenerator{
		"authlete_service": &ServiceGenerator{},
		"authlete_client":  &ClientGenerator{},
	}
}

package client

import (
	cmd "contabo.com/cli/cntb/cmd"
	"contabo.com/cli/cntb/config"
	"contabo.com/cli/cntb/oauth2Client"
	apiClient "contabo.com/cli/cntb/openapi"
)

func ApiClient() *apiClient.APIClient {
	configuration := apiClient.NewConfiguration()
	configuration.AddDefaultHeader("x-trace-id", "arcus_cli")
	configuration.HTTPClient = oauth2Client.BearerHttpClient(config.Conf.Oauth2TokenUrl, config.Conf.Oauth2ClientId, config.Conf.Oauth2ClientSecret, config.Conf.Oauth2User, config.Conf.Oauth2Password)
	var server apiClient.ServerConfiguration
	server.URL = config.Conf.Api
	var serverConfigurations []apiClient.ServerConfiguration
	serverConfigurations = append(serverConfigurations, server)
	configuration.Servers = serverConfigurations
	if cmd.DebugLevel == "trace" {
		configuration.Debug = true
	}

	return apiClient.NewAPIClient(configuration)
}
